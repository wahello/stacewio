package go1

import (
	"errors"
	"log"
	sinterface "stacew/stacewio/sInterface"

	socketio "github.com/googollee/go-socket.io"
)

//NewSocket is
func NewSocket() *socket {
	return &socket{}
}

type socket struct {
}

func (m *socket) MakeCustomSocket(s *socketio.Server, nsp string) {
	s.OnEvent(nsp, "cPressKey", func(c socketio.Conn) {
		sinterface.PressKey(nsp, c)
	})
}

func (m *socket) MakeRoomSocket(s *socketio.Server, nsp string) {
	// s.OnEvent(nsp, "cLeave", func(c socketio.Conn) {
	// 	if c != nil {
	// 		roomName := sinterface.LeaveGoingRoom(nsp, c)
	// 		c.Leave(roomName)
	// 		return
	// 	}

	// 	log.Println("[ConnNil]", nsp)
	// })
}

func (m *socket) MakeBaseSocket(s *socketio.Server, nsp string) {
	s.OnConnect(nsp, func(c socketio.Conn) error {
		roomName := sinterface.JoinGoingRoom(nsp, c)
		c.Join(roomName)
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
