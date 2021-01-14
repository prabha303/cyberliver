package dbscripts

import (
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"

	"ecargoware/alcochange-dtx/dbcon"
	"ecargoware/alcochange-dtx/models"
)

//InitDB initialize DB
func InitDB() {
	db := dbcon.Get()
	CreateTables(db)
	MigrateTables(db)
	CreateIndex(db)
}

//getModels function use to get all the masters from models
func getModels() []interface{} {
	return []interface{}{
		//Application  Masters
		&models.Country{},
	}
}

//CreateTables function is use to create master tables
func CreateTables(db *pg.DB) {
	for _, mod := range getModels() {
		// db := dbcon.Get()

		if err := db.CreateTable(mod, &orm.CreateTableOptions{
			IfNotExists:   true,
			FKConstraints: true,
		}); err != nil {
			log.Printf("Error in creating tables, err:%s", err.Error())
		}
	}
}

//MigrateTables function is use to migrate table mentioned queries
func MigrateTables(db *pg.DB) {
	// db := dbcon.Get()
	migrationQueries := []string{}

	for _, q := range migrationQueries {
		if _, err := db.Exec(q); err != nil {
			log.Println("Error in migration", err.Error())
		}
	}
}

//DropTables function
func DropTables(db *pg.DB) {
	// db := dbcon.Get()
	dropTableQueries := []string{}

	for _, q := range dropTableQueries {
		if _, err := db.Exec(q); err != nil {
			log.Println("Error in drop table query", err.Error())
		}
	}
}

//CreateIndex indexing
func CreateIndex(db *pg.DB) {
	//TODO: add your indexing code here..
	for _, i := range []string{} {
		if _, err := db.Exec(i); err != nil {
			log.Printf("Error in creating the index %s", err.Error())
		}
	}
}
