package conf

import (
	"github.com/joho/godotenv"
	"strings"
	"wkb_comments/model"
	"wkb_comments/src/logging"
)

const (
	userName = "root"
	password = "mirror"
	ip       = "127.0.0.1"
	port     = "4036"
	dbName   = "wkb_comments"
)

func Init() {
	_ = godotenv.Load() //从本地读取环境变量
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		logging.Info(err) //日志内容
		panic(err)
	}
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8&parseTime=true"}, "")
	model.Database(path)
}
