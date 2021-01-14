package dbcon

import (
	"fmt"
	"log"
	"os"

	"ecargoware/alcochange-dtx/conf"

	"github.com/go-pg/pg"
)

var db *pg.DB

//Connect database
func Connect() {
	dbCon := pg.Connect(&pg.Options{
		Addr:     conf.DatabaseAddr,
		User:     conf.DatabaseUsername,
		Password: conf.DatabasePassword,
		Database: conf.DatabaseName,
	})

	db = dbCon
	log.Printf("Connected successfully")

	_, err := db.Exec("SELECT 1")
	if err != nil {
		fmt.Println("PostgreSQL is down")
		log.Fatalf("Unable to connect Postgres Database: %v\n", err)
		os.Exit(1)
	}

	db.AddQueryHook(dbLogger{})
	// isProfile := true
	// if isProfile {
	// 	db.OnQueryProcessed(func(event *pg.QueryEvent) {
	// 		query, err := event.FormattedQuery()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		log.Printf("\033[35m%s %s\033[39m\n\n", time.Since(event.StartTime), query)
	// 	})
	// }

	// if isProfile {
	// 	db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
	// 		query, err := event.FormattedQuery()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		log.Printf("\033[35m%s %s\033[39m\n\n", time.Since(event.StartTime), query)
	// 	})
	// }
}

//Get db connection
func Get() *pg.DB {
	return db
}

//Close db connection
func Close() {
	err := db.Close()

	if err != nil {
		log.Printf("Closing DB err", err)
	}
	log.Printf("DB closed")
}

// type dbLogger struct{}

// func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
// 	return c, nil
// }

// func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
// 	fmt.Println(q.FormattedQuery())
// 	return nil
// }

type dbLogger struct{}

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {
}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	// fmt.Println(q.FormattedQuery())
}
