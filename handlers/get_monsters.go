package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMonster() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"key": "values",
		})
	}
}
