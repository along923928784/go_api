package service

import (
	"net/http"

	"jiyue.im/pkg/errno"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	// always return http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func GetUid(c *gin.Context) (UID uint64) {
	ID, b := c.Get("UID")
	if !b {
		SendResponse(c, errno.ErrForbbiden, nil)
		return
	}
	UID = ID.(uint64)
	return
}
