package g2

import (
	"time"

	socketio "github.com/googollee/go-socket.io"
)

type going struct {
	frame60 int
}

func (m *going) StartUp(frame60 int) {
	m.frame60 = frame60
}
func (m *going) JoinRoom(c socketio.Conn) string {
	return ""
}

func (m *going) LeaveRoom(c socketio.Conn) string {
	return ""
}

func (m *going) newRoom(roomName string) {

}

func (m *going) GOGOGO(s *socketio.Server, nsp string) {
	for {
		for _, roomName := range s.Rooms(nsp) {
			if s.RoomLen(nsp, roomName) == 0 {
				s.ClearRoom(nsp, roomName)
				continue
			}

			println(s.RoomLen(nsp, roomName), nsp)
		}

		time.Sleep(time.Duration(m.frame60) * 16 * time.Millisecond)
	}
}
