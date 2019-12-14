package infrastructure

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	router.GET("/users", func(c *gin.Context) {
		print("実行された")
	})

	Router = router
}
