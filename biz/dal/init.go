package dal

import (
	"tiktok.service/biz/dal/mysql"
)

func Init() {
	mysql.Init()
}
