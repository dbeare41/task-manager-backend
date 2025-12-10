package models

import (
	"my-task-manager/db"
)

type Task struct {
	Id          int64
	Title       string `binding:"required"`
	Description string `binding:"required"`
	Status      string `binding:"required"`
	UserID      int64
}

func (t *Task) SaveTask() error {
	query := "INSERT INTO tasks(title,description,status,user_id) VALUES (?,?,?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	stmt.Exec(t.Title, t.Description, t.Status, t.UserID)
	return err
}

func GetAllTasks() ([]Task, error) {
	query := "SELECT * FROM tasks"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []Task
	for rows.Next() {
		var t Task
		err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.Status, &t.UserID)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)

	}
	return tasks, nil
}

func GetTaskById(id int64) (*Task, error) {
	query := "SELECT * FROM tasks WHERE id =?"
	var t Task
	row := db.DB.QueryRow(query, id)
	err := row.Scan(&t.Id, &t.Title, &t.Description, &t.Status, &t.UserID)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (t Task) UpdateTaskInfo() error {
	query := `
	UPDATE tasks SET title = ?, description = ?, status =? WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(t.Title, t.Description, t.Status, t.Id)
	if err != nil {
		return err
	}
	return nil
}
