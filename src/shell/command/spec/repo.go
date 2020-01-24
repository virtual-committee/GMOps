package spec

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	gmopsProto "GMOps/src/proto"
	"GMOps/src/util"

	"github.com/golang/protobuf/proto"
)

type Repo struct {
	id               string
	user             string
	OriginalRepoName string
}

func (r *Repo) parseOriginalRepoName() (string, string, error) {
	trimedRepoName := strings.Trim(r.OriginalRepoName, "'")
	repoNames := strings.Split(trimedRepoName, "/")
	if len(repoNames) != 2 {
		return "", "", fmt.Errorf("invalid repo name")
	}
	userOrGroupName := repoNames[0]
	repoName := repoNames[1]
	return userOrGroupName, repoName, nil
}

func (r *Repo) tryPassUserFetchRepoId(user string, repo string) (bool, error) {
	req, err := http.NewRequest("GET", "http+unix://gmops/user/repo/"+repo, nil)
	if err != nil {
		return false, err
	}
	req.Header.Add("GMOps-Username", user)

	resp, err := util.GetGMOpsClient().Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != 200 {
		errResp := gmopsProto.Error{}
		proto.Unmarshal(body, &errResp)
		return false, fmt.Errorf(errResp.Reason)
	}

	resRepo := gmopsProto.Repo{}
	proto.Unmarshal(body, &resRepo)

	r.id = resRepo.Id
	r.user = user

	return true, nil
}

func (r *Repo) Parse() error {
	userOrGroupName, repoName, err := r.parseOriginalRepoName()
	if err != nil {
		return err
	}
	fetched, err := r.tryPassUserFetchRepoId(userOrGroupName, repoName)
	if err != nil {
		return err
	}
	if fetched {
		return nil
	}

	return fmt.Errorf("not supported group yet")
}
