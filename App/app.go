package App

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"task/Config"
	"task/Controllers/TaskController"
	"task/Database"
)

func appLoad() error {
	err := Config.Init()
	if err == nil {
		err = Database.Connect()
		if err == nil {
			err = Database.Migrate()
			if err != nil {
				return err
			} else {
				return nil
			}
		} else {
			return err
		}
	} else {
		return err
	}
}
func Start() {
	err := appLoad()
	if err != nil {
		fmt.Println(err)
		return
	}
	app := gin.Default()
	app.POST("/api/createTask", TaskController.CreateTask)
	app.POST("/api/getTasks", TaskController.GetTasks)
	app.PATCH("/api/setDone", TaskController.SetDone)
	app.DELETE("/api/delete/:id", TaskController.Delete)
	app.Run(Config.Config.Port)
}
