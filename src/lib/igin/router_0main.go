package igin

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", pinPong)

	word := r.Group("/word")
	{
		word.GET("/all", fetchAllWords)
		word.POST("/addOne", addOneWord)
		word.POST("/addMany", addManyWords)
		word.PUT("/updateOne", updateOneWord)
		word.PUT("/updateMany", updateManyWords)
		word.DELETE("/deleteOne", deleteOneWord)
		word.DELETE("/deleteMany", deleteManyWord)
		word.DELETE("/deleteAll", deleteAll)
	}

	return r
}
