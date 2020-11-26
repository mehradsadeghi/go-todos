package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Show(dbConn *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody struct {
			Id uint `json:"id" url:"required"`
		}

		if err := c.ShouldBindUri(&requestBody); err != nil {
			// todo logging logrus
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, string(requestBody.Id) + " is not valid")
			return
		}

		id := c.Params.ByName("id")

		var todo Todo

		if result := dbConn.First(&todo, id); result.Error != nil {
			// todo logging logrus
			c.AbortWithStatusJSON(http.StatusNotFound, "Item not found")
			return
		}

		c.JSON(http.StatusOK, todo)
	}
}
