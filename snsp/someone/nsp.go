package someone

import (
	sinterface "github.com/stacew/io/sInterface"
)

//GetNewHandler is
func GetNewHandler(sub string) (string, sinterface.SocketInterface, sinterface.GoingInterface) {
	// nsp := "someone" + sub

	// if sub == "/w1" {
	// 	return nsp, w1.NewSocket(), w1.NewGoing()
	// }
	// else if sub == "/w2" {
	// 	return nsp, w2.NewSocket(), w2.NewGoing()
	// }

	return "", nil, nil
}
