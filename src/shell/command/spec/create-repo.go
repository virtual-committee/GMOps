package spec

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	gmopsProto "GMOps/src/proto"
	"GMOps/src/util"

	"github.com/golang/protobuf/proto"
)

type CreateRepo struct {
	Name string
	user string

	stdout io.Writer
	stderr io.Writer
}

func (c *CreateRepo) Exec() error {
	postJson := fmt.Sprintf("{\"name\": \"%s\", \"descript\":\"none\"}", c.Name)

	req, err := http.NewRequest("POST", "http+unix://gmops/user/repo", bytes.NewBuffer([]byte(postJson)))
	if err != nil {
		return err
	}
	req.Header.Add("GMOps-Username", c.user)
	req.Header.Add("Content-Type", "application/json")

	resp, err := util.GetGMOpsClient().Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode == 201 {
		io.WriteString(c.stdout, fmt.Sprintf("repo: <%s> created", c.Name))
	} else {
		errResp := gmopsProto.Error{}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		proto.Unmarshal(body, &errResp)
		io.WriteString(c.stderr, errResp.Reason)
		os.Exit(1)
	}
	return nil
}

func (c *CreateRepo) SetReadWriter(_ io.Reader, stdout, stderr io.Writer) {
	c.stdout = stdout
	c.stderr = stderr
}

func (c *CreateRepo) RequiredUser() bool { return true }

func (c *CreateRepo) SetUser(user string) {
	c.user = user
}
