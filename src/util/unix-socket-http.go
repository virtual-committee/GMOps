package util

import (
	"net/http"
	"os"
	"time"

	"github.com/tv42/httpunix"
)

func GetGMOpsClient() *http.Client {
	u := &httpunix.Transport{
		DialTimeout:           100 * time.Millisecond,
		RequestTimeout:        1 * time.Second,
		ResponseHeaderTimeout: 1 * time.Second,
	}
	var unixSocketPath string
	if os.Getenv("GMOPS_BI_UNIX_SOCKET") == "" {
		unixSocketPath = "/var/run/gmops.sock"
	} else {
		unixSocketPath = os.Getenv("GMOPS_BI_UNIX_SOCKET")
	}
	u.RegisterLocation("gmops", unixSocketPath)

	t := &http.Transport{}
	t.RegisterProtocol(httpunix.Scheme, u)

	return &http.Client{Transport: t}
}
