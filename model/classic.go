package model

type Art interface {
	FindOne(int32) bool
}

type ClassicModel struct {
	BaseModel
	Status  int32  `json:"status" gorm:"column:status" binding:"required"`
	Image   string `json:"image" gorm:"column:image" binding:"required"`
	Content string `json:"content" gorm:"column:content" binding:"required"`
	Title   string `json:"title" gorm:"column:title" binding:"required"`
	FavNums int32  `json:"fav_nums" gorm:"column:fav_nums" binding:"required"`
	Type    int32  `json:"type" gorm:"column:type" binding:"required"`
}

type MusicModel struct {
	ClassicModel
	Url string `json:"url" gorm:"column:url" binding:"required"`
}

func (music *MusicModel) TableName() string {
	return "tb_music"
}

type MovieModel struct {
	ClassicModel
}

func (movie *MovieModel) TableName() string {
	return "tb_movie"
}

type SentenceModel struct {
	ClassicModel
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
