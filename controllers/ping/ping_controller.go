package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping controller handler
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
