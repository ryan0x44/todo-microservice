package mysql

import (
	"database/sql"
	"fmt"

	"github.com/ryan0x44/todo-microservice/repository"
)

type TaskRepository struct {
	DB *sql.DB
}

const findAllTasksSQL = "SELECT id,project_id,description FROM tasks"
const findTasksByProjectSQL = "SELECT id,project_id,description FROM tasks WHERE project_id = ?"
const insertTaskSQL = "INSERT INTO tasks (project_id,description) VALUES (?,?)"
const updateTaskSQL = "UPDATE tasks SET project_id = ?, description = ? WHERE id = ?"
const deleteTaskSQL = "DELETE FROM tasks WHERE id = ?"

func (r *TaskRepository) FindAll() (results []repository.Task, err error) {
	var rows *sql.Rows
	rows, err = r.DB.Query(findAllTasksSQL)
	if err != nil {
		return []repository.Task{}, err
	}
	defer rows.Close()
	for rows.Next() {
		result := repository.Task{Project: &repository.Project{}}
		err = rows.Scan(&result.ID, &result.Project.ID, &result.Description)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return
}

func (r *TaskRepository) FindByProject(p repository.Project) (results []repository.Task, err error) {
	if p.ID <= 0 {
		return nil, fmt.Errorf("Project ID invalid")
	}
	var rows *sql.Rows
	rows, err = r.DB.Query(findTasksByProjectSQL, p.ID)
	if err != nil {
		return []repository.Task{}, err
	}
	defer rows.Close()
	for rows.Next() {
		result := repository.Task{Project: &repository.Project{}}
		err = rows.Scan(&result.ID, &result.Project.ID, &result.Description)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return
}

func (r *TaskRepository) Store(t repository.Task) (repository.Task, error) {
	var result sql.Result
	err := Transact(r.DB, func(tx *sql.Tx) (err error) {
		if t.ID > 0 {
			result, err = tx.Exec(updateTaskSQL, t.ID, t.Project.ID, t.Description)
		} else {
			result, err = tx.Exec(insertTaskSQL, t.Project.ID, t.Description)
		}
		return
	})
	if t.ID <= 0 && err == nil {
		t.ID, err = result.LastInsertId()
	}
	return t, err
}

func (r *TaskRepository) Delete(p repository.Task) (err error) {
	return Transact(r.DB, func(tx *sql.Tx) (err error) {
		if p.ID > 0 {
			_, err = tx.Exec(deleteTaskSQL, p.ID)
		} else {
			return fmt.Errorf("Task ID invalid")
		}
		return
	})
}
