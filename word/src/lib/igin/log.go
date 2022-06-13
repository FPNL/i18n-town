package igin

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func setupLog() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
}
