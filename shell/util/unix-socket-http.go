package util

import (
	"net/http"
	"time"

	"github.com/tv42/httpunix"
)

func GetGMOpsClient() *http.Client {
	u := &httpunix.Transport{
		DialTimeout:           100 * time.Millisecond,
		RequestTimeout:        1 * time.Second,
		ResponseHeaderTimeout: 1 * time.Second,
	}
	u.RegisterLocation("gmops", "/var/run/gmops.sock")

	t := &http.Transport{}
	t.RegisterProtocol(httpunix.Scheme, u)

	return &http.Client{Transport: t}
}
