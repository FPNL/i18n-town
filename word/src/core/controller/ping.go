package controller

import (
	"net/http"

	"github.com/FPNL/i18n-town/src/core/service"

	"github.com/gin-gonic/gin"
)

type IPingHandler interface {
	PinPong(c *gin.Context)
}

type pingHandler struct {
	pingService service.IPingService
}

var singlePing = pingHandler{}

func Ping(pingService service.IPingService) IPingHandler {
	singlePing.pingService = pingService
	return &singlePing
}

func (hd pingHandler) PinPong(c *gin.Context) {
	res, err := hd.pingService.Pong()
	if err == nil {
		c.String(http.StatusOK, res)
	} else {
		c.String(http.StatusBadRequest, err.Error())
	}
}
