package igin

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	api := r.Group("/api")

	deliveryPing := di_delivery_ping()
	{
		api.GET("/ping", deliveryPing.PinPong)
	}

	adminGroup := api.Group("/v1/user")
	deliveryAdmin := di_delivery_admin()
	{
		adminGroup.POST("/login", deliveryAdmin.Login)
		adminGroup.POST("/register", deliveryAdmin.Register)
	}

	wordApi := api.Group("/v1/word")
	wordApi.Use(deliveryAdmin.Guard_authenticate)
	deliveryWord := di_delivery_word()
	{
		//AddWords
		wordApi.POST("/add/committed", deliveryWord.AddCommittedWords)
		//FetchAdvisedWords
		wordApi.POST("/fetch/committed", deliveryWord.FetchCommittedWords)
		//UpdateWords
		wordApi.PUT("/update/Committed", deliveryWord.UpdateCommittedWords)
		////DeleteWords
		//wordApi.DELETE("/delete/committed", deliveryWord.DeleteCommittedWords)
		//
		////AdviseWords
		//wordApi.POST("/advise", deliveryWord.AdviseWords)
		////FetchStageWords
		//wordApi.POST("/fetch/stage", deliveryWord.FetchStageWords)
		////TODO 當 commit word 與 advice word 不一樣，是否為他只承認一部分的意思？
		//// 這樣實際行為為，我承認了你全部，但是我要修改，於事 who advice 最後會被消失
		////CommitWords
		//wordApi.PUT("/commit/stage", deliveryWord.CommitStageWords)
		////DiscardStageWords
		//wordApi.DELETE("/discard/stage", deliveryWord.DiscardStageWords)
	}
}
