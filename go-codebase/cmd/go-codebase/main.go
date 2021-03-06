package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime/debug"
	"strings"
	"syscall"
	"time"

	"gitlab.com/Wuriyanto/go-codebase/config"
	"gitlab.com/Wuriyanto/go-codebase/internal/app"
)

func main() {
	rootApp, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	rootApp = strings.TrimSuffix(rootApp, "/cmd/go-codebase") // trim this path location, for cleaning root path

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute) // timeout to shutdown server and init configuration
	defer func() {
		cancel()
		if r := recover(); r != nil {
			// notify if this service failed to start (to logstash or slack)
			fmt.Println("Failed to start service:", r)
			fmt.Printf("Stack trace: \n%s\n", debug.Stack())
		}
	}()

	cfg := config.Init(ctx, rootApp)
	defer cfg.Exit(ctx)

	myService := app.New(cfg)

	// serve http server
	go myService.ServeHTTP()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	myService.Shutdown(ctx)
}
