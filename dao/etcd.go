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

package dao

import (
	"strings"
	"time"

	"go.etcd.io/etcd/clientv3"

	"go-mysql-transfer/config"
	"go-mysql-transfer/util/logagent"
)

var (
	_etcdConn      *clientv3.Client
	_etcdOps       clientv3.KV
	_etcdAddresses []string
)

func initEtcd(app *config.AppConfig) error {
	clientv3.SetLogger(logagent.NewEtcdLoggerAgent())
	addresses := strings.Split(app.GetClusterConfig().GetEtcdAddrs(), ",")
	conn, err := clientv3.New(clientv3.Config{
		Endpoints:   addresses,
		Username:    app.GetClusterConfig().GetEtcdUser(),
		Password:    app.GetClusterConfig().GetEtcdPassword(),
		DialTimeout: 1 * time.Second,
	})
	if err != nil {
		return err
	}

	//err = etcdutils.CreateNodeIfNecessary(conf.PositionDir(""), string(data), _ops)
	//if err != nil {
	//	return err
	//}

	_etcdConn = conn
	_etcdOps = clientv3.NewKV(conn)
	_etcdAddresses = addresses

	return nil
}

func closeEtcd() {
	if _etcdConn != nil {
		_etcdConn.Close()
	}
}

func GetEtcdConn() *clientv3.Client {
	return _etcdConn
}

func GetEtcdOps() clientv3.KV {
	return _etcdOps
}
