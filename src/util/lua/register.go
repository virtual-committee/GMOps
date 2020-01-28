package luautil

import (
	lua "github.com/yuin/gopher-lua"
)

func LuaRegisterGlobal(l *lua.LState) {
	luaRegisterHttpClientGlobal(l)
	luaRegisterGMOpsGlobal(l)
}
