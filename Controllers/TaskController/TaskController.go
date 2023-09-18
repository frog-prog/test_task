package TaskController

import (
	"errors"
	"task/Database/Models/Task"
	"unicode/utf8"
)

func CreateTask(task *Task.CreateTask) (*Task.Task, error) {
	if utf8.RuneCountInString((*task).Text) > 1000 {
		return nil, errors.New("500:text too long")
	}
	createdTask, err := Task.Create(&task.Text, &task.Date)
	if err == nil {
		return createdTask, err
	}
	return nil, err
}

func GetTasks(getTasks *Task.GetTasks) (*[]Task.Task, error) {
	tasks, err := Task.GetMany((*getTasks).Page, (*getTasks).Count, (*getTasks).Done)
	if err == nil {
		return tasks, err
	}
	return nil, err
}

func SetDone(task *Task.SetDone) error {
	return Task.Update((*task).Id, (*task).Done)
}

func Delete(id int) error {
	return Task.Delete(id)
}
