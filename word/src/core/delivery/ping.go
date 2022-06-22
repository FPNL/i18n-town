package delivery

import (
	"net/http"

	"github.com/FPNL/i18n-town/src/core/service"

	"github.com/gin-gonic/gin"
)

type IPingDelivery interface {
	PinPong(c *gin.Context)
}

type pingDelivery struct {
	pingService service.IPingService
}

var singlePing = pingDelivery{}

func Ping(pingService service.IPingService) IPingDelivery {
	singlePing.pingService = pingService
	return &singlePing
}

func (hd pingDelivery) PinPong(c *gin.Context) {
	res, err := hd.pingService.Pong()
	if err == nil {
		c.String(http.StatusOK, res)
	} else {
		c.String(http.StatusBadRequest, err.Error())
	}
}
