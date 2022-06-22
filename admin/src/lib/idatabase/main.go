package idatabase

import (
	"gorm.io/gorm"
)

type IDB *gorm.DB

var db *gorm.DB

func Go() (err error) {
	//db, err = setupPSql()
	db, err = setupGorm()
	return
}

func Connect() *gorm.DB {
	if db == nil {
		panic("專案架構層級錯誤")
	}
	return db
}

func Close() {
	if db == nil {
		panic("專案架構層級錯誤")
	}

	sqlcmd, _ := db.DB()
	sqlcmd.Close()
	//db.Close()
}
