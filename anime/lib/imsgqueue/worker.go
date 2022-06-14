package imsgqueue

import (
	"github.com/FPNL/anime/core/spy_and_family"
	"github.com/FPNL/anime/lib/ianime"
)

func setupWorker() error {
	msgs, err := createCN("hello")
	if err != nil {
		return err
	}

	for {
		select {
		case d := <-msgs:
			a := ianime.Conn()
			a.Play(spy_and_family.Name, &d)
		}
	}
}
