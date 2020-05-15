package g2

import (
	"errors"
	"log"

	socketio "github.com/googollee/go-socket.io"
	sinterface "github.com/stacew/io/sInterface"
)

const (
	maxRoomCounter = 10
)

type socket struct {
	conCount    int64
	roomNo      int
	roomCounter int
}

func (m *socket) StartUp() {
	m.conCount = 0
	m.roomNo = 0
	m.roomCounter = 0
}
func (m *going) PressKey(c socketio.Conn) {
}
func (m *socket) MakeCustomSocket(s *socketio.Server, nsp string) {
	s.OnEvent(nsp, "cInputSpace", func(c socketio.Conn) {
		log.Println(nsp, "[TEST]", "cInputSpace:"+c.ID())
	})
}

func (m *socket) MakeRoomSocket(s *socketio.Server, nsp string) {

	s.OnEvent(nsp, "cLeave", func(c socketio.Conn) {

	})
}

func (m *socket) MakeBaseSocket(s *socketio.Server, nsp string) {
	s.OnConnect(nsp, func(c socketio.Conn) error {
		if c == nil {
			return errors.New("[ConnNil] connect nil : " + nsp)
		}
		return nil
	})

	s.OnError(nsp, func(c socketio.Conn, e error) {
		if c != nil {
			roomName := sinterface.LeaveGoingRoom(nsp, c)
			c.Leave(roomName)
			return
		}

		log.Println("[ConnNil]", e, nsp)
	})

	s.OnDisconnect(nsp, func(c socketio.Conn, reason string) {
		if c != nil {
			roomName := sinterface.LeaveGoingRoom(nsp, c)
			c.Leave(roomName)
			return
		}

		log.Println("[ConnNil]", reason, nsp)
	})
}
