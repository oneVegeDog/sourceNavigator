package dal

import (
	"github.com/oneVegeDog/sourceNavigator/internal"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GLOBAL_DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(mysql.Open(internal.GLOBAL_CONFIG.GetString("sourceNavigator.mysql.dns")), &gorm.Config{})
	if err != nil {
		println(err)
		return
	}
	GLOBAL_DB = db
}
