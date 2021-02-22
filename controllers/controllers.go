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

// Function API GET that returns a list of songs searchable by artist, song or name
// Parameters: parameter
func GetSong(c *gin.Context) {
	
	initDB := GetConnectionDB()

	parameter := c.Params.ByName("params")

	var listSongs []models.Song

	query := `SELECT s.song, s.artist, g.name, s.length FROM songs AS s JOIN genres AS g ON g.id = s.genre WHERE `

	rowsArtists, _ := initDB.Query(query + `s.artist = ?`, parameter)
	rowsSongs, _  := initDB.Query(query + `s.song = ?`, parameter)
	rowsGenres, _ := initDB.Query(query + `g.name = ?`, parameter)

	if rowsArtists != nil{
		for rowsArtists.Next() {
			var res models.Song
			rowsArtists.Scan(&res.Song, &res.Artist, &res.Genre, &res.Length)
			listSongs = append(listSongs, res)
		}
	}

	if rowsSongs != nil{
		for rowsSongs.Next() {
			var res models.Song
			rowsSongs.Scan(&res.Song, &res.Artist, &res.Genre, &res.Length)
			listSongs = append(listSongs, res)
		}
	}

	if rowsGenres != nil{
		for rowsGenres.Next() {
			var res models.Song
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


func isEmpty(object interface{}) bool {
	return (object != nil)
    
}

// Function API GET that returns list of genres
// Parameters: Name
func GetGenre(c *gin.Context) {

	initDB := GetConnectionDB()
	
	name := c.Params.ByName("name")

	var res models.Genre

	row := initDB.QueryRow("SELECT ID, name FROM genres WHERE name=?", name)
		
	row.Scan(&res.Id, &res.Name)
	if(isEmpty(res)){
		c.JSON(200, res)
	}else{
		c.JSON(404, gin.H{"error": "Genre not found"})
	}		
}
// Function API GET that returns a list of songs 
// Parameters: Min, Max
func GetSongsByLength(c *gin.Context) {

	initDB := GetConnectionDB()
	
	min := c.Params.ByName("min")
	max := c.Params.ByName("max")

	var listSongs []models.Song
	query := `SELECT s.song, s.artist, g.name, s.length FROM songs AS s JOIN genres AS g ON g.id = s.genre WHERE s.length BETWEEN ? AND ?`
	rows, _ := initDB.Query(query, min, max)
	
	if rows != nil{
		for rows.Next() {
			var res models.Song
			rows.Scan(&res.Song, &res.Artist, &res.Genre, &res.Length)
			listSongs = append(listSongs, res)
		}
	}

	defer rows.Close()

	if listSongs != nil {
		c.JSON(200, listSongs)
	} else {
		c.JSON(404, gin.H{"error": "Songs not found"})
	}
}

// Function API GET  that returns the list of genres with the total songs and total length by genre
func GetGenres(c *gin.Context){

	initDB := GetConnectionDB()

	type genres struct  {
		Name string `json:"genre"`
		TotalSongs int `json:"totalSongs"`
		TotalLength int `json:"totalLength"`
	}

	var listGenres []genres
	query := `SELECT g.name, SUM(s.length) AS totalLength, COUNT(*) AS totalSongs FROM genres AS g JOIN songs AS s ON s.genre = g.id GROUP BY g.name`
	rows, _ := initDB.Query(query)
	
	if rows != nil{
		for rows.Next() {
			var res genres
			rows.Scan(&res.Name, &res.TotalLength, &res.TotalSongs)
			listGenres = append(listGenres, res)
		}
	}

	defer rows.Close()

	if listGenres != nil {
		c.JSON(200, listGenres)
	} else {
		c.JSON(404, gin.H{"error": "Genres not found"})
	}
	
}