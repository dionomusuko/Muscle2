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
	c.JSON(http.StatusOK, users)
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

func CreateUser(c *gin.Context) {
	var user model.User
	db := db.NewDB()

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.Create(&user)
	c.JSON(http.StatusOK, &user)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	db := db.NewDB()
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.BindJSON(&user)
	db.Save(&user)
	c.JSON(http.StatusOK, &user)
}
