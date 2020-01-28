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
	param := strings.Fields(string(input))

	repoId := os.Getenv("GMOPS_REPO_ID")

	if len(param)%3 != 0 {
		os.Exit(1)
	}
	req, err := http.NewRequest("GET", "http+unix://gmops/repo/"+repoId+"/hook/pre-receive", nil)
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

	for i := 0; i < len(param); i += 3 {
		subParam := param[i : i+3]
		for _, hook := range hooks.Hooks {
			if !execLuaSource(hook, subParam[0], subParam[1], subParam[2]) {
				os.Exit(1)
			}
		}
	}

	os.Exit(0)
}

func execLuaSource(hook *gmopsProto.Hook, oldrev, newrev, refname string) bool {
	l := lua.NewState()
	defer l.Close()
	gmopslua.LuaRegisterGlobal(l)

	fmt.Fprintln(os.Stderr, "<pre-receive hook> %s processing", hook.Name)
	if err := l.DoString(hook.Source); err != nil {
		fmt.Fprintln(os.Stderr, "lua source internal error (loading)")
		os.Exit(1)
	}
	if err := l.CallByParam(lua.P{
		Fn:      l.GetGlobal("prereceive"),
		NRet:    1,
		Protect: true,
	}); err != nil {
		fmt.Fprintln(os.Stderr, "lua source internal error (exec)")
		os.Exit(1)
	}

	ret := l.Get(-1)
	l.Pop(-1)

	res, ok := ret.(lua.LBool)
	if !ok {
		fmt.Fprintln(os.Stderr, "<pre-receive hook> %s failed", hook.Name)
		return false
	}
	if res {
		fmt.Fprintln(os.Stderr, "<pre-receive hook> %s successed", hook.Name)
		return true
	} else {
		fmt.Fprintln(os.Stderr, "<pre-receive hook> %s failed", hook.Name)
		return false
	}
}
