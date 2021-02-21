package mappings

import (
	"testDataDeck/controllers"
	"github.com/gin-gonic/gin"
)


func UrlFunctions(router *gin.Engine){

	router.Use(controllers.Cors())

	router.GET("/getSong/:params", controllers.GetSong)
	router.GET("/getGenres", controllers.GetGenres)
}