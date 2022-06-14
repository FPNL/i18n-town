package spy_and_family

import (
	"errors"
	"fmt"
	"github.com/FPNL/anime/package/anime"
)

func ep3(u anime.User) any {
	d, ok := u.Data().(Twilight)
	if !ok {
		return u.StopFor(errors.New("data corrupt"))
	}

	fmt.Printf("destroy data... %#v \n", d)

	return nil
}
