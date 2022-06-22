package main

import (
	"github.com/FPNL/i18n-town/src/lib/icache"
	"github.com/FPNL/i18n-town/src/lib/idatabase"
	"github.com/FPNL/i18n-town/src/lib/igin"
	"github.com/FPNL/i18n-town/src/lib/igrpc"
	"github.com/FPNL/i18n-town/src/lib/imsgqueue"
	"github.com/FPNL/i18n-town/src/setting"
	"log"
)

func main() {
	var err error
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	err = setting.Go()
	if err != nil {
		log.Fatalf("setting 專案架構錯誤: %v", err)
	}

	err = idatabase.Go()
	if err != nil {
		log.Fatalf("idatabase 專案架構錯誤: %v", err)
	}
	defer idatabase.Close()

	err = icache.Go()
	if err != nil {
		log.Fatalf("icache 專案架構錯誤: %v", err)
	}
	defer icache.Close()

	err = igrpc.Go()
	if err != nil {
		log.Fatalf("igrpc 專案架構錯誤: %v", err)
	}
	defer igrpc.Close()

	err = imsgqueue.Go()
	if err != nil {
		log.Fatalf("imsgqueue 專案架構錯誤: %v", err)
	}
	defer imsgqueue.Close()

	err = igin.BlockingGo()
	if err != nil {
		log.Fatalf("igin 專案架構錯誤: %v", err)
	}
}
