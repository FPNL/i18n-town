package igin

import (
	"github.com/FPNL/i18n-town/src/core/delivery"
	"github.com/FPNL/i18n-town/src/core/repository"
	"github.com/FPNL/i18n-town/src/core/service"
	"github.com/FPNL/i18n-town/src/lib/icache"
	"github.com/FPNL/i18n-town/src/lib/idatabase"
	"github.com/FPNL/i18n-town/src/lib/igrpc"
	"github.com/FPNL/i18n-town/src/lib/imsgqueue"
)

func di_delivery_word() delivery.IWordDelivery {
	db := idatabase.Connect()
	cache := icache.Connect()
	//mod := model.Word(db)
	repo := repository.Word(db, cache)
	serv := service.Word(repo)
	return delivery.Word(serv)
}

func di_delivery_ping() delivery.IPingDelivery {
	serv := service.Ping(imsgqueue.ConnectChn(), imsgqueue.GetQueue())
	return delivery.Ping(serv)
}

func di_delivery_admin() delivery.IAdminDelivery {
	return delivery.Admin(service.Admin(igrpc.ConnectAdmin()))
}
