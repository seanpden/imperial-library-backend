package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seanpden/imperial-library-backend/internal/controllers"
)

// PUBLIC FUNC BECAUSE OF UPPERCASE LETTER
func SetupRouter() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", controllers.HomeReponse)

	r.GET("/api/books/all", controllers.GetAllBooks)
	r.GET("/api/books/find", controllers.GetBooksFuzzy)

	r.GET("/debug", func(c *gin.Context) {
		q := c.Query("foo")
		c.IndentedJSON(http.StatusOK, q)
	})

	return r
}
