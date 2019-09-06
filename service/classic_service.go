package service

import (
	"github.com/gin-gonic/gin"
	"jiyue.im/model"
	"jiyue.im/pkg/errno"
	"jiyue.im/pkg/token"
	"jiyue.im/util"
)

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
