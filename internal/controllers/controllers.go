package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seanpden/imperial-library-backend/internal/models"
)

func HomeReponse(c *gin.Context) {
	msg := map[string]string{"message": "Home!"}
	c.IndentedJSON(http.StatusOK, msg)
}

func GetAllBooks(c *gin.Context) {
	value, err := models.GetAllBooks()
	if err != nil {
		log.Panic(err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, value)
}

func GetBooksFuzzy(c *gin.Context) {
	titleq := c.Query("title")
	authorq := c.Query("author")
	textq := c.Query("text")

	value, err := models.GetBooksFuzzy(titleq, authorq, textq)
	if err != nil {
		log.Panic(err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, value)
}
