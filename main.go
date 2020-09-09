package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leandrofreires/crm/database"
	"github.com/leandrofreires/crm/router"
)

func main() {
	client := database.Connect()
	database.Db = client.Database("crm")
	handler := gin.Default()
	//load midlewares
	// r.Use(cors.Default())
	// r.Static("/images", "./images")
	// r.StaticFile("favicon.ico", "./logo.png")
	router.APIRouter(handler)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal("Server forced to disconected from database:", err)
			return
		}
		fmt.Print("Sucess disconnected")
	}()
}
