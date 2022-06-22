package ianime

import (
	"log"

	"github.com/FPNL/anime/core/spy_and_family"
	"github.com/FPNL/anime/package/anime"
)

var ania *anime.Dashboard

func Go() error {
	// dashboard -> series -> episode
	ania = anime.NewDashboard(anime.Options{}) // global Options

	spy_and_family.SetupSpyAndFamily(ania)

	ania.MakeAllRelease() // will make all series release

	return nil
	//data := 1
	//for i := 0; i < 10; i++ {
	//	err := a.Play("Spy and family", data)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
}

func Conn() *anime.Dashboard {
	if ania == nil {
		log.Fatalln("架構錯誤")
	}
	return ania
}
