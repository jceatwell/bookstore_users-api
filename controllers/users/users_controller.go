// Package users provides a domain model for working with users
package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jceatwell/bookstore_users-api/domain/users"
	"github.com/jceatwell/bookstore_users-api/services"
)

//CreateUser creates new user
func CreateUser(c *gin.Context) {
	var user users.User
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO: Handle Error
		return
	}

	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())
		//TODO: Handle json error
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: handle user creation error
		return
	}

	fmt.Println(err)
	fmt.Println(string(bytes))

	c.String(http.StatusNotImplemented, "TODO: Implement!")
}

//GetUser retrieve user details
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "TODO: Implement!")
}
