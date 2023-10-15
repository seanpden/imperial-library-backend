package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seanpden/imperial-library-backend/internal/models"
)

func HomeReponse(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.ReadBookContent())
}

func GetBookContent(c *gin.Context) {
	c.JSON(http.StatusOK, "book content")
}

func GetBookInfo(c *gin.Context) {
	c.JSON(http.StatusOK, "book info")
}
