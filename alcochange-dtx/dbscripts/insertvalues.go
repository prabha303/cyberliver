package dbscripts

import (
	"cyberliver/alcochange-dtx/models"
	"log"

	"github.com/go-pg/pg"
)

func InsertValues(db *pg.DB) {
	createUserType(db)
	createProductAccess(db)
	createQuestionOptionType(db)
	createHealthAssesmentQuestion(db)
	createHealthAssesmentQuestionOptions(db)
}

func createUserType(db *pg.DB) {
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

func createQuestionOptionType(db *pg.DB) {
	questionOptionIns := []models.QuestionOptionType{
		{
			Name:     "Radio",
			Code:     "SINGLE",
			IsActive: true,
		},
		{
			Name:     "Checkbox",
			Code:     "MULTIPLE",
			IsActive: true,
		},
		{
			Name:     "Likert Scal",
			Code:     "SINGLE_LIKERT`",
			IsActive: true,
		},
	}
	for _, rowData := range questionOptionIns {
		quesOptionType := &models.QuestionOptionType{}
		db.Model(quesOptionType).Where("LOWER(code) = LOWER(?)", rowData.Code).Select()
		if quesOptionType.ID == 0 {
			rowData.BeforeInsert("")
			if _, err := db.Model(&rowData).Insert(); err != nil {
				log.Println("Error to insert default question_option_type.", err.Error())
				return
			}
			log.Println("Question Option Type created successfully.")
		}
	}
}

func createProductAccess(db *pg.DB) {
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

func createHealthAssesmentQuestion(db *pg.DB) {
	healthQuestions := []models.AldHealthConditionQuestion{
		{
			Question:             "How much are these health conditions bothering you?",
			QuestionNo:           1,
			QuestionOptionTypeID: 3,
			SequenceOrder:        1,
		},
	}
	for _, health := range healthQuestions {
		healthRow := &models.AldHealthConditionQuestion{}
		db.Model(healthRow).Where("id = ?", health.ID).Select()
		if healthRow.QuestionNo != 1 {
			health.BeforeInsert("")
			if _, err := db.Model(&health).Insert(); err != nil {
				log.Println("Error to insert default createHealthAssesmentQuestion.", err.Error())
				return
			}
			log.Println("Health Assesment Question created successfully.")
		}
	}
}
func createHealthAssesmentQuestionOptions(db *pg.DB) {
	healthQuestions := []models.AldHealthConditionOption{
		{
			Name:                         "Jaundice (yellow eyes or skin)",
			AldHealthConditionQuestionID: 1,
			SequenceOrder:                1,
			Points:                       0.0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Ascites (swelling of abdomen or tummy)",
			AldHealthConditionQuestionID: 1,
			SequenceOrder:                2,
			Points:                       0.0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Swelling of legs",
			AldHealthConditionQuestionID: 1,
			SequenceOrder:                3,
			Points:                       0.0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Low energy levels",
			AldHealthConditionQuestionID: 1,
			SequenceOrder:                4,
			Points:                       0.0,
			MaxPoints:                    10,
		},
	}
	for _, health := range healthQuestions {
		healthRow := models.AldHealthConditionOption{}
		db.Model(&healthRow).Where("ald_health_condition_question_id=? AND sequence_order=?", health.AldHealthConditionQuestionID, health.SequenceOrder).Select()
		if healthRow.ID == 0 {
			health.BeforeInsert("")
			if _, err := db.Model(&health).Insert(); err != nil {
				log.Println("Error to insert default createHealthAssesmentQuestionOptions.", err.Error())
				return
			}
			log.Println("create Health Assesment Question Options created successfully.")
		}
	}
}
