package command

import (
	"fmt"
	"net/http"

	"GMOps/shell/util"
)

type Principal struct {
	Username string
	KeyId    string
}

func ParsePrincipal(args []string) (*Principal, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("invalid principal")
	}
	return &Principal{
		Username: args[0],
		KeyId:    args[1],
	}, nil
}

func (p *Principal) Valid() (bool, error) {
	req, err := http.NewRequest("GET", "http+unix://gmops/user/key/"+p.KeyId+"/valid", nil)
	if err != nil {
		return false, nil
	}
	req.Header.Add("GMOps-Username", p.Username)

	resp, err := util.GetGMOpsClient().Do(req)
	if err != nil {
		return false, nil
	}
	return resp.StatusCode == 200, nil
}
