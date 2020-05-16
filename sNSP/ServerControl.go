package snsp

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
	sinterface "github.com/stacew/io/sInterface"
	"github.com/stacew/io/sNSP/index"
	"github.com/stacew/io/sNSP/yslee"
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
	sinterface.SetHandler(s, nsp, socketI, goI, 60) //60쓰면 1초에 1번씩 보내기
	//1을 쓰면 60프레임, 2를 쓰면 30프레임

	base60Frame := 1
	//yslee
	nsp, socketI, goI = yslee.GetNewHandler("/go1")
	sinterface.SetHandler(s, nsp, socketI, goI, base60Frame)
	nsp, socketI, goI = yslee.GetNewHandler("/go2")
	sinterface.SetHandler(s, nsp, socketI, goI, base60Frame)

	//someone
	// nsp, socketI, goI = someone.GetNewHandler("/w1")
	// sinterface.SetHandler(s, nsp, socketI, goI, base60Frame)
	// nsp, socketI, goI = someone.GetNewHandler("/w2")
	// sinterface.SetHandler(s, nsp, socketI, goI, base60Frame)
}
