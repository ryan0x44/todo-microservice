package mysql

import (
	"database/sql"
	"flag"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ryan0x44/todo-microservice/repository"

	mysqlconfig "github.com/NYTimes/gizmo/config/mysql"
	"github.com/kelseyhightower/envconfig"
)

var mysql = flag.Bool("mysql", false, "run mysql integration tests")
var DB *sql.DB
var projects repository.ProjectRepository
var tasks repository.TaskRepository

func TestMain(m *testing.M) {
	flag.Parse()
	if *mysql {
		var err error
		// Parse env config
		var c mysqlconfig.Config
		err = envconfig.Process("", &c)
		if err != nil {
			log.Fatalf("Error parsing MySQL config", err)
		}
		// Connect to DB
		DB, err = c.DB()
		if err != nil {
			log.Fatalf("Error load DB driver", err)
		}
		err = DB.Ping()
		if err != nil {
			log.Fatalf("Error connecting to DB: ", err)
		}
		// Create repositories
		projects = &ProjectRepository{DB: DB}
		tasks = &TaskRepository{DB: DB}
	}
	// Run tests
	result := m.Run()
	if *mysql {
		// Disconnect from DB
		DB.Close()
	}
	os.Exit(result)
}
