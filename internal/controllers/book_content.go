package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seanpden/imperial-library-backend/models"
)

func getBookContent(c *gin.Context) {
	c.JSON(http.StatusOK, "test")
}
