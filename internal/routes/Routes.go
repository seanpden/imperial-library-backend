package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/seanpden/imperial-library-backend/internal/controllers"
)

// PUBLIC FUNC BECAUSE OF UPPERCASE LETTER
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", controllers.HomeReponse)

	return r
}
