package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/NYTimes/gizmo/config/mysql"
	"github.com/go-kit/kit/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	"github.com/ryan0x44/todo-microservice/go-kit/internal"
)

type config struct {
	MySQL mysql.Config
}

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	)
	flag.Parse()
	// Create logger
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	}
	// Load config from env
	var c config
	err := envconfig.Process("", &c)
	if err != nil {
		logger.Log("Error parsing config", "err", err)
		return
	}
	// Connect to DB
	db, err := c.MySQL.DB()
	if err != nil {
		logger.Log("Error load DB driver", "err", err)
	}
	err = db.Ping()
	if err != nil {
		logger.Log("Error connecting to DB: ", "err", err)
	}
	defer db.Close()
	// Create service
	s := internal.NewService()
	// Create endpoints
	e := internal.MakeServerEndpoints(s)
	// Create http server
	handler := internal.MakeHTTPHandler(e)
	// Handle signals
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()
	// Run web server
	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()
	// Log web server errors:
	logger.Log("exit", <-errs)
}
