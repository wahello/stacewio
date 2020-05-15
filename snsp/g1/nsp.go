package g1

import sinterface "github.com/stacew/io/sInterface"

//GetNewHandler is
func GetNewHandler() (string, sinterface.SocketInterface, sinterface.GoingInterface) {
	return "/g1", &socket{}, &going{}
}
