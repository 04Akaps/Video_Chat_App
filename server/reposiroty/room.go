package reposiroty

import (
	"github.com/gorilla/websocket"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Participant struct {
	Users    []User
	RoomName string
}

type User struct {
	Socket      *websocket.Conn
	Name        string
	SendMessage chan *Message
}

type Message struct {
	Sender  string
	Message string
	To      string
	Time    time.Time
}

type RoomMap struct {
	Mutex            sync.RWMutex
	Map              map[string]*Participant
	ForwardChannel   chan *Message
	JoinUserChannel  chan *User
	LeaveUserChannel chan *User
	OwnerMap         map[string]string
	upgrader         websocket.Upgrader
}

func NewAllRooms() *RoomMap {
	return &RoomMap{
		Mutex:    sync.RWMutex{},
		Map:      make(map[string]*Participant),
		OwnerMap: make(map[string]string),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 256,
		},
	}
}

func (r *RoomMap) CreateRoom(user string) string {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, 8)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	roomID := string(b)

	r.Map[roomID] = &Participant{}
	r.OwnerMap[roomID] = user

	return roomID
}
