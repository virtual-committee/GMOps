package luautil

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	gmopsProto "GMOps/src/proto"
	"GMOps/src/util"
	"github.com/golang/protobuf/proto"
	lua "github.com/yuin/gopher-lua"
)

func luaRegisterGitGlobal(l *lua.LState) {
	meta := l.NewTable()
	l.SetGlobal("git", meta)

	l.SetField(meta, "repo", l.NewFunction(getGitRepo))
}

func getGitRepo(l *lua.LState) int {
	switch l.GetTop() {
	case 1:
		l.Push(lua.LString(os.Getenv("GMOPS_REPO_ID")))
	case 2:
		repoName := l.Get(-1).String()
		l.Pop(-1)
		repoNameArr := strings.Split(repoName, "/")
		var repoId string = ""
		switch len(repoNameArr) {
		case 1:
			repoId = fetchGitRepoId(os.Getenv("GMOPS_USER"), repoNameArr[0])
		case 2:
			repoId = fetchGitRepoId(repoNameArr[1], repoNameArr[0])
		}
		l.Push(lua.LString(repoId))
	default:
		l.Push(lua.LString(""))
	}
	return 1
}

func fetchGitRepoId(user, repo string) string {
	req, err := http.NewRequest("GET", "http+unix://gmops/user/repo/"+repo, nil)
	if err != nil {
		return ""
	}
	req.Header.Add("GMOps-Username", user)

	resp, err := util.GetGMOpsClient().Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	protoRepo := gmopsProto.Repo{}
	proto.Unmarshal(body, &protoRepo)

	return protoRepo.Id
}
