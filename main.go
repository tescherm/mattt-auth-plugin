package main

import (
	"os"
	"strconv"

	"github.com/Sirupsen/logrus"
	authz "github.com/docker/go-plugins-helpers/authorization"
)

const (
	socketAddress = "/run/docker/plugins/mattt.sock"
)

type authzPlugin struct {
	authz.Plugin
}

func (p *authzPlugin) AuthZReq(r authz.Request) authz.Response {
	return authz.Response{
		Allow: false,
		Msg:   "You are not authorized",
		Err:   "",
	}
}

func (p *authzPlugin) AuthZRes(r authz.Request) authz.Response {
	return authz.Response{
		Allow: false,
		Msg:   "You are not authorized",
		Err:   "",
	}
}

func main() {
	debug := os.Getenv("DEBUG")
	if ok, _ := strconv.ParseBool(debug); ok {
		logrus.SetLevel(logrus.DebugLevel)
	}

	d := &authzPlugin{}
	h := authz.NewHandler(d)
	logrus.Infof("listening on %s", socketAddress)
	logrus.Error(h.ServeUnix("", socketAddress))
}
