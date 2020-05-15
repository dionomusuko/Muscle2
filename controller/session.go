package controller

import (
	"github.com/dionomusuko/muscle2/db"
	"github.com/dionomusuko/muscle2/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignIn(c *gin.Context) {
	var user model.User
	db := db.NewDB()
	//formからuserを取得
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	name := db.Find(&user.Name)
	email := db.Find(&user.Email)
	//nameまたはemailが該当しているか確認
	if name != nil || email != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	c.JSON(http.StatusOK, name)
}
