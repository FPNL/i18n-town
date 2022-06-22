package igin

import (
	"github.com/gin-gonic/gin"
	"os"
)

// BlockingGo will block until it failed
func BlockingGo() error {
	r := gin.Default()

	SetupLog(r)
	SetupRouter(r)

	return r.Run(":" + os.Getenv("IWORD_PORT")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
