package controller

import (
	"github.com/dionomusuko/muscle2/db"
	"github.com/dionomusuko/muscle2/model"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []model.User
	db := db.GetDB()
	db.Find(&users)
	c.JSON(200, users)
}
