package todo

import (
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ToggleDone(dbConn *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody struct {
			Id uint `json:"id" url:"required"`
		}

		if err := c.ShouldBindUri(&requestBody); err != nil {
			log.Error(err)
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, string(requestBody.Id) + " is not valid")
			return
		}

		var todo Todo
		dbConn.First(&todo, c.Param("id"))

		if result := dbConn.First(&todo, c.Param("id")); result.Error != nil {
			log.Error(result.Error)
			c.AbortWithStatusJSON(http.StatusNotFound, "Item not found")
			return
		}

		done := true

		if todo.Done == true {
			done = false
		}

		dbConn.Model(&todo).Update("Done", done)

		c.JSON(http.StatusOK, "Done is toggled")
	}
}