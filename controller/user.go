package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leandrofreires/crm/model"
	"go.mongodb.org/mongo-driver/mongo"
)

//CreateUser is a Handler for router
func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := user.Save(); err != nil {
		merr := err.(mongo.WriteException)
		c.JSON(http.StatusBadRequest, gin.H{"error": merr.WriteErrors[0].Message})
		return
	}
	c.JSON(http.StatusCreated, user)
}
func GetUsers(c *gin.Context) {
	var user model.User
	users, err := user.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, users)
}
