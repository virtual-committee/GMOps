package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"

	gmopsProto "GMOps/src/proto"
	"GMOps/src/util"
	gmopslua "GMOps/src/util/lua"

	"github.com/golang/protobuf/proto"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	input, _ := ioutil.ReadAll(os.Stdin)
	params := strings.Fields(string(input))

	repoId := os.Getenv("GMOPS_REPO_ID")

	if len(params)%3 != 0 {
		os.Exit(1)
	}
	req, err := http.NewRequest("GET", "http+unix://gmops/repo/"+repoId+"/hook/post-receive", nil)
	if err != nil {
		os.Exit(1)
	}
	resp, err := util.GetGMOpsClient().Do(req)
	if err != nil {
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		os.Exit(1)
	}
	hooks := gmopsProto.Hooks{}
	proto.Unmarshal(body, &hooks)

	for _, hook := range hooks.Hooks {
		if !execLuaSource(hook, params) {
			os.Exit(1)
		}
	}

	os.Exit(0)
}

func execLuaSource(hook *gmopsProto.Hook, params []string) bool {
	l := lua.NewState()
	defer l.Close()
	gmopslua.LuaRegisterGlobal(l)

	luaParams := l.NewTable()
	for i := 0; i < len(params); i += 3 {
		luaParam := l.NewTable()
		l.SetField(luaParam, "oldrev", lua.LString(params[i]))
		l.SetField(luaParam, "newrev", lua.LString(params[i+1]))
		l.SetField(luaParam, "refname", lua.LString(params[i+2]))
		luaParams.Append(luaParam)
	}

	fmt.Fprintln(os.Stderr, "<post-receive hook> %s processing", hook.Name)
	if err := l.DoString(hook.Source); err != nil {
		fmt.Fprintln(os.Stderr, "lua source internal error (loading)")
		os.Exit(1)
	}
	if err := l.CallByParam(lua.P{
		Fn:      l.GetGlobal("hook"),
		NRet:    1,
		Protect: true,
	}, luaParams); err != nil {
		fmt.Fprintln(os.Stderr, "lua source internal error (exec)")
		os.Exit(1)
	}

	ret := l.Get(-1)
	l.Pop(-1)

	res, ok := ret.(lua.LBool)
	if !ok {
		fmt.Fprintln(os.Stderr, "<post-receive hook> %s failed", hook.Name)
		return false
	}
	if res {
		fmt.Fprintln(os.Stderr, "<post-receive hook> %s successed", hook.Name)
		return true
	} else {
		fmt.Fprintln(os.Stderr, "<post-receive hook> %s failed", hook.Name)
		return false
	}
}
