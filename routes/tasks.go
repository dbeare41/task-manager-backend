package routes

import (
	"my-task-manager/models"
	"net/http"
	"strconv"

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

func UpdateTask(context *gin.Context) {
	var task *models.Task
	taskId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse request"})
		return
	}

	task, err = models.GetTaskById(taskId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get task"})
		return
	}
	if task.UserID != context.GetInt64("userId") {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorised"})
		return
	}
	err = context.ShouldBindJSON(&task)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse request"})
		return
	}
	err = task.UpdateTaskInfo()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not update task"})
		return
	}
	context.JSON(http.StatusAccepted, gin.H{"Message": "Task updated"})
}

func GetTasks(context *gin.Context) {
	tasks, err := models.GetAllTasks()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Error fetching tasks"})
		return
	}
	context.JSON(http.StatusOK, tasks)
}
