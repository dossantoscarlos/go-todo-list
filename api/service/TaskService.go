package service

import (
	"database/sql"

	"github.com/dossantoscarlos/go-todo-list/api/db"
	"github.com/dossantoscarlos/go-todo-list/api/models"
)

func AllTask() ([]models.Task, error) {

	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Name, &task.Status)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func SaveTask(task models.Task) (int64, error) {

	conn, err := db.Connect()
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	result, err := conn.Exec("INSERT INTO tasks (name, status) VALUES (?, ?)", task.Name, task.Status)
	if err != nil {
		return 0, err
	}

	taskID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int64(taskID), nil
}
func FindOneTask(ID int64) (*models.Task, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM tasks where id = ?", ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var task models.Task
	for rows.Next() {
		err := rows.Scan(&task.ID, &task.Name, &task.Status)
		if err != nil {
			return nil, err
		}
	}

	if task.ID == 0 {
		return nil, sql.ErrNoRows
	}

	return &task, nil
}
