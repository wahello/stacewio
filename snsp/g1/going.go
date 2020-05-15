package g1

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	socketio "github.com/googollee/go-socket.io"
)

const (
	maxBlockCount      = 40
	canvasX            = 1000
	canvasY            = 1000
	minRoomPlayerCount = 8
	maxRoomPlayerCount = 16
)

//blockType
const (
	blockTypeNormal = iota
	blockTypeSlow
	blockTypeFast
	blockTypeBig
	blockTypeBigSlow
)

type block struct {
	x, y      int //0~1000
	blockType int
}
type player struct {
	x      int //0 ~ 1000
	height int //20~250px
	speed  int //1~3 px per frame
	c      socketio.Conn
}

type roomInfo struct {
	blockArr  [maxBlockCount]block
	playerMap map[string]*player //c.ID()
	roomName  string
}

type going struct {
	frame60   int
	roomSlice []*roomInfo
}

func (m *going) StartUp(frame60 int) {
	m.frame60 = frame60
}

func (m *going) PressKey(c socketio.Conn) {
	if c.Context() == nil {
		return
	}

	roomSiceLen := len(m.roomSlice)
	roomIndex := c.Context().(int)
	if roomIndex >= roomSiceLen {
		return
	}

	currentPlayer := m.roomSlice[roomIndex].playerMap[c.ID()]
	if currentPlayer != nil {
		currentPlayer.speed = -currentPlayer.speed
	}
}

func initBlock(blockArr *[maxBlockCount]block) {
	for i := 0; i < maxBlockCount; i++ {
		blockArr[i].x = rand.Intn(canvasX)
		blockArr[i].y = canvasY + rand.Intn(canvasY*3) //최초에만 시차
		blockArr[i].blockType = rand.Intn(5)           //0,1,2,3,4 //normal, slow, fast, big, bigslow
	}
}

func newRoom(roomName string) *roomInfo {
	room := &roomInfo{playerMap: make(map[string]*player), roomName: roomName}
	initBlock(&room.blockArr)
	return room
}

func (m *going) jjoinRoom(room *roomInfo, i int, c socketio.Conn) {
	c.SetContext(i)
	room.playerMap[c.ID()] = &player{x: 0, height: 20, speed: 1, c: c} //height 20 no die
}
func (m *going) JoinRoom(c socketio.Conn) string {
	nMinIndex := -1
	nMinCount := maxRoomPlayerCount
	//최소 인원도 없는 방 앞에서부터 찾으면 바로 조인
	for i, room := range m.roomSlice {
		roomLen := len(room.playerMap)
		if roomLen < minRoomPlayerCount {
			m.jjoinRoom(room, i, c)
			return room.roomName
		}

		if roomLen < nMinCount {
			nMinCount = roomLen
			nMinIndex = i
		}
	}

	//새로 만들어야 하는 경우.
	if nMinIndex == -1 {
		newIndex := len(m.roomSlice)
		roomName := strconv.Itoa(newIndex)
		room := newRoom(roomName)
		m.roomSlice = append(m.roomSlice, room)
		m.jjoinRoom(room, newIndex, c)
		return room.roomName
	}

	//최소 인원 이상인 경우, 적은 방 넣는 경우.
	room := m.roomSlice[nMinIndex]
	m.jjoinRoom(room, nMinIndex, c)
	return room.roomName
}

func (m *going) LeaveRoom(c socketio.Conn) string {

	if c.Context() == nil { //예외처리: 만약 js에서 join OnEvent 추가 액션으로 들어오면, js를 안 읽는지 context가 nil인 경우가 생김.
		log.Println("[Context Nil] LeavRoom 1")
		return ""
	}

	roomSiceLen := len(m.roomSlice)
	roomIndex := c.Context().(int)
	if roomIndex >= roomSiceLen { //예외처리: Join없이 Leave 두 번 타는 경우 있음.
		log.Println("[Context Nil] LeavRoom 2")
		return ""
	}

	//추가 bug : 서버 print resume걸리면서 소켓이 leave가 안 되는 경우가 존재.
	//전체 룸 순회하면서 정리해주는 함수 작성 하려다가 일단은 1분 정도 지나면 leave 호출되는 것 봐서 skip

	//테스트 기록 : 브라우저 켜놓고 놔둘 시, 5분간 무사고. 다음에 로그로 켜놓고 추가 확인하기.
	roomName := strconv.Itoa(roomIndex)
	delete(m.roomSlice[roomIndex].playerMap, c.ID())
	if len(m.roomSlice[roomIndex].playerMap) == 0 {
		m.roomSlice = append(m.roomSlice[:roomIndex], m.roomSlice[roomIndex+1:]...) //delete center index
	}

	return roomName
}

func (m *going) GOGOGO(s *socketio.Server, nsp string) {
	for {

		for _, roomName := range s.Rooms(nsp) {
			if s.RoomLen(nsp, roomName) != 0 {
				continue
			}
			log.Println("[ROOM is Empty]") //걸린적 없음, 걸리면 게임 룸도 클리어
			for i, room := range m.roomSlice {
				if room.roomName != roomName {
					continue
				}
				m.roomSlice = append(m.roomSlice[:i], m.roomSlice[i+1:]...)
				break
			}
		}

		//Play
		for _, room := range m.roomSlice {

			//블록 이동
			for i := 0; i < maxBlockCount; i++ {
				speed := 2
				block := &room.blockArr[i]
				if block.blockType == blockTypeSlow ||
					block.blockType == blockTypeBigSlow {
					speed = 1
				} else if block.blockType == blockTypeFast {
					speed = 3
				}
				block.y = block.y - speed

				if block.y <= 0 {
					block.y = canvasY
					block.x = rand.Intn(canvasX)
				}
			}

			//유저 이동
			for _, user := range room.playerMap {
				user.x = user.x + user.speed
				if user.x <= 0 {
					user.x = 0
				} else if user.x >= canvasX {
					user.x = canvasX
				}
			}
			//충돌 체크

			msg := ""
			//블록어레이를 문자열로 만들어주는 함수
			for _, block := range room.blockArr {
				msg = msg + ".b" + strconv.Itoa(block.x) + "," + strconv.Itoa(block.y) + "," + strconv.Itoa(block.blockType) + ","
			}

			//유저정볼르 문자열로 만들어주는함수
			//frame 갱신
			s.BroadcastToRoom(nsp, room.roomName, "sFrame", msg)
		}

		time.Sleep(time.Duration(m.frame60) * 16 * time.Millisecond)
	}
}
