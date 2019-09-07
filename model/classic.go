package model

type Art interface {
	FindOne(int32) bool
	GetModel() interface{}
}

type ParamIndexPre struct {
	Id int32 `uri:"id" binding:"required"`
}
type ParamFavorReq struct {
	ParamIndexPre
	Type int32 `uri:"type" binding:"required"`
}

type ParamFavorResp struct {
	FavNums    int32 `json:"fav_nums"`
	LikeStatus bool  `json:"like_status"`
}

type ClassicModel struct {
	BaseModel
	Status  int32  `json:"-" gorm:"column:status" binding:"required"`
	Image   string `json:"image" gorm:"column:image" binding:"required"`
	Content string `json:"content" gorm:"column:content" binding:"required"`
	Title   string `json:"title" gorm:"column:title" binding:"required"`
	FavNums int32  `json:"fav_nums" gorm:"column:fav_num" binding:"required"`
	Type    int32  `json:"type" gorm:"column:type" binding:"required"`
}

type MusicModel struct {
	ClassicModel
	Url string `json:"url" gorm:"column:url" binding:"required"`
}

type MovieModel struct {
	ClassicModel
}

type SentenceModel struct {
	ClassicModel
}

func (music *MusicModel) TableName() string {
	return "tb_music"
}
func (movie *MovieModel) TableName() string {
	return "tb_movie"
}
func (sentence *SentenceModel) TableName() string {
	return "sentence"
}

func (this *MovieModel) FindOne(art_id int32) bool {
	notFound := DB.Where("id = ?", art_id).First(this).RecordNotFound()
	return notFound
}

func (this *MusicModel) FindOne(art_id int32) bool {
	notFound := DB.Where("id = ?", art_id).First(this).RecordNotFound()
	return notFound
}

func (this *SentenceModel) FindOne(art_id int32) bool {
	notFound := DB.Where("id = ?", art_id).First(this).RecordNotFound()
	return notFound
}

func (this *MovieModel) GetModel() interface{} {
	return this
}

func (this *MusicModel) GetModel() interface{} {
	return this
}

func (this *SentenceModel) GetModel() interface{} {
	return this
}
