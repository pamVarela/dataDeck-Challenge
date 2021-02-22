package main

import (
	"testDataDeck/routes"
)

func main() {
	router := routes.UrlFunctions()
	router.Run()
}

