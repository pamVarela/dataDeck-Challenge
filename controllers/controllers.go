package controllers

import (
	"log"
	"testDataDeck/models"
	
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func GetSong(c *gin.Context) {
	
	var songs []*models.Songs
	rows, err := db.Query("SELECT * FROM songs order by ID")
	log.Println("HERE")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		
		var song = new(models.Songs)
		rows.Scan(&song.Id, &song.Artist, &song.Song, &song.Genre, &song.Length)
		songs = append(songs, song)
		if err != nil {
			c.JSON(200, songs)
		} else {
			c.JSON(404, gin.H{"error": "Song not found"})
		}

	}

}

func GetGenres(c *gin.Context) {
}

