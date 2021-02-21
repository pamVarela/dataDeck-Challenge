package controllers

import (
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
	
	initDB := GetConnectionDB()

	parameter := c.Params.ByName("params")

	var listSongs []models.Songs

	query := `SELECT s.song, s.artist, g.name, s.length FROM songs AS s JOIN genres AS g ON g.id = s.genre WHERE `

	rowsArtists, _ := initDB.Query(query + `s.artist = ?`, parameter)
	rowsSongs, _  := initDB.Query(query + `s.song = ?`, parameter)
	rowsGenres, _ := initDB.Query(query + `g.name = ?`, parameter)

	if rowsArtists != nil{
		for rowsArtists.Next() {
			var res models.Songs
			rowsArtists.Scan(&res.Song, &res.Artist, &res.Genre, &res.Length)
			listSongs = append(listSongs, res)
		}
	}

	if rowsSongs != nil{
		for rowsSongs.Next() {
			var res models.Songs
			rowsSongs.Scan(&res.Song, &res.Artist, &res.Genre, &res.Length)
			listSongs = append(listSongs, res)
		}
	}

	if rowsGenres != nil{
		for rowsGenres.Next() {
			var res models.Songs
			rowsGenres.Scan(&res.Song, &res.Artist, &res.Genre, &res.Length)
			listSongs = append(listSongs, res)
		}
	}

	defer rowsArtists.Close()

	if listSongs != nil {
		c.JSON(200, listSongs)
	} else {
		c.JSON(404, gin.H{"error": "Songs not found"})
	}
}

func GetGenres(c *gin.Context) {
	
}
