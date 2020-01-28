package util

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"

	lua "github.com/yuin/gopher-lua"
)

func LuaRegisterHttpClientGlobal(l *lua.LState) {
	meta := l.NewTable()
	l.SetGlobal("http", meta)

	l.SetField(meta, "req", l.NewFunction(newGMOpsHttpReq))
	l.SetField(meta, "send", l.NewFunction(sendGMOpsHttpReq))
}

func newGMOpsHttpReq(l *lua.LState) int {
	c := l.NewTable()
	method := lua.LString(l.CheckString(2))
	url := lua.LString(l.CheckString(3))

	l.SetField(c, "method", method)
	l.SetField(c, "url", url)
	l.SetField(c, "header", l.NewTable())
	l.SetField(c, "body", lua.LString(""))

	l.Push(c)
	return 1
}

func sendGMOpsHttpReq(l *lua.LState) int {
	self := l.CheckTable(2)

	method := l.GetField(self, "method").(lua.LString)
	url := l.GetField(self, "url")
	body := l.GetField(self, "body").(lua.LString)

	req, err := http.NewRequest(method.String(), url.String(), bytes.NewBuffer([]byte(body.String())))
	if err != nil {
		luaResp := newGMOpsHttpRes(l, 500, "new request error")
		l.Push(luaResp)
		return 1
	}
	l.GetField(self, "header").(*lua.LTable).ForEach(func(key, value lua.LValue) {
		req.Header.Add(key.String(), value.String())
	})
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		luaResp := newGMOpsHttpRes(l, 500, "new request error")
		l.Push(luaResp)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		luaResp := newGMOpsHttpRes(l, 500, "new request error")
		l.Push(luaResp)
		return 1
	}
	luaResp := newGMOpsHttpRes(l, resp.StatusCode, string(respBody))
	l.Push(luaResp)

	return 1
}

func newGMOpsHttpRes(l *lua.LState, statusCode int, body string) *lua.LTable {
	c := l.NewTable()

	l.SetField(c, "statusCode", lua.LNumber(statusCode))
	l.SetField(c, "body", lua.LString(body))

	return c
}
