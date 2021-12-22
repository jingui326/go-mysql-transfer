/*
 * Copyright 2021-2022 the original author(https://github.com/wj596)
 *
 * <p>
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 * </p>
 */

package mongodb

import (
	"context"

	"github.com/juju/errors"
	"go.mongodb.org/mongo-driver/mongo"

	"go-mysql-transfer/domain/bo"
	"go-mysql-transfer/domain/po"
	"go-mysql-transfer/util/log"
)

type StreamEndpoint struct {
	endpoint *Endpoint
}

func NewStreamEndpoint(info *po.EndpointInfo) *StreamEndpoint {
	return &StreamEndpoint{
		endpoint: &Endpoint{
			info:        info,
			collections: make(map[string]*mongo.Collection),
		},
	}
}

func (s *StreamEndpoint) Stream(requests []*bo.RowEventRequest) error {
	models := make(map[string][]mongo.WriteModel, 0)

	for _, request := range requests {
		ctx := request.Context
		if ctx.GetTableColumnCount() != len(request.Data) {
			log.Warnf("管道[%s]，表[%s]的结构发生变更，忽略此条数据", ctx.GetPipelineName(), ctx.GetTableFullName())
			continue
		}

		if ctx.IsLuaEnable() {
			err := s.endpoint.parseByLua(request, ctx, models, nil)
			if err != nil {
				return err
			}
		} else {
			err := s.endpoint.parseByRegular(request, ctx, models)
			if err != nil {
				return err
			}
		}
	}

	hasDuplicateKey := false
	for collectionName, array := range models {
		collection := s.endpoint.getCollection(collectionName)
		if nil == collection {
			return errors.New("collection为空")
		}
		_, err := collection.BulkWrite(context.Background(), array)
		if err != nil {
			if s.endpoint.isDuplicateKeyError(err.Error()) {
				log.Warnf(err.Error())
				hasDuplicateKey = true
				break
			} else {
				return err
			}
		}
	}

	if hasDuplicateKey {
		_, err := s.endpoint.doSlowly(models)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *StreamEndpoint) Connect() error {
	return s.endpoint.Connect()
}

func (s *StreamEndpoint) Ping() error {
	return s.endpoint.Ping()
}

func (s *StreamEndpoint) Close() {
	s.endpoint.Close()
}
