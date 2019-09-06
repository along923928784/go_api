package conf

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lexkong/log"
	"jiyue.im/model"
)

func Init(addr *string) {
	// 从本地读取环境变量
	godotenv.Load()
	*addr = os.Getenv("API_ADDR")
	gin.SetMode(os.Getenv("GIN_MODE"))
	model.Database(os.Getenv("MYSQL_DSN"))
	initLog()
}

func initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        "stdout",
		LoggerLevel:    "DEBUG",
		LoggerFile:     "log/apiserver.log",
		LogFormatText:  true,
		RollingPolicy:  "size",
		LogRotateDate:  1,
		LogRotateSize:  1,
		LogBackupCount: 7,
	}

	log.InitWithConfig(&passLagerCfg)
}
