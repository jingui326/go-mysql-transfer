// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: proto/pipeline.proto

package po

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 定义'ES索引映射'结构
type EsIndexMapping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Column     string `protobuf:"bytes,1,opt,name=column,proto3" json:"column,omitempty"`         // 数据库列名称
	EsField    string `protobuf:"bytes,2,opt,name=esField,proto3" json:"esField,omitempty"`       // 映射后的ES字段名称
	EsType     string `protobuf:"bytes,3,opt,name=esType,proto3" json:"esType,omitempty"`         // ES字段类型
	EsAnalyzer string `protobuf:"bytes,4,opt,name=esAnalyzer,proto3" json:"esAnalyzer,omitempty"` // ES分词器
}

func (x *EsIndexMapping) Reset() {
	*x = EsIndexMapping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pipeline_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EsIndexMapping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EsIndexMapping) ProtoMessage() {}

func (x *EsIndexMapping) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pipeline_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EsIndexMapping.ProtoReflect.Descriptor instead.
func (*EsIndexMapping) Descriptor() ([]byte, []int) {
	return file_proto_pipeline_proto_rawDescGZIP(), []int{0}
}

func (x *EsIndexMapping) GetColumn() string {
	if x != nil {
		return x.Column
	}
	return ""
}

func (x *EsIndexMapping) GetEsField() string {
	if x != nil {
		return x.EsField
	}
	return ""
}

func (x *EsIndexMapping) GetEsType() string {
	if x != nil {
		return x.EsType
	}
	return ""
}

func (x *EsIndexMapping) GetEsAnalyzer() string {
	if x != nil {
		return x.EsAnalyzer
	}
	return ""
}

// 定义'转换规则'结构
type Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        int32 `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`               //0规则 1脚本
	ReceiveType int32 `protobuf:"varint,2,opt,name=receiveType,proto3" json:"receiveType,omitempty"` //0:insert&update&delete  1:insert&update 2:insert&delete 3:insert
	// ===================== read =============================
	Schema                       string            `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema,omitempty"`
	Table                        string            `protobuf:"bytes,4,opt,name=table,proto3" json:"table,omitempty"`
	ColumnNameFormatter          int32             `protobuf:"varint,5,opt,name=columnNameFormatter,proto3" json:"columnNameFormatter,omitempty"`                                                                                                          //列名转换格式 0:列名称转为小写 1:列名称转为大写 2:列名称下划线转驼峰
	ExcludeColumnList            []string          `protobuf:"bytes,6,rep,name=excludeColumnList,proto3" json:"excludeColumnList,omitempty"`                                                                                                               // 排除掉的列
	ColumnNameMapping            map[string]string `protobuf:"bytes,7,rep,name=columnNameMapping,proto3" json:"columnNameMapping,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`                       // 列名称映射
	AdditionalColumnValueMapping map[string]string `protobuf:"bytes,8,rep,name=additionalColumnValueMapping,proto3" json:"additionalColumnValueMapping,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 附加的字段和值
	DataEncoder                  int32             `protobuf:"varint,9,opt,name=dataEncoder,proto3" json:"dataEncoder,omitempty"`                                                                                                                          //数据编码类型，0: json、 1:表达式
	DataExpression               string            `protobuf:"bytes,10,opt,name=dataExpression,proto3" json:"dataExpression,omitempty"`                                                                                                                    //数据expression,{id}表示字段id的值、{name}表示字段name的值
	DateFormatter                string            `protobuf:"bytes,11,opt,name=dateFormatter,proto3" json:"dateFormatter,omitempty"`                                                                                                                      //date类型格式化， 不填写默认2006-01-02
	DatetimeFormatter            string            `protobuf:"bytes,12,opt,name=datetimeFormatter,proto3" json:"datetimeFormatter,omitempty"`                                                                                                              //datetime、timestamp类型格式化，不填写默认RFC3339(2006-01-02T15:04:05Z07:00)
	ReserveCoveredData           bool              `protobuf:"varint,13,opt,name=reserveCoveredData,proto3" json:"reserveCoveredData,omitempty"`                                                                                                           // 保留update之前的数据
	OrderColumn                  string            `protobuf:"bytes,14,opt,name=orderColumn,proto3" json:"orderColumn,omitempty"`                                                                                                                          // 排序列
	// ===================== write =============================
	// redis
	RedisStructure            int32  `protobuf:"varint,15,opt,name=redisStructure,proto3" json:"redisStructure,omitempty"`                      //对应redis的5种数据类型 1:String、2:Hash(字典) 、3:List(列表) 、4:Set(集合)、5:Sorted Set(有序集合)
	RedisKeyPrefix            string `protobuf:"bytes,16,opt,name=redisKeyPrefix,proto3" json:"redisKeyPrefix,omitempty"`                       //key的前缀
	RedisKeyBuilder           int32  `protobuf:"varint,17,opt,name=redisKeyBuilder,proto3" json:"redisKeyBuilder,omitempty"`                    //key生成方式，0:使用列值(默认使用主键)、1:表达式、2固定值
	RedisKeyColumn            string `protobuf:"bytes,18,opt,name=redisKeyColumn,proto3" json:"redisKeyColumn,omitempty"`                       //key生成方式，使用列值，默认使用主键
	RedisKeyExpression        string `protobuf:"bytes,19,opt,name=redisKeyExpression,proto3" json:"redisKeyExpression,omitempty"`               // key生成方式，格式化,如{id}-{name}；{id}表示字段id的值、{name}表示字段name的值
	RedisKeyFixValue          string `protobuf:"bytes,20,opt,name=redisKeyFixValue,proto3" json:"redisKeyFixValue,omitempty"`                   // key固定值
	RedisHashFieldPrefix      string `protobuf:"bytes,21,opt,name=redisHashFieldPrefix,proto3" json:"redisHashFieldPrefix,omitempty"`           // hash的field前缀，仅redis_structure为hash时起作用
	RedisHashFieldColumn      string `protobuf:"bytes,22,opt,name=redisHashFieldColumn,proto3" json:"redisHashFieldColumn,omitempty"`           // 使用哪个列的值作为hash的field，仅redis_structure为hash时起作用
	RedisSortedSetScoreColumn string `protobuf:"bytes,23,opt,name=redisSortedSetScoreColumn,proto3" json:"redisSortedSetScoreColumn,omitempty"` // Sorted Set(有序集合)的Score
	// mongodbRule
	MongodbDatabase   string `protobuf:"bytes,24,opt,name=mongodbDatabase,proto3" json:"mongodbDatabase,omitempty"`     //mongodb database 不能为空
	MongodbCollection string `protobuf:"bytes,25,opt,name=mongodbCollection,proto3" json:"mongodbCollection,omitempty"` //mongodb collection，可以为空，默认使用表(Table)名称
	// elasticsearch
	EsIndexBuildType int32  `protobuf:"varint,26,opt,name=esIndexBuildType,proto3" json:"esIndexBuildType,omitempty"` //Index名称创建方式，0使用已经存在的、1自动创建
	EsIndexName      string `protobuf:"bytes,27,opt,name=esIndexName,proto3" json:"esIndexName,omitempty"`            //Index名称,可以为空，默认使用表(Table)名称
	//string esType          =31;         //es6.x以后一个Index只能拥有一个Type,可以为空，默认使用_doc; es7.x版本此属性无效
	EsIndexMappings []*EsIndexMapping `protobuf:"bytes,28,rep,name=esIndexMappings,proto3" json:"esIndexMappings,omitempty"` //Elasticsearch mappings映射关系,可以为空，为空时根据数据类型自己推导
	// Rocketmq Kafka Rabbitmq
	MqTopic string `protobuf:"bytes,29,opt,name=mqTopic,proto3" json:"mqTopic,omitempty"` // topic名称，可以为空，为空时使用表名称
	// HTTP 参数名称
	HttpParameterName  string `protobuf:"bytes,30,opt,name=httpParameterName,proto3" json:"httpParameterName,omitempty"`    // http参数名称
	HttpReserveRawData bool   `protobuf:"varint,31,opt,name=httpReserveRawData,proto3" json:"httpReserveRawData,omitempty"` // 保留update之前的数据
	// =============================  LUA  ==============================================
	LuaScript string `protobuf:"bytes,32,opt,name=LuaScript,proto3" json:"LuaScript,omitempty"` //lua 脚本
}

func (x *Rule) Reset() {
	*x = Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pipeline_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule) ProtoMessage() {}

func (x *Rule) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pipeline_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rule.ProtoReflect.Descriptor instead.
func (*Rule) Descriptor() ([]byte, []int) {
	return file_proto_pipeline_proto_rawDescGZIP(), []int{1}
}

func (x *Rule) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Rule) GetReceiveType() int32 {
	if x != nil {
		return x.ReceiveType
	}
	return 0
}

func (x *Rule) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *Rule) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *Rule) GetColumnNameFormatter() int32 {
	if x != nil {
		return x.ColumnNameFormatter
	}
	return 0
}

func (x *Rule) GetExcludeColumnList() []string {
	if x != nil {
		return x.ExcludeColumnList
	}
	return nil
}

func (x *Rule) GetColumnNameMapping() map[string]string {
	if x != nil {
		return x.ColumnNameMapping
	}
	return nil
}

func (x *Rule) GetAdditionalColumnValueMapping() map[string]string {
	if x != nil {
		return x.AdditionalColumnValueMapping
	}
	return nil
}

func (x *Rule) GetDataEncoder() int32 {
	if x != nil {
		return x.DataEncoder
	}
	return 0
}

func (x *Rule) GetDataExpression() string {
	if x != nil {
		return x.DataExpression
	}
	return ""
}

func (x *Rule) GetDateFormatter() string {
	if x != nil {
		return x.DateFormatter
	}
	return ""
}

func (x *Rule) GetDatetimeFormatter() string {
	if x != nil {
		return x.DatetimeFormatter
	}
	return ""
}

func (x *Rule) GetReserveCoveredData() bool {
	if x != nil {
		return x.ReserveCoveredData
	}
	return false
}

func (x *Rule) GetOrderColumn() string {
	if x != nil {
		return x.OrderColumn
	}
	return ""
}

func (x *Rule) GetRedisStructure() int32 {
	if x != nil {
		return x.RedisStructure
	}
	return 0
}

func (x *Rule) GetRedisKeyPrefix() string {
	if x != nil {
		return x.RedisKeyPrefix
	}
	return ""
}

func (x *Rule) GetRedisKeyBuilder() int32 {
	if x != nil {
		return x.RedisKeyBuilder
	}
	return 0
}

func (x *Rule) GetRedisKeyColumn() string {
	if x != nil {
		return x.RedisKeyColumn
	}
	return ""
}

func (x *Rule) GetRedisKeyExpression() string {
	if x != nil {
		return x.RedisKeyExpression
	}
	return ""
}

func (x *Rule) GetRedisKeyFixValue() string {
	if x != nil {
		return x.RedisKeyFixValue
	}
	return ""
}

func (x *Rule) GetRedisHashFieldPrefix() string {
	if x != nil {
		return x.RedisHashFieldPrefix
	}
	return ""
}

func (x *Rule) GetRedisHashFieldColumn() string {
	if x != nil {
		return x.RedisHashFieldColumn
	}
	return ""
}

func (x *Rule) GetRedisSortedSetScoreColumn() string {
	if x != nil {
		return x.RedisSortedSetScoreColumn
	}
	return ""
}

func (x *Rule) GetMongodbDatabase() string {
	if x != nil {
		return x.MongodbDatabase
	}
	return ""
}

func (x *Rule) GetMongodbCollection() string {
	if x != nil {
		return x.MongodbCollection
	}
	return ""
}

func (x *Rule) GetEsIndexBuildType() int32 {
	if x != nil {
		return x.EsIndexBuildType
	}
	return 0
}

func (x *Rule) GetEsIndexName() string {
	if x != nil {
		return x.EsIndexName
	}
	return ""
}

func (x *Rule) GetEsIndexMappings() []*EsIndexMapping {
	if x != nil {
		return x.EsIndexMappings
	}
	return nil
}

func (x *Rule) GetMqTopic() string {
	if x != nil {
		return x.MqTopic
	}
	return ""
}

func (x *Rule) GetHttpParameterName() string {
	if x != nil {
		return x.HttpParameterName
	}
	return ""
}

func (x *Rule) GetHttpReserveRawData() bool {
	if x != nil {
		return x.HttpReserveRawData
	}
	return false
}

func (x *Rule) GetLuaScript() string {
	if x != nil {
		return x.LuaScript
	}
	return ""
}

// 定义'管道信息'结构
type PipelineInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                    uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                  string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SourceId              uint64  `protobuf:"varint,3,opt,name=sourceId,proto3" json:"sourceId,omitempty"`
	EndpointId            uint64  `protobuf:"varint,4,opt,name=endpointId,proto3" json:"endpointId,omitempty"`
	EndpointType          uint32  `protobuf:"varint,5,opt,name=endpointType,proto3" json:"endpointType,omitempty"`
	CreateTime            string  `protobuf:"bytes,6,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime            string  `protobuf:"bytes,7,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
	Status                uint32  `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	StreamBulkSize        uint64  `protobuf:"varint,9,opt,name=streamBulkSize,proto3" json:"streamBulkSize,omitempty"`
	StreamFlushInterval   uint32  `protobuf:"varint,10,opt,name=streamFlushInterval,proto3" json:"streamFlushInterval,omitempty"`
	BatchCoroutines       uint32  `protobuf:"varint,11,opt,name=batchCoroutines,proto3" json:"batchCoroutines,omitempty"`
	BatchBulkSize         uint32  `protobuf:"varint,12,opt,name=batchBulkSize,proto3" json:"batchBulkSize,omitempty"`
	AlarmMailList         string  `protobuf:"bytes,13,opt,name=alarmMailList,proto3" json:"alarmMailList,omitempty"`
	AlarmWebhook          string  `protobuf:"bytes,14,opt,name=alarmWebhook,proto3" json:"alarmWebhook,omitempty"`
	AlarmWebhookSecretKey string  `protobuf:"bytes,15,opt,name=alarmWebhookSecretKey,proto3" json:"alarmWebhookSecretKey,omitempty"`
	AlarmItemList         string  `protobuf:"bytes,16,opt,name=alarmItemList,proto3" json:"alarmItemList,omitempty"`
	DataVersion           int32   `protobuf:"varint,17,opt,name=dataVersion,proto3" json:"dataVersion,omitempty"` //数据版本
	Rules                 []*Rule `protobuf:"bytes,18,rep,name=rules,proto3" json:"rules,omitempty"`              // 转换规则列表
}

func (x *PipelineInfo) Reset() {
	*x = PipelineInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pipeline_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineInfo) ProtoMessage() {}

func (x *PipelineInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pipeline_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineInfo.ProtoReflect.Descriptor instead.
func (*PipelineInfo) Descriptor() ([]byte, []int) {
	return file_proto_pipeline_proto_rawDescGZIP(), []int{2}
}

func (x *PipelineInfo) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PipelineInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PipelineInfo) GetSourceId() uint64 {
	if x != nil {
		return x.SourceId
	}
	return 0
}

func (x *PipelineInfo) GetEndpointId() uint64 {
	if x != nil {
		return x.EndpointId
	}
	return 0
}

func (x *PipelineInfo) GetEndpointType() uint32 {
	if x != nil {
		return x.EndpointType
	}
	return 0
}

func (x *PipelineInfo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *PipelineInfo) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

func (x *PipelineInfo) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *PipelineInfo) GetStreamBulkSize() uint64 {
	if x != nil {
		return x.StreamBulkSize
	}
	return 0
}

func (x *PipelineInfo) GetStreamFlushInterval() uint32 {
	if x != nil {
		return x.StreamFlushInterval
	}
	return 0
}

func (x *PipelineInfo) GetBatchCoroutines() uint32 {
	if x != nil {
		return x.BatchCoroutines
	}
	return 0
}

func (x *PipelineInfo) GetBatchBulkSize() uint32 {
	if x != nil {
		return x.BatchBulkSize
	}
	return 0
}

func (x *PipelineInfo) GetAlarmMailList() string {
	if x != nil {
		return x.AlarmMailList
	}
	return ""
}

func (x *PipelineInfo) GetAlarmWebhook() string {
	if x != nil {
		return x.AlarmWebhook
	}
	return ""
}

func (x *PipelineInfo) GetAlarmWebhookSecretKey() string {
	if x != nil {
		return x.AlarmWebhookSecretKey
	}
	return ""
}

func (x *PipelineInfo) GetAlarmItemList() string {
	if x != nil {
		return x.AlarmItemList
	}
	return ""
}

func (x *PipelineInfo) GetDataVersion() int32 {
	if x != nil {
		return x.DataVersion
	}
	return 0
}

func (x *PipelineInfo) GetRules() []*Rule {
	if x != nil {
		return x.Rules
	}
	return nil
}

var File_proto_pipeline_proto protoreflect.FileDescriptor

var file_proto_pipeline_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x0e, 0x45, 0x73, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x65, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x73,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x73, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x73, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a,
	0x65, 0x72, 0x22, 0xa5, 0x0c, 0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x30, 0x0a, 0x13, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x63, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x12, 0x2c, 0x0a, 0x11, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x43, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x65, 0x78,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x4a, 0x0a, 0x11, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x52, 0x75, 0x6c,
	0x65, 0x2e, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x11, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x6b, 0x0a, 0x1c, 0x61,
	0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x1c, 0x61, 0x64, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x64,
	0x61, 0x74, 0x61, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x74, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x64, 0x61, 0x74, 0x65,
	0x74, 0x69, 0x6d, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x11, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x46, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x12, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x43, 0x6f, 0x76, 0x65, 0x72,
	0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x43,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x64, 0x69,
	0x73, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x4b, 0x65, 0x79, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x4b,
	0x65, 0x79, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x64, 0x69,
	0x73, 0x4b, 0x65, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x4b, 0x65, 0x79, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x4b, 0x65, 0x79, 0x43, 0x6f,
	0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x64, 0x69,
	0x73, 0x4b, 0x65, 0x79, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x72, 0x65,
	0x64, 0x69, 0x73, 0x4b, 0x65, 0x79, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x64, 0x69, 0x73, 0x4b, 0x65, 0x79,
	0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65,
	0x64, 0x69, 0x73, 0x4b, 0x65, 0x79, 0x46, 0x69, 0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x72, 0x65, 0x64, 0x69, 0x73, 0x4b, 0x65, 0x79, 0x46, 0x69,
	0x78, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x72, 0x65, 0x64, 0x69, 0x73, 0x48,
	0x61, 0x73, 0x68, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x15,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x72, 0x65, 0x64, 0x69, 0x73, 0x48, 0x61, 0x73, 0x68, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x32, 0x0a, 0x14, 0x72, 0x65,
	0x64, 0x69, 0x73, 0x48, 0x61, 0x73, 0x68, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x6f, 0x6c, 0x75,
	0x6d, 0x6e, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x72, 0x65, 0x64, 0x69, 0x73, 0x48,
	0x61, 0x73, 0x68, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x3c,
	0x0a, 0x19, 0x72, 0x65, 0x64, 0x69, 0x73, 0x53, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x53, 0x65, 0x74,
	0x53, 0x63, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x17, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x19, 0x72, 0x65, 0x64, 0x69, 0x73, 0x53, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x53, 0x65,
	0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x28, 0x0a, 0x0f,
	0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18,
	0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64,
	0x62, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x19, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x6d, 0x6f, 0x6e, 0x67, 0x6f, 0x64, 0x62, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x65, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10,
	0x65, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x65, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0f, 0x65, 0x73, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x1c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x45, 0x73,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x0f, 0x65, 0x73,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x71, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x71, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x2c, 0x0a, 0x11, 0x68, 0x74, 0x74, 0x70, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x68, 0x74, 0x74, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x52, 0x61, 0x77, 0x44, 0x61, 0x74, 0x61, 0x18, 0x1f, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x12, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x52, 0x61,
	0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x75, 0x61, 0x53, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4c, 0x75, 0x61, 0x53, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x1a, 0x44, 0x0a, 0x16, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4f, 0x0a, 0x21, 0x41, 0x64, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf9, 0x04, 0x0a, 0x0c, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x42, 0x75, 0x6c, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x42, 0x75, 0x6c, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x30, 0x0a, 0x13, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x12, 0x28, 0x0a, 0x0f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6f, 0x72, 0x6f, 0x75, 0x74,
	0x69, 0x6e, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x43, 0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x42, 0x75, 0x6c, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0d, 0x62, 0x61, 0x74, 0x63, 0x68, 0x42, 0x75, 0x6c, 0x6b, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x4d, 0x61, 0x69, 0x6c, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x4d,
	0x61, 0x69, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x6c, 0x61, 0x72, 0x6d,
	0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61,
	0x6c, 0x61, 0x72, 0x6d, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12, 0x34, 0x0a, 0x15, 0x61,
	0x6c, 0x61, 0x72, 0x6d, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x61, 0x6c, 0x61, 0x72,
	0x6d, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65,
	0x79, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x49,
	0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x64, 0x61,
	0x74, 0x61, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x05, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x52,
	0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2f, 0x70, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_pipeline_proto_rawDescOnce sync.Once
	file_proto_pipeline_proto_rawDescData = file_proto_pipeline_proto_rawDesc
)

func file_proto_pipeline_proto_rawDescGZIP() []byte {
	file_proto_pipeline_proto_rawDescOnce.Do(func() {
		file_proto_pipeline_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_pipeline_proto_rawDescData)
	})
	return file_proto_pipeline_proto_rawDescData
}

var file_proto_pipeline_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_pipeline_proto_goTypes = []interface{}{
	(*EsIndexMapping)(nil), // 0: EsIndexMapping
	(*Rule)(nil),           // 1: Rule
	(*PipelineInfo)(nil),   // 2: PipelineInfo
	nil,                    // 3: Rule.ColumnNameMappingEntry
	nil,                    // 4: Rule.AdditionalColumnValueMappingEntry
}
var file_proto_pipeline_proto_depIdxs = []int32{
	3, // 0: Rule.columnNameMapping:type_name -> Rule.ColumnNameMappingEntry
	4, // 1: Rule.additionalColumnValueMapping:type_name -> Rule.AdditionalColumnValueMappingEntry
	0, // 2: Rule.esIndexMappings:type_name -> EsIndexMapping
	1, // 3: PipelineInfo.rules:type_name -> Rule
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_pipeline_proto_init() }
func file_proto_pipeline_proto_init() {
	if File_proto_pipeline_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_pipeline_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EsIndexMapping); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_pipeline_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rule); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_pipeline_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_pipeline_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_pipeline_proto_goTypes,
		DependencyIndexes: file_proto_pipeline_proto_depIdxs,
		MessageInfos:      file_proto_pipeline_proto_msgTypes,
	}.Build()
	File_proto_pipeline_proto = out.File
	file_proto_pipeline_proto_rawDesc = nil
	file_proto_pipeline_proto_goTypes = nil
	file_proto_pipeline_proto_depIdxs = nil
}
