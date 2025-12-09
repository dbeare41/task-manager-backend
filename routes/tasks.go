package routes

import (
	"my-task-manager/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTask(context *gin.Context) {
	var task models.Task
	err := context.ShouldBindJSON(&task)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse task"})
		return
	}
	userid := context.GetInt64("userId")
	task.UserID = userid
	err = task.SaveTask()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Error saving task"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Task saved"})
}

func GetTasks(context *gin.Context) {
	tasks, err := models.GetAllTasks()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Error fetching tasks"})
		return
	}
	context.JSON(http.StatusOK, tasks)
}
