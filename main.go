package main

import (
	"testDataDeck/mappings"
	_ "testDataDeck/mappings"

	"github.com/gin-gonic/gin"

)

func main() {
	
	Router := gin.Default();
	
	mappings.UrlFunctions(Router)
	
	Router.Run()

}
