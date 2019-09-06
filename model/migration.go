package model

//执行数据迁移

func migration() {
	// 自动迁移模式
	DB.AutoMigrate(&UserModel{})
	DB.AutoMigrate(&FlowModel{})
	DB.AutoMigrate(&MovieModel{})
	DB.AutoMigrate(&MusicModel{})
	DB.AutoMigrate(&SentenceModel{})
	DB.AutoMigrate(&FavorModel{})
}
