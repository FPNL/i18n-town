package igin

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
)

func SetupLog(r *gin.Engine) {
	r.Use(ginBodyLogMiddleware)
}

type bodyLogWriter struct {
	gin.ResponseWriter
	out  io.Writer
	body *bytes.Buffer
}

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

var f, _ = os.OpenFile("gin.log", os.O_RDWR|os.O_CREATE, 0755)

func ginBodyLogMiddleware(c *gin.Context) {
	//defer f.Close()
	blw := &bodyLogWriter{
		body:           bytes.NewBufferString(""),
		ResponseWriter: c.Writer,
		out:            f,
	}
	c.Writer = blw
	c.Next()

	blw.reqWrite(c)
	blw.resWrite(c)
}

func (blw *bodyLogWriter) resWrite(c *gin.Context) {
	_, err := fmt.Fprintf(blw.out, " Res[%v]: %s\n", c.Writer.Status(), blw.body.String())
	if err != nil {
		log.Println(err)
		return
	}
}

func (blw *bodyLogWriter) reqWrite(c *gin.Context) {
	// Start timer
	start := time.Now()
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery

	// Process request
	c.Next()

	// Log only when path is not being skipped
	param := gin.LogFormatterParams{
		Request: c.Request,
		Keys:    c.Keys,
	}

	// Stop timer
	param.TimeStamp = time.Now()
	param.Latency = param.TimeStamp.Sub(start)

	param.ClientIP = c.ClientIP()
	param.Method = c.Request.Method
	param.StatusCode = c.Writer.Status()
	param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()

	param.BodySize = c.Writer.Size()

	if raw != "" {
		path = path + "?" + raw
	}

	param.Path = path

	_, err := fmt.Fprintf(blw.out, `%s - [%s] "%s %s %s %d %s" "%s" "%s" `,
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
	if err != nil {
		log.Println(err)
	}
}
