package app

import (
	"github.com/jceatwell/bookstore_users-api/controllers/ping"
	"github.com/jceatwell/bookstore_users-api/controllers/users"
)

// mapUrls maps routes to Controllers
func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
