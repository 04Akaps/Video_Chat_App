package router

import (
	"github.com/04Akaps/Video_Chat_App/reposiroty"
	"github.com/04Akaps/Video_Chat_App/service"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type socket struct {
	router  Router
	rooms   *reposiroty.RoomMap
	service *service.Service
}

var (
	userNameToSocketId = make(map[string]string)
	socketIdToUsername = make(map[string]string)
	upgrader           = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

var users = make(map[string][]*websocket.Conn)

// socket.ID 기준으로 어떤 방에 들어있는지
var socketRoom = make(map[*websocket.Conn]string)

// 방의 최대 인원수
const maximum = 2

func socketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(conn.RemoteAddr())

	for {
		var msg struct {
			Event     string      `json:"event"`
			Room      string      `json:"room"`
			SDP       interface{} `json:"sdp"`
			Candidate interface{} `json:"candidate"`
		}

		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}

		switch msg.Event {
		case "join_room":
			// 방이 기존에 생성되어 있다면
			if _, ok := users[msg.Room]; ok {
				// 현재 입장하려는 방에 있는 인원수
				currentRoomLength := len(users[msg.Room])
				if currentRoomLength == maximum {
					// 인원수가 꽉 찼다면 돌아갑니다.
					conn.WriteJSON(map[string]interface{}{
						"event": "room_full",
					})
					continue
				}

				// 여분의 자리가 있다면 해당 방 배열에 추가해줍니다.
				users[msg.Room] = append(users[msg.Room], conn)
			} else {
				// 방이 존재하지 않다면 값을 생성하고 추가해줍니다.
				users[msg.Room] = []*websocket.Conn{conn}
			}
			socketRoom[conn] = msg.Room

			// 입장
			conn.WriteJSON(map[string]interface{}{
				"event": "joined_room",
				"room":  msg.Room,
			})

			// 입장하기 전 해당 방의 다른 유저들이 있는지 확인하고
			// 다른 유저가 있었다면 offer-answer을 위해 알려줍니다.
			others := make([]string, 0)
			for _, c := range users[msg.Room] {
				if c != conn {
					others = append(others, socketRoom[c])
				}
			}
			if len(others) > 0 {
				conn.WriteJSON(map[string]interface{}{
					"event":      "all_users",
					"otherUsers": others,
				})
			}

		case "offer":
			roomName := socketRoom[conn]
			// offer를 전달받고 다른 유저들에게 전달해 줍니다.
			for _, c := range users[roomName] {
				if c != conn {
					c.WriteJSON(map[string]interface{}{
						"event": "getOffer",
						"sdp":   msg.SDP,
					})
				}
			}

		case "answer":
			roomName := socketRoom[conn]
			// answer를 전달받고 방의 다른 유저들에게 전달해 줍니다.
			for _, c := range users[roomName] {
				if c != conn {
					c.WriteJSON(map[string]interface{}{
						"event": "getAnswer",
						"sdp":   msg.SDP,
					})
				}
			}

		case "candidate":
			roomName := socketRoom[conn]
			// candidate를 전달받고 방의 다른 유저들에게 전달해 줍니다.
			for _, c := range users[roomName] {
				if c != conn {
					c.WriteJSON(map[string]interface{}{
						"event":     "getCandidate",
						"candidate": msg.Candidate,
					})
				}
			}

		case "disconnect":
			roomID := socketRoom[conn]

			if conns, ok := users[roomID]; ok {
				for i, c := range conns {
					if c == conn {
						// 방을 나가게 된다면 socketRoom과 users의 정보에서 해당 유저를 지워줍니다.
						users[roomID] = append(conns[:i], conns[i+1:]...)
						if len(users[roomID]) == 0 {
							delete(users, roomID)
							return
						}
						break
					}
				}
			}
			delete(socketRoom, conn)

			for _, c := range users[roomID] {
				c.WriteJSON(map[string]interface{}{
					"event": "user_exit",
					"id":    conn.RemoteAddr().String(),
				})
			}
		}
	}

	defer conn.Close()
}

func test() {
	http.HandleFunc("/socket", socketHandler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
