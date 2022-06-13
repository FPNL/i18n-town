package main

import (
	"github.com/FPNL/i18n-town/src/lib/icache"
	"github.com/FPNL/i18n-town/src/lib/idatabase"
	"github.com/FPNL/i18n-town/src/lib/igin"
	"github.com/FPNL/i18n-town/src/lib/igrpc"
	"github.com/FPNL/i18n-town/src/setting"
)

func main() {
	var err error

	err = setting.Go()
	if err != nil {
		panic(err)
	}

	err = idatabase.Go()
	if err != nil {
		panic(err)
	}
	defer idatabase.Close()

	err = icache.Go()
	if err != nil {
		panic(err)
	}
	defer icache.Close()

	c, err := igrpc.Go()
	if err != nil {
		panic("專案架構錯誤")
	}
	defer c.Close()

	err = igin.BlockingGo()
	if err != nil {
		panic(err)
	}
}
