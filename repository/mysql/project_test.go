package mysql

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ryan0x44/todo-microservice/repository"
)

func TestProjectRepository_FindAll(t *testing.T) {
	if !*mysql {
		t.Skip()
	}
	_, err := projects.FindAll()
	if err != nil {
		t.Error(err)
	}
}

func TestProjectRepository_StoreAndDelete(t *testing.T) {
	if !*mysql {
		t.Skip()
	}
	var err error
	// Insert
	project := repository.Project{Name: "test a"}
	project, err = projects.Store(project)
	if err != nil {
		t.Error(err)
	}
	if project.ID <= 0 {
		t.Errorf("Invalid project ID after insert: %d", project.ID)
	}
	// Update
	project.Name = "test b"
	project, err = projects.Store(project)
	if err != nil {
		t.Error(err)
	}
	// Delete
	err = projects.Delete(project)
	if err != nil {
		t.Error(err)
	}
}
