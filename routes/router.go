package routes

import (
	"ginblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	router := gin.Default()

	router.Group("api/v1")
	{
		router.GET("api/v1/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}

	router.Run(utils.HttpPort)
}
