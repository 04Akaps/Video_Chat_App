package router

import (
	"github.com/04Akaps/Video_Chat_App/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Router) JoinRoom(c *gin.Context) {

	var req types.JoinRoomReq

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.NewRespHeader(http.StatusBadRequest, "Bind Error : ", err.Error()))
		return
	}

}
