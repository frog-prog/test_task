package Task

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"task/Database"
	"time"
)

type Task struct {
	Id   int       `json:"id" db:"id"`
	Text string    `json:"text" db:"text"`
	Done bool      `json:"done" db:"done"`
	Date time.Time `json:"date" db:"date"`
}
type CreateTask struct {
	Text string `json:"text"`
	Date string `json:"date"`
}
type GetTasks struct {
	Page  int  `json:"page"`
	Count int  `json:"count"`
	Done  bool `json:"done"`
}
type SetDone struct {
	Id   int  `json:"id"`
	Done bool `json:"done"`
}

func Create(text *string, date *string) (*Task, error) {
	var newTask Task
	ctx := context.Background()
	query := `INSERT INTO tasks (text, date) VALUES ($1, $2) returning id,text,done,date;`
	err := Database.Connection.QueryRow(ctx, query, *text, *date).Scan(
		&newTask.Id,
		&newTask.Text,
		&newTask.Done,
		&newTask.Date)
	if err == nil {
		fmt.Println(newTask.Id, newTask.Text, newTask.Done, newTask.Date)
		return &newTask, nil
	}
	return nil, errors.New("500:database error")
}

func GetMany(page int, count int, done bool) (*[]Task, error) {
	offset := (page - 1) * count
	ctx := context.Background()
	query := `SELECT * FROM tasks WHERE done=$1 OFFSET $2 LIMIT $3`
	rows, err := Database.Connection.Query(ctx, query, done, offset, count)
	if err == nil {
		taskList, err := pgx.CollectRows[Task](rows, pgx.RowToStructByName[Task])
		if len(taskList) == 0 {
			return nil, errors.New("404:page not found")
		}
		if err != nil {
			return nil, errors.New("500:parse error")
		}
		return &taskList, nil
	}
	return nil, errors.New("500:database error")
}

func Update(id int, done bool) error {
	ctx := context.Background()
	query := `UPDATE tasks SET done=$1 WHERE id=$2 returning id, text, done, date`
	result, err := Database.Connection.Exec(ctx, query, done, id)
	fmt.Println(result)
	if err == nil {
		if result.RowsAffected() == 0 {
			return errors.New("404:record not found")
		}
		return nil
	}
	return errors.New("500:database error")
}

func Delete(id int) error {
	ctx := context.Background()
	query := `DELETE FROM tasks WHERE id=$1`
	result, err := Database.Connection.Exec(ctx, query, id)
	if err == nil {
		if result.RowsAffected() == 0 {
			return errors.New("404:record not found")
		}
		return nil
	}
	return errors.New("500:database error")
}
