package main

import (
	"github.com/FPNL/i18n-town/src/lib/idatabase"
	"github.com/FPNL/i18n-town/src/lib/igin"
)

func main() {
	err := idatabase.Go(idatabase.Option{})
	if err != nil {
		panic(err)
	}

	err = igin.BlockingGo()
	if err != nil {
		panic(err)
	}
}
