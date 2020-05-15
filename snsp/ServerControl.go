package snsp

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
	sinterface "github.com/stacew/io/sInterface"
	"github.com/stacew/io/sNSP/g1"
	"github.com/stacew/io/sNSP/g2"
	"github.com/stacew/io/sNSP/index"
)

//NewSocketioServer is
func NewSocketServer() *socketio.Server {
	s, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatalln(err)
	}
	return s
}

//RegisterHandler is
func RegisterHandler(s *socketio.Server) {
	nsp, socketI, goI := index.GetNewHandler()
	sinterface.SetHandler(s, nsp, socketI, goI, 60) //1초에 1번씩 보내기

	nsp, socketI, goI = g1.GetNewHandler()
	sinterface.SetHandler(s, nsp, socketI, goI, 1) //1을 쓰면 60프레임 2를 쓰면 30 프레임

	nsp, socketI, goI = g2.GetNewHandler()
	sinterface.SetHandler(s, nsp, socketI, goI, 120) //1초에 30번
}
