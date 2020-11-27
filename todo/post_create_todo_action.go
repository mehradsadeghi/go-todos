package todo

import (
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(dbConn *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody struct {
			Title string `json:"title" binding:"required"`
			Body  string `json:"body" binding:"required"`
		}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			log.Error(err)
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "Provided data is not valid")
			return
		}

		dbConn.Create(New(requestBody.Title, requestBody.Body))

		c.JSON(http.StatusCreated, "todo is created")
	}
}