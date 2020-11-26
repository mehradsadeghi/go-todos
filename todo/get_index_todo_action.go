package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(dbConn *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var todos []Todo
		dbConn.Find(&todos)
		c.JSON(http.StatusOK, todos)
	}
}
