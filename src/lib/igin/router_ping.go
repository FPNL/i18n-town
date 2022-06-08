package igin

import (
	"net/http"

	"github.com/FPNL/i18n-town/src/core/service"

	"github.com/gin-gonic/gin"
)

func pinPong(c *gin.Context) {
	res, err := service.PingService().Pong()
	if err == nil {
		c.String(http.StatusOK, res)
	} else {
		c.String(http.StatusBadRequest, err.Error())
	}
}
