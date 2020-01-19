package options

import (
	"github.com/spf13/pflag"
)

type ServerOption struct {
	BIAddrPath     string
	WIHostPort     string
	MongoConnector string
	Debug          bool
}

func NewServerOption() *ServerOption {
	return &ServerOption{}
}

func (opt *ServerOption) Add(fs *pflag.FlagSet) {
	fs.StringVar(&opt.BIAddrPath, "bipath", "/var/run/gmops.sock", "The Unix-Socket path for BI serve on")
	fs.StringVar(&opt.WIHostPort, "wihostport", "0.0.0.0:8080", "The host/port for WI serve on")
	fs.StringVar(&opt.MongoConnector, "mongo", "127.0.0.1:27017/gmops", "The Mongo connection string")
	fs.BoolVar(&opt.Debug, "debug", true, "debug flag")
}
