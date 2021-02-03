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
	createAuditAssesmentQuestion(db)
	createAudiAssesmentQuestionOptions(db)
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
		db.Model(healthRow).Where("question_no = ?", health.QuestionNo).Select()
		if healthRow.ID == 0 {
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
			log.Println("Health Assesment Question Options created successfully.")
		}
	}
}

func createAuditAssesmentQuestion(db *pg.DB) {
	auditQuestions := []models.AldAuditAssessmentQuestion{
		{
			Question:             "How often do you have a drink containing alcohol?",
			SequenceOrder:        1,
			QuestionNo:           1,
			QuestionOptionTypeID: 1,
		},
		{
			Question:             "How many units of alcohol do you drink on a typical day when you are drinking?",
			SequenceOrder:        2,
			QuestionNo:           2,
			QuestionOptionTypeID: 1,
		},
		{
			Question:             "How often have you had 6 or more units if female, or 8 or more if male, on a single occasion in the last year?",
			SequenceOrder:        3,
			QuestionNo:           3,
			QuestionOptionTypeID: 1,
		},
		{
			Question:             "How often during the last year have you found that you were not able to stop drinking once you had started?",
			SequenceOrder:        4,
			QuestionNo:           4,
			QuestionOptionTypeID: 1,
		},
		{
			Question:             "How often during the last year have you needed an alcoholic drink in the morning to get yourself going after a heavy drinking session?",
			SequenceOrder:        5,
			QuestionNo:           5,
			QuestionOptionTypeID: 1,
		},
		{
			Question:             "How often during the last year have you needed an alcoholic drink in the morning to get yourself going after a heavy drinking session?",
			SequenceOrder:        6,
			QuestionNo:           6,
			QuestionOptionTypeID: 1,
		},
		{
			Question:             "How often during the last year have you had a feeling of guilt or remorse after drinking?",
			SequenceOrder:        7,
			QuestionNo:           7,
			QuestionOptionTypeID: 1,
		},
		{
			Question:             "How often during the last year have you been unable to remember what happened the night before because you had been drinking?",
			SequenceOrder:        8,
			QuestionNo:           8,
			QuestionOptionTypeID: 1,
		},
		{
			Question:             "Have you or someone else been injured as a result of your drinking?",
			SequenceOrder:        9,
			QuestionNo:           9,
			QuestionOptionTypeID: 1,
		},
		{
			Question:             "Has a relative or friend or a doctor or another health worker been concerned about your drinking or suggested you cut down?",
			SequenceOrder:        10,
			QuestionNo:           10,
			QuestionOptionTypeID: 1,
		},
	}
	for _, auditRow := range auditQuestions {
		audit := &models.AldAuditAssessmentQuestion{}
		db.Model(audit).Where("question_no = ?", auditRow.QuestionNo).Select()
		if audit.ID == 0 {
			auditRow.BeforeInsert("")
			if _, err := db.Model(&auditRow).Insert(); err != nil {
				log.Println("Error to insert default createAuditAssesmentQuestion.", err.Error())
				return
			}
			log.Println("Audit Assesment Question created successfully.")
		}
	}
}

func createAudiAssesmentQuestionOptions(db *pg.DB) {
	healthQuestions := []models.AldAuditAssessmentOption{
		{
			Name:                         "Never",
			SequenceOrder:                1,
			AldAuditAssessmentQuestionID: 1,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Monthly or less",
			SequenceOrder:                2,
			AldAuditAssessmentQuestionID: 1,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "2 to 4 times a MONTH",
			SequenceOrder:                3,
			AldAuditAssessmentQuestionID: 1,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "2 to 3 times a WEEK",
			SequenceOrder:                4,
			AldAuditAssessmentQuestionID: 1,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "4 or more times a week",
			SequenceOrder:                5,
			AldAuditAssessmentQuestionID: 1,
			Points:                       0,
			MaxPoints:                    10,
		},

		// 2nd Question options
		{
			Name:                         "1 or 2 drinks",
			SequenceOrder:                1,
			AldAuditAssessmentQuestionID: 2,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "3 or 4 drinks",
			SequenceOrder:                2,
			AldAuditAssessmentQuestionID: 2,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "5 or 6 drinks",
			SequenceOrder:                3,
			AldAuditAssessmentQuestionID: 2,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "7 or 8 or 9 drinks",
			SequenceOrder:                4,
			AldAuditAssessmentQuestionID: 2,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "10 or more drinks",
			SequenceOrder:                5,
			AldAuditAssessmentQuestionID: 2,
			Points:                       0,
			MaxPoints:                    10,
		},

		//3rd Question options
		{
			Name:                         "Never",
			SequenceOrder:                1,
			AldAuditAssessmentQuestionID: 3,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Less than monthly",
			SequenceOrder:                2,
			AldAuditAssessmentQuestionID: 3,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Monthly",
			SequenceOrder:                3,
			AldAuditAssessmentQuestionID: 3,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Weekly",
			SequenceOrder:                4,
			AldAuditAssessmentQuestionID: 3,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Daily or almost daily",
			SequenceOrder:                5,
			AldAuditAssessmentQuestionID: 3,
			Points:                       0,
			MaxPoints:                    10,
		},

		//4th Question options
		{
			Name:                         "Never",
			SequenceOrder:                1,
			AldAuditAssessmentQuestionID: 4,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Less than monthly",
			SequenceOrder:                2,
			AldAuditAssessmentQuestionID: 4,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Monthly",
			SequenceOrder:                3,
			AldAuditAssessmentQuestionID: 4,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Weekly",
			SequenceOrder:                4,
			AldAuditAssessmentQuestionID: 4,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Daily or almost daily",
			SequenceOrder:                5,
			AldAuditAssessmentQuestionID: 4,
			Points:                       0,
			MaxPoints:                    10,
		},

		//5th Question options
		{
			Name:                         "Never",
			SequenceOrder:                1,
			AldAuditAssessmentQuestionID: 5,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Less than monthly",
			SequenceOrder:                2,
			AldAuditAssessmentQuestionID: 5,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Monthly",
			SequenceOrder:                3,
			AldAuditAssessmentQuestionID: 5,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Weekly",
			SequenceOrder:                4,
			AldAuditAssessmentQuestionID: 5,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Daily or almost daily",
			SequenceOrder:                5,
			AldAuditAssessmentQuestionID: 5,
			Points:                       0,
			MaxPoints:                    10,
		},

		//6th Question options
		{
			Name:                         "Never",
			SequenceOrder:                1,
			AldAuditAssessmentQuestionID: 6,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Less than monthly",
			SequenceOrder:                2,
			AldAuditAssessmentQuestionID: 6,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Monthly",
			SequenceOrder:                3,
			AldAuditAssessmentQuestionID: 6,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Weekly",
			SequenceOrder:                4,
			AldAuditAssessmentQuestionID: 6,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Daily or almost daily",
			SequenceOrder:                5,
			AldAuditAssessmentQuestionID: 6,
			Points:                       0,
			MaxPoints:                    10,
		},

		//7th Question options
		{
			Name:                         "Never",
			SequenceOrder:                1,
			AldAuditAssessmentQuestionID: 7,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Less than monthly",
			SequenceOrder:                2,
			AldAuditAssessmentQuestionID: 7,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Monthly",
			SequenceOrder:                3,
			AldAuditAssessmentQuestionID: 7,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Weekly",
			SequenceOrder:                4,
			AldAuditAssessmentQuestionID: 7,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Daily or almost daily",
			SequenceOrder:                5,
			AldAuditAssessmentQuestionID: 7,
			Points:                       0,
			MaxPoints:                    10,
		},

		//8th Question options
		{
			Name:                         "Never",
			SequenceOrder:                1,
			AldAuditAssessmentQuestionID: 8,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Less than monthly",
			SequenceOrder:                2,
			AldAuditAssessmentQuestionID: 8,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Monthly",
			SequenceOrder:                3,
			AldAuditAssessmentQuestionID: 8,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Weekly",
			SequenceOrder:                4,
			AldAuditAssessmentQuestionID: 8,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Daily or almost daily",
			SequenceOrder:                5,
			AldAuditAssessmentQuestionID: 8,
			Points:                       0,
			MaxPoints:                    10,
		},

		//9th Question options
		{
			Name:                         "No, never",
			SequenceOrder:                1,
			AldAuditAssessmentQuestionID: 9,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Yes, but not in the last year",
			SequenceOrder:                2,
			AldAuditAssessmentQuestionID: 9,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Yes, during the last year",
			SequenceOrder:                3,
			AldAuditAssessmentQuestionID: 9,
			Points:                       0,
			MaxPoints:                    10,
		},

		//10th Question options
		{
			Name:                         "No, never",
			SequenceOrder:                1,
			AldAuditAssessmentQuestionID: 10,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Yes, but not in the last year",
			SequenceOrder:                2,
			AldAuditAssessmentQuestionID: 10,
			Points:                       0,
			MaxPoints:                    10,
		},
		{
			Name:                         "Yes, during the last year",
			SequenceOrder:                3,
			AldAuditAssessmentQuestionID: 10,
			Points:                       0,
			MaxPoints:                    10,
		},
	}
	for _, auditOption := range healthQuestions {
		auditRow := &models.AldAuditAssessmentOption{}
		db.Model(auditRow).Where("ald_audit_assessment_question_id=? AND sequence_order=?", auditOption.AldAuditAssessmentQuestionID, auditOption.SequenceOrder).Select()
		if auditRow.ID == 0 {
			auditOption.BeforeInsert("")
			if _, err := db.Model(&auditOption).Insert(); err != nil {
				log.Println("Error to insert default createAudiAssesmentQuestionOptions.", err.Error())
				return
			}
			log.Println("Audit Assesment Question Options created successfully.")
		}
	}
}
