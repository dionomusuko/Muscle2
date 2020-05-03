package controller

import (
	"github.com/dionomusuko/muscle2/db"
	"github.com/dionomusuko/muscle2/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//ユーザーIDをとって、表示する仕組みに変える
func GetTasks(c *gin.Context) {
	var tasks []model.Task
	db := db.NewDB()
	db.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

func ShowTask(c *gin.Context) {
	id := c.Param("id")
	var task model.Task
	db := db.NewDB()
	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task model.Task
	db := db.NewDB()
	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

}

func CreateTask() {
	//User_idを取る方法を決めてから
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task model.Task
	db := db.NewDB()
	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
}
