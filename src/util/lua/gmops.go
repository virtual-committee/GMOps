package luautil

import (
	"os"

	lua "github.com/yuin/gopher-lua"
)

func luaRegisterGMOpsGlobal(l *lua.LState) {
	meta := l.NewTable()
	l.SetGlobal("gmops", meta)

	l.SetField(meta, "user", l.NewFunction(getGMOpsUser))
}

func getGMOpsUser(l *lua.LState) int {
	c := l.NewTable()

	l.SetField(c, "name", lua.LString(os.Getenv("GMOPS_USER")))

	l.Push(c)
	return 1
}
