package controller

import (
	"github.com/dionomusuko/muscle2/db"
	"github.com/dionomusuko/muscle2/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	var users []model.User
	db := db.NewDB()
	db.Find(&users)
	c.JSON(200, users)
}

func ShowUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	db := db.NewDB()
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, &user)
}
