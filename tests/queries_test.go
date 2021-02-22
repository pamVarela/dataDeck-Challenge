package tests

import (
	"testDataDeck/routes"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var testRouter *gin.Engine = routes.UrlFunctions()

//unit testing API function in getGenre if it exists "Pop"
func TestQueries(t *testing.T) {
	
	gin.SetMode(gin.TestMode)

	t.Run("getGenre", func(t *testing.T) {
		response := httptest.NewRecorder()
		req, err := http.NewRequest("GET", "/getGenre/Pop", nil)
		if err != nil {
			t.Error(err)
		}
		testRouter.ServeHTTP(response, req)
		assert.Equal(t, response.Code, 200)
	})

}
