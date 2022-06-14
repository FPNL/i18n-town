package igin

import (
	"github.com/FPNL/i18n-town/src/core/controller"
	"github.com/FPNL/i18n-town/src/core/model"
	"github.com/FPNL/i18n-town/src/core/repository"
	"github.com/FPNL/i18n-town/src/core/service"
	"github.com/FPNL/i18n-town/src/lib/icache"
	"github.com/FPNL/i18n-town/src/lib/idatabase"
	"github.com/FPNL/i18n-town/src/lib/igrpc"
	"github.com/FPNL/i18n-town/src/lib/imsgqueue"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")

	{
		serv := service.Ping(imsgqueue.ConnectChn(), imsgqueue.GetQueue())
		hand := controller.Ping(serv)
		api.GET("/ping", hand.PinPong)
	}

	adminHandler := controller.Admin(service.Admin(igrpc.Connect()))
	adminGroup := api.Group("/v1/admin")
	{
		adminGroup.GET("ping", adminHandler.Ping)
		adminGroup.POST("login", adminHandler.Login)
		adminGroup.POST("register", adminHandler.Register)
	}

	wordApi := api.Group("/v1/word")
	{
		db := idatabase.Connect()
		cache := icache.Connect()
		mod := model.Word(db)
		repo := repository.Word(mod, cache)
		serv := service.Word(repo)
		hand := controller.Word(serv)
		wordApi.GET("/all", hand.FetchAllWords)
		wordApi.Use(adminHandler.Auth)
		wordApi.POST("/addOne", hand.AddOneWord)
		wordApi.POST("/addMany", hand.AddManyWords)
		wordApi.PUT("/updateOne", hand.UpdateOneWord)
		wordApi.PUT("/updateMany", hand.UpdateManyWords)
		wordApi.DELETE("/deleteOne", hand.DeleteOneWord)
		wordApi.DELETE("/deleteMany", hand.DeleteManyWord)
		wordApi.DELETE("/deleteAll", hand.DeleteAll)
	}

	return r
}
