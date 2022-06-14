package spy_and_family

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/FPNL/anime/package/anime"
	"github.com/streadway/amqp"
)

type Twilight struct {
	Hint   string
	Person string
}

func ep1(user anime.User) any {
	// data assertion
	d, ok := user.Data().(*amqp.Delivery)
	if !ok {
		return user.StopFor(errors.New("data is not from Twilight"))
	}

	var Thorn_Princess Twilight
	err := json.Unmarshal(d.Body, &Thorn_Princess)
	if err != nil {
		return user.StopFor(err)
	}

	if Thorn_Princess.Hint != "Strix" {
		fmt.Printf("This is Thorn_Princess got message from %s\n", Thorn_Princess.Hint)

		return user.SkipTo(2, Thorn_Princess) // what is five?
	}

	fmt.Printf("This is Thorn_Princess got message from Strix, target is %s\n", Thorn_Princess.Person)

	return Thorn_Princess

	// how do episode know if it is done?
	//return d
}
