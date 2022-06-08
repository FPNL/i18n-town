package igin

import (
	"fmt"
)

func BlockingGo() error {
	//fmt.Println("setup log...")
	//setupLog()
	//fmt.Println("log setup...")

	fmt.Println("setup router...")
	r := SetupRouter()
	fmt.Println("router setup...")
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	return r.Run()
}
