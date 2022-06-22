package main

import (
	"log"

	"github.com/FPNL/anime/lib/ianime"
	"github.com/FPNL/anime/lib/imsgqueue"
	"github.com/FPNL/anime/src/setting"
)

func main() {
	var err error
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	err = setting.Go()
	if err != nil {
		log.Fatalf("setting 專案架構錯誤: %v", err)
	}

	err = ianime.Go()
	if err != nil {
		log.Fatalln(err)
	}

	err = imsgqueue.Go()
	if err != nil {
		log.Fatalln(err)
	}
}
