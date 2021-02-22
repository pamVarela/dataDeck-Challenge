package routes

import (
	"testDataDeck/controllers"
	"github.com/gin-gonic/gin"
)

func UrlFunctions() *gin.Engine{
	
	router := gin.Default()

	router.Use(controllers.Cors())

	router.GET("/getSongsBySong/:song", controllers.GetSongsBySong)
	router.GET("/getSongsByArtist/:artist", controllers.GetSongsByArtist)
	router.GET("/getSongsByGenre/:genre", controllers.GetSongsByGenre)
	router.GET("/getGenre/:name", controllers.GetGenre)
	//Extra functions
	router.GET("/getSongsByLength/min/:min/max/:max", controllers.GetSongsByLength)
	router.GET("/getGenres/", controllers.GetGenres)	

	return router	
}
