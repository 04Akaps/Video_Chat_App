package router

import (
	"fmt"
	"github.com/04Akaps/Video_Chat_App/reposiroty"
	"github.com/04Akaps/Video_Chat_App/service"
	"github.com/04Akaps/Video_Chat_App/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type room struct {
	router  Router
	rooms   *reposiroty.RoomMap
	service *service.Service
}

func newRoom(r Router, service *service.Service) *room {
	a := &room{
		router:  r,
		rooms:   reposiroty.NewAllRooms(),
		service: service,
	}

	baseUri := "/room"
	//a.router.verifyAuthToken(),
	//a.router.verifyAuthToken(),
	r.engine.GET(baseUri+"/my-room-list", a.router.verifyAuthToken(), a.myRoomList)
	r.engine.GET(baseUri+"/total-room-list", a.totalRoomList)
	r.engine.GET(baseUri+"/room-by-hash/:hash", a.getRoomByHash)
	r.engine.POST(baseUri+"/change-broadcast/:hash", a.router.verifyAuthToken(), a.changeBroadCast)
	r.engine.GET(baseUri+"/recently-created-room-list", a.recentlyCreatedRoom)

	r.engine.POST(baseUri+"/create", a.router.verifyAuthToken(), a.createRoom)

	return a
}

func (r *room) recentlyCreatedRoom(c *gin.Context) {
	if res, err := r.service.Mysql.RecentlyCreatedRoomList(); err != nil {
		msg := fmt.Sprintf("Get GetAllRoomByPaging Failed %s", err.Error())
		c.JSON(http.StatusConflict, msg)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (r *room) changeBroadCast(c *gin.Context) {
	var req types.ChangeRoomStatusByHashReq

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.NewRespHeader(http.StatusBadRequest, "Bind Error : ", err.Error()))
		return
	}

	if userName, err := r.router.extractToken(c); err != nil {
		c.JSON(http.StatusConflict, "Not Auth")
	} else {
		if err := r.service.Mysql.ChangeBroadCastStatus(userName, req.Hash); err != nil {
			msg := fmt.Sprintf("Get GetRoomDataByHash Failed %s", err.Error())
			c.JSON(http.StatusConflict, msg)
		} else {
			c.JSON(http.StatusOK, "Success.jsx")
		}
	}

}

func (r *room) getRoomByHash(c *gin.Context) {
	var req types.GetRoomByHashReq

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.NewRespHeader(http.StatusBadRequest, "Bind Error : ", err.Error()))
		return
	}

	if room, err := r.service.Mysql.GetRoomDataByHash(req.Hash); err != nil {
		msg := fmt.Sprintf("Get GetRoomDataByHash Failed %s", err.Error())
		c.JSON(http.StatusConflict, msg)
	} else {
		c.JSON(http.StatusOK, types.GetRoomByHashRes{Room: room})
	}
}

func (r *room) totalRoomList(c *gin.Context) {
	var req types.TotalRoomListReq

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.NewRespHeader(http.StatusBadRequest, "Bind Error : ", err.Error()))
		return
	}

	if res, err := r.service.Mysql.GetAllRoomByPaging(&req.Paging); err != nil {
		msg := fmt.Sprintf("Get GetAllRoomByPaging Failed %s", err.Error())
		c.JSON(http.StatusConflict, msg)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (r *room) myRoomList(c *gin.Context) {
	if userName, err := r.router.extractToken(c); err != nil {
		c.JSON(http.StatusAccepted, err.Error())
	} else {
		if res, err := r.service.Mysql.GetRoomByOwner(userName); err != nil {
			msg := fmt.Sprintf("Get Room Failed %s", err.Error())
			c.JSON(http.StatusConflict, msg)
		} else {
			c.JSON(http.StatusOK, res)
		}
	}
}

func (r *room) createRoom(c *gin.Context) {
	if userName, err := r.router.extractToken(c); err != nil {
		c.JSON(http.StatusConflict, "Not Auth")
	} else {

		roomID := r.rooms.CreateRoom(userName)

		if count, err := r.service.Mysql.RoomCountByName(userName); err != nil {
			msg := fmt.Sprintf("Create Room Failed %s", err.Error())
			c.JSON(http.StatusConflict, msg)
		} else {
			if count == 3 {
				msg := fmt.Sprintf("Create Room Failed %s", "count limit")
				c.JSON(http.StatusConflict, msg)
			} else {
				if err = r.service.Mysql.InsertRoom(roomID, userName); err != nil {
					msg := fmt.Sprintf("Create Room Failed %s", err.Error())
					c.JSON(http.StatusConflict, msg)
				} else {

					createdRoom := types.Room{
						RoomHash:  roomID,
						OwnerName: userName,
						CreatedAt: time.Now().Unix(),
					}

					c.JSON(http.StatusOK, createdRoom)
				}
			}
		}
	}
}
