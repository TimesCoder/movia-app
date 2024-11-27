package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/TimesCoder/movie-app/config"
	"github.com/TimesCoder/movie-app/internal/builder"
	"github.com/TimesCoder/movie-app/pkg/database"
	"github.com/TimesCoder/movie-app/pkg/server"
)

func main() {
	// load configuration via env
	cfg, err := config.NewConfig(".env")
	checkError(err)

	// Init & start database
	db, err := database.InitDatabase(cfg.MySQLConfig)
	checkError(err)

	// Build routes
	publicRoutes := builder.BuildPublicRoutes(db)
	privateRoutes := builder.BuildPrivateRoutes(db)

	// Init sever
	srv := server.NewServer(publicRoutes, privateRoutes)
	runServer(srv, cfg.PORT)
	waitForShutdown(srv)

	// start server

}

func waitForShutdown(srv *server.Server) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			srv.Logger.Fatal(err)
		}
	}()
}

func runServer(srv *server.Server, port string) {
	go func() {
		err := srv.Start(fmt.Sprintf(":%s", port))
		log.Fatal(err)
	}()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
