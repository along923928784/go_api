package service

import (
	"github.com/gin-gonic/gin"
	"jiyue.im/model"
	"jiyue.im/pkg/errno"
	"jiyue.im/pkg/token"
	"jiyue.im/util"
)

// @Summary 获取最新期刊
// @Tags Classic
// @Accept  json
// @Produce  json
// @Param x-auth-token header string  true "x-auth-token"
// @Success 200 {object} service.Response "{"content": "人生不能像做菜，把所有的料准备好才下锅","favnums": 0,"id": 7,"image": "images/movie.8.png","index": 8,"like_status": true,"pubdate": "2019-04-05T17:12:04+08:00","title": "李安《饮食男女 》","type": 100}"
// @Router /v1/classic/latest [get]
func GetLatest(c *gin.Context) {

	ctx, b := c.Get("userContext")
	if !b {
		SendResponse(c, errno.ErrForbbiden, nil)
		return
	}
	ID := ctx.(*token.Context).ID

	d, err := model.GetLatest()
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	data, isNotFound := model.GetData(d.ArtId, d.Type)
	if isNotFound {
		SendResponse(c, errno.ErrNotFound, nil)
		return
	}

	isFavor := model.UserLikeIt(d.ArtId, d.Type, ID)
	ret := util.StructToMap(data)
	delete(ret, "basemodel")
	delete(ret, "status")
	ret["like_status"] = isFavor
	ret["index"] = d.Index
	ret["id"] = d.Id
	ret["pubdate"] = d.CreatedAt
	SendResponse(c, errno.OK, ret)
	// SendResponse(c, errno.OK, data)
}
