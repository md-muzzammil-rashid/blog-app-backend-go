package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/md-muzzammil-rashid/blog-app-backend-go/cmd/api"
	"github.com/md-muzzammil-rashid/blog-app-backend-go/config"
	"github.com/md-muzzammil-rashid/blog-app-backend-go/database"
)


func main() {
	// load configuration
	
	cfg, err := config.InitConfig()
	
	if err != nil {
		log.Fatalf("Error in initialization of configuration  \n %v", err)
	}
	// database configuration

	db, err := database.NewRepository(cfg); if err != nil {
		log.Fatal(err.Error())
	} 
	defer db.Close()

	slog.Info("Database Connection established!")

	//setup router

	router := mux.NewRouter()

	//setup server
	server := http.Server{
		Handler: router,
		Addr: cfg.Server.Address+":"+cfg.Server.Port,
	}
	
	api.ApiHandler(router, db)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func () {
		slog.Info("Server is running on port ", cfg.Server.Port)
	
		err = server.ListenAndServe(); if err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	<- done

	// gracefully shutting down
	slog.Info("Shutting down server ...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)

	defer cancel()

	if err := server.Shutdown(ctx); err!= nil {
        slog.Error("Error gracefully shutting down the server: %v", err)
    }

	slog.Info("Server shutdown completed successfully")

}