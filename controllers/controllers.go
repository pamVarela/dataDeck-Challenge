package controllers

import (
	"log"
	"testDataDeck/models"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// Function enables Header for make requests
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

// Bool Function returns if is empty the object Genre
func isEmptyGenre(object models.Genre) bool {
	if(object.Id != 0 && object.Name != ""){
		return true
	}
	return false
}

// Function API GET that returns a list of songs searchable by song
// Parameters: song
func GetSongsBySong(c *gin.Context) {
	
	parameter := c.Params.ByName("song")

	query := `SELECT s.song, s.artist, g.name, s.length FROM songs AS s JOIN genres AS g ON g.id = s.genre WHERE LOWER(s.song) = LOWER(?) `

	listSongs, err := getSongsByParameter(query,parameter)

	if listSongs != nil && err == nil {
		c.JSON(200, listSongs)
	} else {
		c.JSON(404, gin.H{"error": "Songs not found"})
	}
}

// Function API GET that returns a list of songs searchable by artist
// Parameters: artist
func GetSongsByArtist(c *gin.Context) {

	parameter := c.Params.ByName("artist")

	query := `SELECT s.song, s.artist, g.name, s.length FROM songs AS s JOIN genres AS g ON g.id = s.genre WHERE LOWER(s.artist) = LOWER(?)`

	listSongs, err := getSongsByParameter(query, parameter)

	if listSongs != nil && err == nil {
		c.JSON(200, listSongs)
	} else {
		c.JSON(404, gin.H{"error": "Songs not found"})
	}
}

// Function API GET that returns a list of songs searchable by genre
// Parameters: genre
func GetSongsByGenre(c *gin.Context) {

	parameter := c.Params.ByName("genre")

	query := `SELECT s.song, s.artist, g.name, s.length FROM songs AS s JOIN genres AS g ON g.id = s.genre WHERE LOWER(g.name) = LOWER(?)`

	listSongs, err := getSongsByParameter(query, parameter)

	if listSongs != nil && err == nil {
		c.JSON(200, listSongs)
	} else {
		c.JSON(404, gin.H{"error": "Songs not found"})
	}
}

// Function is generic for the queries on the table 'Songs' that returns a list and error
// Parameters: query, parameter
func getSongsByParameter(query string, parameter string) (list []models.Song, err error){

	initDB := GetConnectionDB()

	var  listSongs []models.Song

	rows, err := initDB.Query(query, parameter)
	
	if rows != nil{
		for rows.Next() {
			var res models.Song
			rows.Scan(&res.Song, &res.Artist, &res.Genre, &res.Length)
			listSongs = append(listSongs, res)
		}
	}

	defer rows.Close()

	return listSongs, err
}

// Function API GET that returns list of genres
// Parameters: Name
func GetGenre(c *gin.Context) {

	initDB := GetConnectionDB()
	
	name := c.Params.ByName("name")

	var res models.Genre

	row := initDB.QueryRow("SELECT ID, name FROM genres WHERE LOWER(name)=LOWER(?)", name)	
	row.Scan(&res.Id, &res.Name)

	log.Println(isEmptyGenre(res))

	if(isEmptyGenre(res)){
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
	rows, err := initDB.Query(query, min, max)
	
	if rows != nil{
		for rows.Next() {
			var res models.Song
			rows.Scan(&res.Song, &res.Artist, &res.Genre, &res.Length)
			listSongs = append(listSongs, res)
		}
	}

	defer rows.Close()

	if listSongs != nil && err == nil {
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
	rows, err := initDB.Query(query)
	
	if rows != nil{
		for rows.Next() {
			var res genres
			rows.Scan(&res.Name, &res.TotalLength, &res.TotalSongs)
			listGenres = append(listGenres, res)
		}
	}

	defer rows.Close()

	if listGenres != nil && err == nil{
		c.JSON(200, listGenres)
	} else {
		c.JSON(404, gin.H{"error": "Genres not found"})
	}
	
}