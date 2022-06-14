package main

import (
	"github.com/FPNL/anime/lib/ianime"
	"github.com/FPNL/anime/lib/imsgqueue"
	"log"
)

func main() {
	var err error
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	err = ianime.Go()
	if err != nil {
		log.Fatalln(err)
	}

	err = imsgqueue.Go()
	if err != nil {
		log.Fatalln(err)
	}
}
