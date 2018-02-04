package mysql

import (
	"database/sql"
	"fmt"

	"github.com/ryan0x44/todo-microservice/repository"
)

type ProjectRepository struct {
	DB *sql.DB
}

const findAllProjectsSQL = "SELECT id,name FROM projects"
const insertProjectSQL = "INSERT INTO projects (name) VALUES (?)"
const updateProjectSQL = "UPDATE projects SET name = ? WHERE id = ?"
const deleteProjectSQL = "DELETE FROM projects WHERE id = ?"

func (r *ProjectRepository) FindAll() (results []repository.Project, err error) {
	var rows *sql.Rows
	rows, err = r.DB.Query(findAllProjectsSQL)
	if err != nil {
		return []repository.Project{}, err
	}
	defer rows.Close()
	for rows.Next() {
		result := repository.Project{}
		err = rows.Scan(&result.ID, &result.Name)
		if err != nil {
			return results, err
		}
		results = append(results, result)
	}
	return
}

func (r *ProjectRepository) Store(p repository.Project) (repository.Project, error) {
	var result sql.Result
	err := Transact(r.DB, func(tx *sql.Tx) (err error) {
		if p.ID > 0 {
			result, err = tx.Exec(updateProjectSQL, p.ID, p.Name)
		} else {
			result, err = tx.Exec(insertProjectSQL, p.Name)
		}
		return
	})
	if p.ID <= 0 && err == nil {
		p.ID, err = result.LastInsertId()
	}
	return p, err
}

func (r *ProjectRepository) Delete(p repository.Project) (err error) {
	return Transact(r.DB, func(tx *sql.Tx) (err error) {
		if p.ID > 0 {
			_, err = tx.Exec(deleteProjectSQL, p.ID)
		} else {
			return fmt.Errorf("Project ID invalid")
		}
		return
	})
}
