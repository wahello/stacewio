package index

import socketio "github.com/googollee/go-socket.io"

//NewGoing is ...
func NewGoing() *going {
	return &going{}
}

type going struct {
	frame60 int
}

func (m *going) StartUp(frame60 int) {
	m.frame60 = frame60
}
func (m *going) PressKey(c socketio.Conn) {
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
}
