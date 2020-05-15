package sinterface

import (
	socketio "github.com/googollee/go-socket.io"
)

//SocketInterface is
type SocketInterface interface {
	StartUp()
	MakeBaseSocket(s *socketio.Server, nsp string)
	MakeRoomSocket(s *socketio.Server, nsp string)
	MakeCustomSocket(s *socketio.Server, nsp string)
}

//GoingInterface is
type GoingInterface interface {
	StartUp(frame int)
	GOGOGO(s *socketio.Server, nsp string)

	JoinRoom(c socketio.Conn) string
	LeaveRoom(c socketio.Conn) string
	PressKey(c socketio.Conn)
}

type nspManager struct {
	goIMap  map[string]GoingInterface
	joinCnt map[string]int
}

var nspMgr *nspManager

func init() {
	nspMgr = &nspManager{}
	nspMgr.goIMap = make(map[string]GoingInterface)
	nspMgr.joinCnt = make(map[string]int)
}

//SetHandler is
func SetHandler(s *socketio.Server, nsp string, socketI SocketInterface, goI GoingInterface, frame60 int) {
	socketI.StartUp()
	socketI.MakeBaseSocket(s, nsp)
	socketI.MakeRoomSocket(s, nsp)
	socketI.MakeCustomSocket(s, nsp)

	nspMgr.goIMap[nsp] = goI
	nspMgr.goIMap[nsp].StartUp(frame60)
	go nspMgr.goIMap[nsp].GOGOGO(s, nsp)
}

//JoinGoingRoom is
func JoinGoingRoom(nsp string, c socketio.Conn) string {
	nspMgr.joinCnt[nsp]++
	return nspMgr.goIMap[nsp].JoinRoom(c)
}

//LeaveGoingRoom is
func LeaveGoingRoom(nsp string, c socketio.Conn) string {
	return nspMgr.goIMap[nsp].LeaveRoom(c)
}

//PressKey is
func PressKey(nsp string, c socketio.Conn) {
	nspMgr.goIMap[nsp].PressKey(c)
}
