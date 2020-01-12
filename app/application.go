package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// StartApplication initialization method
func StartApplication() {
	mapUrls()
	router.Run(":8082")
}
