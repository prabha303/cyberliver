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
	createRole()
	createProductAccess()
}

//getModels function use to get all the masters from models
func getModels() []interface{} {
	return []interface{}{
		//Application  Masters
		&models.UserType{},
		&models.User{},
		&models.Country{},
		&models.WarningLabel{},
		&models.UserActionConfirmation{},
		&models.AlcoChangeTermsAndPrivacy{},
		&models.PatientAccessCode{},
		&models.ProductAccess{},
		&models.LoginDeviceDetails{},
		&models.LoginLogs{},
		&models.UserAccess{},
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
		fmt.Sprintf("CREATE UNIQUE INDEX IF NOT EXISTS uuid_eid_unique ON %s (device_uuid, user_id)", "user_action_confirmations"),
		fmt.Sprintf("CREATE UNIQUE INDEX IF NOT EXISTS idx_code_name_role ON %s (name, code)", "user_types"),
		fmt.Sprintf("CREATE UNIQUE INDEX IF NOT EXISTS idx_code_name_pa ON %s (name, code)", "product_accesses"),
	} {
		if _, err := db.Exec(i); err != nil {
			log.Printf("Error in creating the index %s", err.Error())
		} else {
			//log.Printf("CreateIndex success")
		}
	}
}

func createRole() {
	db := dbcon.Get()
	userTypeC := []models.UserType{
		{
			Name:        "patient",
			Code:        "PATIENT",
			Description: "Patient login",
			IsActive:    true,
		},
		{
			Name:        "Others",
			Code:        "OTHERS",
			Description: "Others role for user",
			IsActive:    true,
		},
	}
	for _, rowData := range userTypeC {
		userType := &models.UserType{}
		db.Model(userType).Where("LOWER(code) = LOWER(?)", rowData.Code).Select()
		if userType.ID == 0 {
			rowData.BeforeInsert("")
			if _, err := db.Model(&rowData).Insert(); err != nil {
				log.Println("Error to insert default user_types.", err.Error())
				return
			}
			log.Println("User Types created successfully.")
		}
	}

}

func createProductAccess() {
	db := dbcon.Get()
	pAccess := []models.ProductAccess{
		{
			Name:     "AlcoChange ",
			Code:     "ALCOCHANGE-DTX",
			IsActive: true,
		},
	}
	for _, accessData := range pAccess {
		access := &models.ProductAccess{}
		db.Model(access).Where("LOWER(code) = LOWER(?)", &accessData.Code).Select()
		if access.ID == 0 {
			accessData.BeforeInsert("")
			if _, err := db.Model(&accessData).Insert(); err != nil {
				log.Println("Error to insert default ProductAccess.", err.Error())
				return
			}
			log.Println("ProductAccess created successfully.")
		}
	}

}
