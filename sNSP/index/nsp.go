package index

import sinterface "stacew/stacewio/sInterface"

//GetNewHandler is
func GetNewHandler() (string, sinterface.SocketInterface, sinterface.GoingInterface) {
	return "/", NewSocket(), NewGoing()
}
