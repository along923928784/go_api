package model

import (
	"github.com/gin-gonic/gin"
	"jiyue.im/pkg/errno"
	"github.com/jinzhu/gorm"
)
type FavorModel struct {
	BaseModel
	Uid   uint64 `json:"uid" gorm:"column:uid" binding:"required"`
	ArtId int32  `json:"art_id" gorm:"column:art_id" binding:"required"`
	Type  int32  `json:"type" gorm:"column:type" binding:"required"`
}

func (f *FavorModel) TableName() string {
	return "tb_favor"
}

func UserLikeIt(art_id, art_type int32, uid uint64) bool {
	f := FavorModel{}
	notFound := DB.Where(map[string]interface{}{"uid": uid, "art_id": art_id, "type": art_type}).First(&f).RecordNotFound()
	return !notFound
}

func GetClassicFavors(uid uint64) (art map[string][]interface{}, notFound bool) {

	var favors []*FavorModel
	DB.Where("uid=?", uid).Find(&favors)
	// 查询不到直接返回
	if len(favors) == 0 {
		notFound = true
		return
	}

	var artInfoMap map[int32][]int32
	artInfoMap = make(map[int32][]int32, 5)
	art = make(map[string][]interface{}, 5)
	// 按期刊类型归类 切片形式
	for _, value := range favors {
		artInfoMap[value.Type] = append(artInfoMap[value.Type], value.ArtId)
	}

	// 找到每个期刊类型的ids
	var movies []*MovieModel
	var musics []*MusicModel
	var sentences []*SentenceModel

	for key, ids := range artInfoMap {
		if len(ids) == 0 {
			continue
		}
		switch key {
		case 100:
			DB.Where("id in (?)", ids).Find(&movies)
			if len(movies) != 0 {
				for _, v := range movies {
					art["arts"] = append(art["arts"], v)
				}
			}
		case 200:
			DB.Where("id in (?)", ids).Find(&musics)
			if len(musics) != 0 {
				for _, v := range musics {
					art["arts"] = append(art["arts"], v)
				}
			}
		case 300:
			DB.Where("id in (?)", ids).Find(&sentences)
			if len(sentences) != 0 {
				for _, v := range sentences {
					art["arts"] = append(art["arts"], v)
				}
			}
		case 400:
			break
		default:
			break
		}
	}
	return

}

// like true 点赞 false 取消点赞
func ClassicLikeOrDisLike(c *gin.Context,UID uint64, like bool) (err error) {
	var expr string
	var likeParam ParamFavorReq

	if err := c.Bind(&likeParam); err != nil {
		return errno.ErrBind
	}
	
	data, notFound := GetData(likeParam.Id, likeParam.Type)
	if notFound {
		return errno.ErrNotFound
	}

	tx := DB.Begin()

	f := FavorModel{ArtId: likeParam.Id, Type: likeParam.Type, Uid: UID}
	if like {
		expr = "fav_num + ?"

		isLike := UserLikeIt(likeParam.Id, likeParam.Type, UID)
		if isLike {
			return errno.ErrLike
		}

		if err = tx.Create(&f).Error; err != nil {
			tx.Rollback()
			return errno.ErrCreate
		}

	} else {
		expr = "fav_num - ?"

		isLike := UserLikeIt(likeParam.Id, likeParam.Type, UID)
		if !isLike {
			return errno.ErrDisLike
		}

		if err = tx.Delete(&f).Error; err != nil {
			tx.Rollback()
			return errno.ErrDelete
		}
	}
	switch data.(type) {
	case *MovieModel:
		classic := data.(*MovieModel)
		err = DB.Model(&classic).UpdateColumn("fav_num", gorm.Expr(expr, 1)).Error
	case *MusicModel:
		classic := data.(*MusicModel)
		err = DB.Model(&classic).UpdateColumn("fav_num", gorm.Expr(expr, 1)).Error
	case *SentenceModel:
		classic := data.(*SentenceModel)
		err = DB.Model(&classic).UpdateColumn("fav_num", gorm.Expr(expr, 1)).Error
	default:

	}
	tx.Commit()
	return
}
