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

package vo

type CounterVO struct {
	Table string `json:"table"`
	Count uint64 `json:"count"`
}

type PipelineRuntimeVO struct {
	PipelineId     uint64       `json:"pipelineId,string"`
	PipelineName   string       `json:"pipelineName"`
	Status         uint32       `json:"status"`
	StartTime      string       `json:"startTime"`
	UpdateTime     string       `json:"updateTime"`
	PositionName   string       `json:"positionName"`
	PositionIndex  uint32       `json:"positionIndex"`
	InsertCount    uint64       `json:"insertCount"`
	UpdateCount    uint64       `json:"updateCount"`
	DeleteCount    uint64       `json:"deleteCount"`
	Message        string       `json:"message"`
	Node           string       `json:"node"`
	BatchStartTime string       `json:"batchStartTime"`
	BatchEndTime   string       `json:"batchEndTime"`
	TotalCounters  []*CounterVO `json:"totalCounters"`
	InsertCounters []*CounterVO `json:"insertCounters"`
	BatchPercent   float64      `json:"batchPercent"`
}
