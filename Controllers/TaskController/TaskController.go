package TaskController

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"task/Database/Models/Task"
	"task/Services/TaskService"
	"task/Utils"
)

type Message struct {
	Message string `json:"message"`
}

func CreateTask(ctx *gin.Context) {
	var CreateData Task.CreateTask
	err := ctx.BindJSON(&CreateData)
	if err == nil {
		result, err := TaskService.CreateTask(&CreateData)
		if err == nil {
			ctx.JSON(200, *result)
		} else {
			status, message := Utils.GetMessageAndStatus(err)
			ctx.JSON(status, Message{
				Message: message,
			})
		}
	} else {
		ctx.JSON(500, Message{
			Message: "error on json parse",
		})
	}
}

func GetTasks(ctx *gin.Context) {
	var GetData Task.GetTasks
	err := ctx.BindJSON(&GetData)
	if err == nil {
		result, err := TaskService.GetTasks(&GetData)
		if err == nil {
			ctx.JSON(200, *result)
		} else {
			status, message := Utils.GetMessageAndStatus(err)
			ctx.JSON(status, Message{
				Message: message,
			})
		}
	} else {
		ctx.JSON(500, Message{
			Message: "error on json parse",
		})
	}
}

func SetDone(ctx *gin.Context) {
	var SetDone Task.SetDone
	err := ctx.BindJSON(&SetDone)
	if err == nil {
		err := TaskService.SetDone(&SetDone)
		if err == nil {
			ctx.JSON(200, Message{
				Message: "ok",
			})
		} else {
			status, message := Utils.GetMessageAndStatus(err)
			ctx.JSON(status, Message{
				Message: message,
			})
		}
	} else {
		ctx.JSON(500, Message{
			Message: "error on json parse",
		})
	}
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := strconv.Atoi(id)
	if err == nil {
		err := TaskService.Delete(result)
		if err == nil {
			ctx.JSON(200, Message{
				Message: "ok",
			})
		} else {
			status, message := Utils.GetMessageAndStatus(err)
			ctx.JSON(status, Message{
				Message: message,
			})
		}
	} else {
		ctx.JSON(500, Message{
			Message: "error on params parse",
		})
	}
}
