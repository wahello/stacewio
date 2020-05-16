package yslee

import (
	"log"

	sinterface "github.com/stacew/io/sInterface"
	"github.com/stacew/io/sNSP/yslee/go1"
	"github.com/stacew/io/sNSP/yslee/go2"
)

//GetNewHandler is
func GetNewHandler(sub string) (string, sinterface.SocketInterface, sinterface.GoingInterface) {
	if sub == "/go1" {
		return "/yslee" + sub, go1.NewSocket(), go1.NewGoing()
	} else if sub == "/go2" {
		return "/yslee" + sub, go2.NewSocket(), go2.NewGoing()
	}

	log.Fatal("/yslee"+sub, "is Not")
	return "", nil, nil
}
