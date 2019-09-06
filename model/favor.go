package model

type FavorModel struct {
	BaseModel
	Uid   int32 `json:"uid" gorm:"column:uid" binding:"required"`
	ArtId int32 `json:"art_id" gorm:"column:art_id" binding:"required"`
	Type  int32 `json:"type" gorm:"column:type" binding:"required"`
}

func (f *FavorModel) TableName() string {
	return "tb_favor"
}

func UserLikeIt(art_id, art_type int32, uid uint64) bool {
	f := FavorModel{}
	notFound := DB.Where(map[string]interface{}{"uid": uid, "art_id": art_id, "type": art_type}).First(&f).RecordNotFound()
	return !notFound
}
