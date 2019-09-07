package service

import (
	"github.com/gin-gonic/gin"
	"jiyue.im/model"
	"jiyue.im/pkg/errno"
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

	d, err := model.GetLatest()
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	UID := GetUid(c)
	ret, notFound := getClassicData(UID, d.ArtId, d.Type)
	if notFound {
		SendResponse(c, errno.ErrNotFound, nil)
		return
	}
	ret["index"] = d.Index
	SendResponse(c, errno.OK, ret)
}

// @Summary 获取当前期刊的下一个期刊
// @Tags Classic
// @Accept  json
// @Produce  json
// @Param x-auth-token header string  true "x-auth-token"
// @Param id path int32 true "Id"
// @Success 200 {object} service.Response "{}"
// @Router /v1/classic/next/{id} [get]
func ClassicNext(c *gin.Context) {
	var next model.ParamIndexPre
	if err := c.ShouldBindUri(&next); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	next.Id += 1
	d, isFound := model.FlowOne(next.Id)
	if isFound {
		SendResponse(c, errno.ErrNotFound, nil)
		return
	}
	UID := GetUid(c)
	ret, notFound := getClassicData(UID, d.ArtId, d.Type)
	if notFound {
		SendResponse(c, errno.ErrNotFound, nil)
		return
	}
	ret["index"] = d.Index
	SendResponse(c, errno.OK, ret)
}

// @Summary 获取当前期刊的上一个期刊
// @Tags Classic
// @Accept  json
// @Produce  json
// @Param x-auth-token header string  true "x-auth-token"
// @Param id path int32 true "Id"
// @Success 200 {object} service.Response "{}"
// @Router /v1/classic/previous/{id} [get]
func ClassicPrevious(c *gin.Context) {
	var next model.ParamIndexPre
	if err := c.ShouldBindUri(&next); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	next.Id -= 1
	d, isFound := model.FlowOne(next.Id)
	if isFound {
		SendResponse(c, errno.ErrNotFound, nil)
		return
	}
	UID := GetUid(c)
	ret, notFound := getClassicData(UID, d.ArtId, d.Type)
	if notFound {
		SendResponse(c, errno.ErrNotFound, nil)
		return
	}
	ret["index"] = d.Index
	SendResponse(c, errno.OK, ret)
}

// @Summary 用户是否喜欢一个期刊
// @Tags Classic
// @Accept  json
// @Produce  json
// @Param x-auth-token header string  true "x-auth-token"
// @Param type path int32 true "Type"
// @Param id path int32 true "Id"
// @Success 200 {object} service.Response "[{"id": 1,"image": "images/music.7.png","content": "无人问我粥可温 风雨不见江湖人", "title": "月之门《枫华谷的枫林》","fav_nums": 145,"type": 200,"url": "http://music.163.com/song/media/outer/url?id=393926.mp3"}]"
// @Router /v1/classic/favor/{type}/{id} [get]
func ClassicFavor(c *gin.Context) {
	var paramFavor model.ParamFavorReq
	if err := c.ShouldBindUri(&paramFavor); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	data, isNotFound := model.GetData(paramFavor.Id, paramFavor.Type)

	var v interface{} = data
	var num int32
	switch v.(type) {
	case *model.MovieModel:
		num = v.(*model.MovieModel).FavNums
	case *model.MusicModel:
		num = v.(*model.MusicModel).FavNums
	case *model.SentenceModel:
		num = v.(*model.SentenceModel).FavNums
	}

	if isNotFound {
		SendResponse(c, errno.ErrNotFound, nil)
		return
	}

	UID := GetUid(c)
	isFavor := model.UserLikeIt(paramFavor.Id, paramFavor.Type, UID)
	SendResponse(c, errno.OK, &model.ParamFavorResp{FavNums: num, LikeStatus: isFavor})
}

// @Summary 获取用户喜欢的期刊
// @Tags Classic
// @Accept  json
// @Produce  json
// @Param x-auth-token header string  true "x-auth-token"
// @Success 200 {object} service.Response "{"content": "人生不能像做菜，把所有的料准备好才下锅","favnums": 0,"id": 7,"image": "images/movie.8.png","index": 8,"like_status": true,"pubdate": "2019-04-05T17:12:04+08:00","title": "李安《饮食男女 》","type": 100}"
// @Router /v1/classic/favors [get]
func ClassicFavors(c *gin.Context) {
	UID := GetUid(c)
	art, isNotFound := model.GetClassicFavors(UID)
	if isNotFound {
		SendResponse(c, errno.ErrNotFound, nil)
		return
	}
	SendResponse(c, errno.OK, art["arts"])
}

// @Summary 获取期刊详情
// @Tags Classic
// @Accept  json
// @Produce  json
// @Param x-auth-token header string  true "x-auth-token"
// @Param type path int32 true "Type"
// @Param id path int32 true "Id"
// @Success 200 {object} service.Response "[{"id": 1,"image": "images/music.7.png","content": "无人问我粥可温 风雨不见江湖人", "title": "月之门《枫华谷的枫林》","fav_nums": 145,"type": 200,"url": "http://music.163.com/song/media/outer/url?id=393926.mp3"}]"
// @Router /v1/classic/detail/{type}/{id} [get]
func ClassicDetail(c *gin.Context) {

	var paramFavor model.ParamFavorReq
	if err := c.ShouldBindUri(&paramFavor); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	UID := GetUid(c)
	ret, notFound := getClassicData(UID, paramFavor.Id, paramFavor.Type)
	if notFound {
		SendResponse(c, errno.ErrNotFound, nil)
		return
	}
	SendResponse(c, errno.OK, ret)
}

func getClassicData(uid uint64, art_id, art_type int32) (ret map[string]interface{}, isNotFound bool) {
	data, isNotFound := model.GetData(art_id, art_type)
	if isNotFound {
		return
	}
	isFavor := model.UserLikeIt(art_id, art_type, uid)
	ret = util.StructToMap(data)
	ret["id"] = ret["basemodel"].(model.BaseModel).Id
	ret["fav_num"] = ret["favnums"]
	delete(ret, "basemodel")
	delete(ret, "status")
	delete(ret, "favnums")
	ret["like_status"] = isFavor
	return
}
