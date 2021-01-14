package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rs/cors"

	"ecargoware/ec-user/conf"
	"ecargoware/ec-user/dbcon"
	"ecargoware/ec-user/dbscripts"
	"ecargoware/ec-user/routes"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flag.BoolVar(&conf.ServerFlag, "prod", false, "prod Flag for producton mode run")
	flag.Parse()

	if conf.ServerFlag {
		conf.InitateServerConfigurations()
	}

	dbcon.Connect()
	defer dbcon.Close()

	dbscripts.InitDB()

	router := routes.RouterConfig()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders: []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
	})

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", conf.Port),
		ReadTimeout:  90 * time.Second,
		WriteTimeout: 90 * time.Second,
		Handler:      c.Handler(router),
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	//Graceful shut down
	go func() {
		<-quit
		log.Println("Server is shutting down...")

		//Close resources before shut down
		dbcon.Close()

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		//Shutdown server
		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Unable to gracefully shutdown the server: %v\n", err)
		}

		//Close channels
		close(quit)
		close(done)
	}()

	log.Printf("Listening on: %d", conf.Port)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Error in listening server: %s", err.Error())
	}

	<-done
	log.Fatal("Server stopped")

}
