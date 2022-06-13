package igin

import "fmt"

// BlockingGo will block until it failed
func BlockingGo() error {
	//fmt.Println("setup log...")
	//setupLog()
	//fmt.Println("log setup...")

	r := SetupRouter()
	fmt.Println("router done")
	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
