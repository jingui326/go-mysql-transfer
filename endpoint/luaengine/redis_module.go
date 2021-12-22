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

package luaengine

import (
	"github.com/siddontang/go-mysql/canal"
	"github.com/yuin/gopher-lua"

	"go-mysql-transfer/domain/constants"
	"go-mysql-transfer/util/stringutils"
)

const (
	_string    = "0"
	_hash      = "1"
	_list      = "2"
	_set       = "3"
	_sortedSet = "4"
)

func preloadRedisModule(L *lua.LState) {
	L.PreloadModule("redis", redisModuleLoader)
}

func redisModuleLoader(L *lua.LState) int {
	t := L.NewTable()
	L.SetFuncs(t, redisApis)
	L.Push(t)
	return 1
}

var redisApis = map[string]lua.LGFunction{
	"SET": redisApiSet,
	"DEL": redisApiDel,

	"HSET": redisApiHashSet,
	"HDEL": redisApiHashDel,

	"RPUSH": redisApiRightPush,
	"LREM":  redisApiLeftRem,

	"SADD": redisApiSetAdd,
	"SREM": redisApiSetRem,

	"ZADD": redisApiSortedSetAdd,
	"ZREM": redisApiSortedSetRem,

	"EXPIRE": redisApiExpire,
}

func redisApiSet(L *lua.LState) int {
	key := L.CheckString(1)
	value := L.CheckAny(2)
	combine := stringutils.JoinWithUnderline(canal.InsertAction, _string, key)

	result := L.GetGlobal(constants.LuaGlobalVariableResult)
	L.SetTable(result, lua.LString(combine), value)
	return 0
}

func redisApiDel(L *lua.LState) int {
	key := L.CheckString(1)
	combine := stringutils.JoinWithUnderline(canal.DeleteAction, _string, key)

	result := L.GetGlobal(constants.LuaGlobalVariableResult)
	L.SetTable(result, lua.LString(combine), lua.LBool(true))
	return 0
}

func redisApiHashSet(L *lua.LState) int {
	key := L.CheckString(1)
	field := L.CheckAny(2)
	value := L.CheckAny(3)
	combine := stringutils.JoinWithUnderline(canal.InsertAction, _hash, nextId())

	hash := L.NewTable()
	L.SetTable(hash, lua.LString("key"), lua.LString(key))
	L.SetTable(hash, lua.LString("field"), field)
	L.SetTable(hash, lua.LString("value"), value)

	result := L.GetGlobal(constants.LuaGlobalVariableResult)
	L.SetTable(result, lua.LString(combine), hash)
	return 0
}

func redisApiHashDel(L *lua.LState) int {
	key := L.CheckAny(1)
	field := L.CheckAny(2)
	combine := stringutils.JoinWithUnderline(canal.DeleteAction, _hash, nextId())

	hash := L.NewTable()
	L.SetTable(hash, lua.LString("key"), key)
	L.SetTable(hash, lua.LString("field"), field)
	L.SetTable(hash, lua.LString("value"), lua.LNumber(1))

	result := L.GetGlobal(constants.LuaGlobalVariableResult)
	L.SetTable(result, lua.LString(combine), hash)
	return 0
}

func redisApiRightPush(L *lua.LState) int {
	key := L.CheckString(1)
	value := L.CheckAny(2)
	combine := stringutils.JoinWithUnderline(canal.InsertAction, _list, key)

	result := L.GetGlobal(constants.LuaGlobalVariableResult)
	L.SetTable(result, lua.LString(combine), value)
	return 0
}

func redisApiLeftRem(L *lua.LState) int {
	key := L.CheckString(1)
	value := L.CheckAny(2)
	combine := stringutils.JoinWithUnderline(canal.DeleteAction, _list, key)

	result := L.GetGlobal(constants.LuaGlobalVariableResult)
	L.SetTable(result, lua.LString(combine), value)
	return 0
}

func redisApiSetAdd(L *lua.LState) int {
	key := L.CheckString(1)
	value := L.CheckAny(2)
	combine := stringutils.JoinWithUnderline(canal.InsertAction, _set, key)

	result := L.GetGlobal(constants.LuaGlobalVariableResult)
	L.SetTable(result, lua.LString(combine), value)
	return 0
}

func redisApiSetRem(L *lua.LState) int {
	key := L.CheckString(1)
	value := L.CheckAny(2)
	combine := stringutils.JoinWithUnderline(canal.DeleteAction, _set, key)

	result := L.GetGlobal(constants.LuaGlobalVariableResult)
	L.SetTable(result, lua.LString(combine), value)
	return 0
}

func redisApiSortedSetAdd(L *lua.LState) int {
	key := L.CheckString(1)
	score := L.CheckAny(2)
	value := L.CheckAny(3)
	combine := stringutils.JoinWithUnderline(canal.InsertAction, _sortedSet, nextId())

	hash := L.NewTable()
	L.SetTable(hash, lua.LString("key"), lua.LString(key))
	L.SetTable(hash, lua.LString("score"), score)
	L.SetTable(hash, lua.LString("value"), value)

	result := L.GetGlobal(constants.LuaGlobalVariableResult)
	L.SetTable(result, lua.LString(combine), hash)
	return 0
}

func redisApiSortedSetRem(L *lua.LState) int {
	key := L.CheckString(1)
	value := L.CheckAny(2)
	combine := stringutils.JoinWithUnderline(canal.DeleteAction, _sortedSet, key)

	result := L.GetGlobal(constants.LuaGlobalVariableResult)
	L.SetTable(result, lua.LString(combine), value)
	return 0
}

func redisApiExpire(L *lua.LState) int {
	key := L.CheckString(1)
	value := L.CheckNumber(2)
	combine := stringutils.JoinWithUnderline(constants.ExpireAction, key)

	result := L.GetGlobal(constants.LuaGlobalVariableResult)
	L.SetTable(result, lua.LString(combine), value)
	return 0
}
