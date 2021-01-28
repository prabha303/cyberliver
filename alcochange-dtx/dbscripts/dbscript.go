package dbscripts

import (
	"fmt"
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
		&models.WarningLabel{},
		&models.UserActionConfirmation{},
		&models.AlcoChangeTermsAndPrivacy{},
		&models.PatientAccessCode{},
		&models.Constants{},
		&models.AldHealthConditionQuestion{},
		&models.AldHealthConditionOption{},
		&models.AldAuditAssessmentQuestion{},
		&models.AldAuditAssessmentOption{},
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

func CreateIndex(db *pg.DB) {
	//TODO: add your indexing code here..
	for _, i := range []string{
		fmt.Sprintf("CREATE UNIQUE INDEX IF NOT EXISTS uuid_eid_unique ON %s (device_uuid, email_id)", "user_action_confirmations"),
	} {
		if _, err := db.Exec(i); err != nil {
			log.Printf("Error in creating the index %s", err.Error())
		} else {
			log.Printf("CreateIndex success")
		}
	}
}
