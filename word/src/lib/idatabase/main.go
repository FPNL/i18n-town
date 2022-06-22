package idatabase

import (
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type Model interface{}

func Go() (err error) {
	//db, err = setupPSql()
	//if err != nil {
	//	fmt.Println("database done")
	//}
	db, err = setupGorm()
	if err != nil {
		log.Fatalln(err)
	}

	return
}

func Connect() *gorm.DB {
	if db == nil {
		log.Fatalln("專案架構級別錯誤")
	}
	return db
}

func Close() {
	if db == nil {
		log.Fatalln("專案架構級別錯誤")
	}
	//db.Close()
}
