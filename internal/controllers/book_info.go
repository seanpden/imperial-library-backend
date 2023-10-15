package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/seanpden/imperial-library-backend/models"
)

func getBookInfo(c *gin.Context) {
	c.JSON(http.StatusOK, "test")
}
