package index

import (
	"errors"
	"log"

	socketio "github.com/googollee/go-socket.io"
)

//NewSocket is
func NewSocket() *socket {
	return &socket{conCount: 0}
}

type socket struct {
	conCount uint64
}

func (m *socket) MakeCustomSocket(s *socketio.Server, nsp string) {
	log.Println("Index MakeCustomSocket() is not developed.")
}

func (m *socket) MakeRoomSocket(s *socketio.Server, nsp string) {
	log.Println("Index MakeRoomSocket() is not developed.")
}

func (m *socket) MakeBaseSocket(s *socketio.Server, nsp string) {
	s.OnConnect(nsp, func(c socketio.Conn) error {
		if c == nil {
			return errors.New("[ConnNil] connect nil : " + nsp)
		}
		m.conCount++
		return nil
	})

	s.OnError(nsp, func(c socketio.Conn, e error) {
		if c == nil {
			log.Println("[ConnNil]", e, nsp)
		}
	})

	s.OnDisconnect(nsp, func(c socketio.Conn, reason string) {
		if c == nil {
			log.Println("[ConnNil]", reason, nsp)
		}
	})
}
