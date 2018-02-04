package mysql

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ryan0x44/todo-microservice/repository"
)

func TestTaskRepository_FindAll(t *testing.T) {
	if !*mysql {
		t.Skip()
	}
	_, err := tasks.FindAll()
	if err != nil {
		t.Error(err)
	}
}

func TestTaskRepository_StoreAndDelete(t *testing.T) {
	if !*mysql {
		t.Skip()
	}
	var err error
	// Insert project
	project := repository.Project{Name: "project with task"}
	project, err = projects.Store(project)
	if err != nil {
		t.Error(err)
	}
	if project.ID <= 0 {
		t.Errorf("Invalid project ID after insert: %d", project.ID)
	}
	// Insert task
	task := repository.Task{
		Description: "task a",
		Project:     &project,
	}
	task, err = tasks.Store(task)
	if err != nil {
		t.Error(err)
	}
	if task.ID <= 0 {
		t.Errorf("Invalid task ID after insert: %d", task.ID)
	}
	// Get task by project
	tasksByProject, err := tasks.FindByProject(project)
	if err != nil {
		t.Error(err)
	}
	if len(tasksByProject) != 1 {
		t.Errorf("Expected to find exactly 1 task for project", err)
	}
	if tasksByProject[0].ID != task.ID {
		t.Errorf(
			"Task ID mismatch. Expected: %d, Got: %d",
			tasksByProject[0].ID,
			task.ID,
		)
	}
	// Update task
	task.Description = "task b"
	task, err = tasks.Store(task)
	if err != nil {
		t.Error(err)
	}
	// Delete task
	err = tasks.Delete(task)
	if err != nil {
		t.Error(err)
	}
	// Delete project
	err = projects.Delete(project)
	if err != nil {
		t.Error(err)
	}
}
