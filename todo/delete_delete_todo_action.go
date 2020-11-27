package todo

import (
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Delete(dbConn *gorm.DB) gin.HandlerFunc {

	return func (c *gin.Context) {
		var requestBody struct {
			Id string `json:"Id" url:"required"`
		}

		if err := c.ShouldBindUri(&requestBody); err != nil {
			log.Error(err)
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, requestBody.Id + " is not valid")
			return
		}

		id := c.Params.ByName("id")

		var todo Todo

		if result := dbConn.Delete(&todo, id); result.Error != nil {
			log.Error(result.Error)
			c.AbortWithStatusJSON(http.StatusNotFound, "Item not found")
			return
		}

		c.JSON(http.StatusOK, "todo is deleted")
	}
}
