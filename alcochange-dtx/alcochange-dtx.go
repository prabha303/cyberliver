package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"time"

	"github.com/rs/cors"

	"ecargoware/alcochange-dtx/conf"
	"ecargoware/alcochange-dtx/dbcon"
	"ecargoware/alcochange-dtx/dbscripts"
	"ecargoware/alcochange-dtx/routes"
	"ecargoware/alcochange-dtx/sentryaccounts"
)

func main() {
	log.Println("-----------------------------------------------------------")
	log.Println(time.Now().UTC())

	var configFile = flag.String("conf", "", "configuration file(mandatory)")

	flag.Parse()
	if flag.NFlag() != 1 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Parsing configuration
	if err := conf.Parse(*configFile); err != nil {
		log.Fatalln("ERROR: ", err)
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Configuring Goroutine to use all the available CPUs
	// Note: GOMAXPROCS is manually set as application is deployed in Docker container with replication
	cpu, _ := strconv.Atoi(os.Getenv("GOMAXPROCS"))
	runtime.GOMAXPROCS(cpu)
	log.Println("INFO: Number of cpu configured - ", cpu)

	// Setting application mode
	setAppMode()

	// if conf.ServerFlag {
	// 	conf.InitateServerConfigurations()
	// }

	dbcon.Connect()
	defer dbcon.Close()

	dbscripts.InitDB()
	sentryaccounts.InitiateSentryLog()

	router := routes.RouterConfig()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders: []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
	})

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", conf.Cfg.PORT),
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

// Helper function for setting app mode
func setAppMode() {
	conf.Cfg.DEV_MODE = !conf.Cfg.PROD_MODE && !conf.Cfg.STAG_MODE
	if conf.Cfg.DEV_MODE {
		log.Println("Running as dev mode")
		return
	}

	if conf.Cfg.PROD_MODE {
		log.Println("Running as prod mode")
		return
	}

	log.Println("Running as staging mode")
}
