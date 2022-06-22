package spy_and_family

import (
	"errors"
	"fmt"

	"github.com/FPNL/anime/package/anime"
)

func ep2(u anime.User) any {
	d, ok := u.Data().(Twilight)
	if !ok {
		return u.StopFor(errors.New("data is not from Westalis"))
	}

	fmt.Printf("Protecting target %s ... \n", d.Person)

	return d
}
