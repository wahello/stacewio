package g2

import sinterface "github.com/stacew/io/sInterface"

//GetNewHandler is
func GetNewHandler() (string, sinterface.SocketInterface, sinterface.GoingInterface) {
	return "/g2", &socket{}, &going{}
}
