package service

import (
	"github.com/gin-gonic/gin"
	"jiyue.im/pkg/errno"
	"jiyue.im/model"
)

// @Summary 期刊点赞
// @Tags Like
// @Accept  json
// @Produce  json
// @Param x-auth-token header string  true "x-auth-token"
// @Param likeParam body model.ParamFavorReq true "likeParam"
// @Success 200 {object} service.Response "{}"
// @Router /v1/like/  [post]
func Like(c *gin.Context) {
	UID := GetUid(c)
	err := model.ClassicLikeOrDisLike(c, UID, true)
	if err != nil {
		SendResponse(c, errno.OK, err)
		return
	}
	SendResponse(c, errno.OK, nil)
}

// @Summary 期刊取消点赞
// @Tags Like
// @Accept  json
// @Produce  json
// @Param x-auth-token header string  true "x-auth-token"
// @Param likeParam body model.ParamFavorReq true "likeParam"
// @Success 200 {object} service.Response "{}"
// @Router /v1/like/cancel  [post]
func DisLike(c *gin.Context) {
	UID := GetUid(c)
	err :=model.ClassicLikeOrDisLike(c, UID, false)
	if err != nil {
		SendResponse(c, errno.OK, err)
		return
	}
	SendResponse(c, errno.OK, nil)
}


