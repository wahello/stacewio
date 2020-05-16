package index

import sinterface "github.com/stacew/io/sInterface"

//GetNewHandler is
func GetNewHandler() (string, sinterface.SocketInterface, sinterface.GoingInterface) {
	return "/", NewSocket(), NewGoing()
}
