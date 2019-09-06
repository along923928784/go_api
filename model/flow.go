package model

type FlowModel struct {
	BaseModel
	Status int32 `json:"status" gorm:"column:status" binding:"required"`
	Index  int32 `json:"index" gorm:"column:index" binding:"required"`
	Type   int32 `json:"type" gorm:"column:type" binding:"required"`
	ArtId  int32 `json:"art_id" gorm:"column:art_id" binding:"required"`
}

func (f *FlowModel) TableName() string {
	return "tb_flow"
}

func GetLatest() (*FlowModel, error) {
	f := &FlowModel{}
	d := DB.Order("index", true).Find(&f)
	// d := DB.Exec("SELECT `id`, `index`, `art_id`, `type` FROM `tb_flow` AS `Flow` ORDER BY `Flow`.`index`").Find(&f)
	return f, d.Error
}
