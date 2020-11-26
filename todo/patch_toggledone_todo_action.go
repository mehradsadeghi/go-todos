package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ToggleDone(dbConn *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody struct {
			ID string `json:"Id" url:"required"`
		}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, requestBody.ID+" is not valid")
			return
		}

		// it can get better
		id := c.Params.ByName("Id")

		var todo Todo

		dbConn.First(&todo, id)

		if result := dbConn.First(&todo, id); result.Error != nil {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "Item not found")
			return
		}

		done := true

		if(todo.Done == true) {
			done = true
		}

		dbConn.Model(&todo).Update("Done", done)

		c.JSON(http.StatusOK, "OK")
	}
}