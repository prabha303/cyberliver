package saveAssessmentService

import (
	"cyberliver/alcochange-dtx/dtos"
	"cyberliver/alcochange-dtx/models"
	"cyberliver/alcochange-dtx/sentryaccounts"
)

// HealthConditionAssessment will insert the data into DB
func (s *SaveAssessment) HealthConditionAssessment(req dtos.HealthConditionAssessmentAnswer, userID int64) error {

	healthConditionLog := models.AldHealthAssessmentLog{}
	for _, healthCondUserAnswer := range req.UserAnswer {
		assessmentHeader := s.saveAssessmentDao.IsExistsHealthConditionAssessment(userID, healthCondUserAnswer.QuestionID)
		if assessmentHeader.ID > 0 {
			assessmentHeader.AldHealthConditionOptionID = healthCondUserAnswer.OptionID
			assessmentHeader.Points = healthCondUserAnswer.Points
			assessmentHeader.MaxPoints = healthCondUserAnswer.MaxPoints

			assessmentHeader.BeforeUpdate()

			errHealthUpt := s.saveAssessmentDao.UpdateHealthConditionAssessment(assessmentHeader)
			if errHealthUpt != nil {
				s.l.Error("HealthConditionAssessment Error--", errHealthUpt)
				sentryaccounts.SentryLogExceptions(errHealthUpt)
				return errHealthUpt
			}
			healthConditionLog.AldHealthAssessmentHeaderID = assessmentHeader.ID
		} else {
			assessmentHeader.UserID = userID
			assessmentHeader.AldHealthConditionQuestionID = healthCondUserAnswer.QuestionID
			assessmentHeader.AldHealthConditionOptionID = healthCondUserAnswer.OptionID
			assessmentHeader.Points = healthCondUserAnswer.Points
			assessmentHeader.MaxPoints = healthCondUserAnswer.MaxPoints

			assessmentHeader.BeforeInsert()

			hcIns, errHealthInsert := s.saveAssessmentDao.SaveHealthConditionAssessment(assessmentHeader)
			if errHealthInsert != nil {
				s.l.Error("HealthConditionAssessment Error--", errHealthInsert)
				sentryaccounts.SentryLogExceptions(errHealthInsert)

				return errHealthInsert
			}
			healthConditionLog.AldHealthAssessmentHeaderID = hcIns.ID
		}

		healthConditionLog.UserID = assessmentHeader.UserID
		healthConditionLog.AldHealthConditionQuestionID = assessmentHeader.AldHealthConditionQuestionID
		healthConditionLog.AldHealthConditionOptionID = assessmentHeader.AldHealthConditionOptionID
		healthConditionLog.Points = assessmentHeader.Points
		healthConditionLog.AvgPoints = assessmentHeader.AvgPoints
		healthConditionLog.MaxPoints = assessmentHeader.MaxPoints

		healthConditionLog.BeforeInsert()

		errHealthInsertLog := s.saveAssessmentDao.SaveHealthConditionAssessmentLog(healthConditionLog)
		if errHealthInsertLog != nil {
			s.l.Error("HealthConditionAssessment Error--", errHealthInsertLog)
			sentryaccounts.SentryLogExceptions(errHealthInsertLog)
			return errHealthInsertLog
		}

	}

	return nil

}

// AuditAssessment will insert the data into DB
func (s *SaveAssessment) AuditAssessment(req dtos.AuditAssessmentAnswer, userID int64) error {

	auditLog := models.AldAuditAssessmentLog{}

	for _, userAnswer := range req.UserAnswer {
		assessmentHeader := s.saveAssessmentDao.IsExistsAuditAssessment(userID, userAnswer.QuestionID)
		if assessmentHeader.ID > 0 {
			assessmentHeader.AldAuditAssessmentOptionID = userAnswer.OptionID
			assessmentHeader.Points = userAnswer.Points
			assessmentHeader.MaxPoints = userAnswer.MaxPoints
			assessmentHeader.AvgPoints = float64((userAnswer.Points / float64(userAnswer.MaxPoints)) * 100)

			assessmentHeader.BeforeUpdate()

			errAuditUpt := s.saveAssessmentDao.UpdateAuditAssessment(assessmentHeader)
			if errAuditUpt != nil {
				s.l.Error("AuditAssessment Error--", errAuditUpt)
				sentryaccounts.SentryLogExceptions(errAuditUpt)
				return errAuditUpt
			}
			auditLog.AldAuditAssessmentHeaderID = assessmentHeader.ID
		} else {
			assessmentHeader.UserID = userID
			assessmentHeader.AldAuditAssessmentQuestionID = userAnswer.QuestionID
			assessmentHeader.AldAuditAssessmentOptionID = userAnswer.OptionID
			assessmentHeader.Points = userAnswer.Points
			assessmentHeader.MaxPoints = userAnswer.MaxPoints
			assessmentHeader.AvgPoints = float64((userAnswer.Points / float64(userAnswer.MaxPoints)) * 100)

			assessmentHeader.BeforeInsert()

			auditIns, errAuditInsert := s.saveAssessmentDao.SaveAuditAssessment(assessmentHeader)
			if errAuditInsert != nil {
				s.l.Error("AuditAssessment Error--", errAuditInsert)
				sentryaccounts.SentryLogExceptions(errAuditInsert)
				return errAuditInsert
			}
			auditLog.AldAuditAssessmentHeaderID = auditIns.ID
		}

		auditLog.UserID = assessmentHeader.UserID
		auditLog.AldAuditAssessmentQuestionID = assessmentHeader.AldAuditAssessmentQuestionID
		auditLog.AldAuditAssessmentOptionID = assessmentHeader.AldAuditAssessmentOptionID
		auditLog.Points = assessmentHeader.Points
		auditLog.AvgPoints = assessmentHeader.AvgPoints
		auditLog.MaxPoints = assessmentHeader.MaxPoints

		auditLog.BeforeInsert()

		errAuditInsertLog := s.saveAssessmentDao.SaveAuditAssessmentLog(auditLog)
		if errAuditInsertLog != nil {
			s.l.Error("AuditAssessment Error--", errAuditInsertLog)
			sentryaccounts.SentryLogExceptions(errAuditInsertLog)
			return errAuditInsertLog
		}

	}

	return nil

}

// GoalSettingAssessment will insert the data into DB
func (s *SaveAssessment) GoalSettingAssessment(req dtos.GoalSettingAssessmentAnswer, userID int64) error {

	goalSettingLog := models.AldGoalSettingAssessmentLog{}

	for _, userAnswer := range req.UserAnswer {
		assessmentHeader := s.saveAssessmentDao.IsExistsGoalSettingAssessment(userID, userAnswer.QuestionID)
		if assessmentHeader.ID > 0 {
			assessmentHeader.AldGoalSettingAssessmentOptionID = userAnswer.OptionID
			assessmentHeader.Points = userAnswer.Points
			assessmentHeader.MaxPoints = userAnswer.MaxPoints
			assessmentHeader.AvgPoints = float64((userAnswer.Points / float64(userAnswer.MaxPoints)) * 100)

			assessmentHeader.BeforeUpdate()

			errGoalUpt := s.saveAssessmentDao.UpdateGoalSettingAssessment(assessmentHeader)
			if errGoalUpt != nil {
				s.l.Error("GoalSettingAssessment Error--", errGoalUpt)
				sentryaccounts.SentryLogExceptions(errGoalUpt)
				return errGoalUpt
			}
			goalSettingLog.AldGoalSettingAssessmentHeaderID = assessmentHeader.ID
		} else {
			assessmentHeader.UserID = userID
			assessmentHeader.AldGoalSettingAssessmentQuestionID = userAnswer.QuestionID
			assessmentHeader.AldGoalSettingAssessmentOptionID = userAnswer.OptionID
			assessmentHeader.Points = userAnswer.Points
			assessmentHeader.MaxPoints = userAnswer.MaxPoints
			assessmentHeader.AvgPoints = float64((userAnswer.Points / float64(userAnswer.MaxPoints)) * 100)

			assessmentHeader.BeforeInsert()

			goalSettingIns, errGoalInsert := s.saveAssessmentDao.SaveGoalSettingAssessment(assessmentHeader)
			if errGoalInsert != nil {
				s.l.Error("GoalSettingAssessment Error--", errGoalInsert)
				sentryaccounts.SentryLogExceptions(errGoalInsert)
				return errGoalInsert
			}
			goalSettingLog.AldGoalSettingAssessmentHeaderID = goalSettingIns.ID

		}

		goalSettingLog.UserID = assessmentHeader.UserID
		goalSettingLog.AldGoalSettingAssessmentQuestionID = assessmentHeader.AldGoalSettingAssessmentQuestionID
		goalSettingLog.AldGoalSettingAssessmentOptionID = assessmentHeader.AldGoalSettingAssessmentOptionID
		goalSettingLog.Points = assessmentHeader.Points
		goalSettingLog.AvgPoints = assessmentHeader.AvgPoints
		goalSettingLog.MaxPoints = assessmentHeader.MaxPoints

		goalSettingLog.BeforeInsert()

		errGoalInsertLog := s.saveAssessmentDao.SaveGoalSettingAssessmentLog(goalSettingLog)
		if errGoalInsertLog != nil {
			s.l.Error("GoalSettingAssessment Error--", errGoalInsertLog)
			sentryaccounts.SentryLogExceptions(errGoalInsertLog)
			return errGoalInsertLog
		}

	}

	return nil

}

// DrinkAssessment will insert the data into DB
func (s *SaveAssessment) DrinkAssessment(req dtos.DrinkHabitAssessmentAnswer, userID int64) error {

	err := s.DrinkHabitAssessment(req, userID)
	if err != nil {
		s.l.Error("DrinkAssessment Error--", err)
		sentryaccounts.SentryLogExceptions(err)
		return err
	}

	dherr := s.DrinkProfileAssessment(req, userID)
	if dherr != nil {
		s.l.Error("DrinkAssessment Error--", dherr)
		sentryaccounts.SentryLogExceptions(dherr)
		return dherr
	}

	return nil

}

func (s *SaveAssessment) DrinkProfileAssessment(req dtos.DrinkHabitAssessmentAnswer, userID int64) error {

	drinkProfileLog := models.AldDrinkProfileLog{}

	for _, userAnswer := range req.DrinkProfile {
		assessmentHeader := s.saveAssessmentDao.IsExistsDrinkProfileAssessment(userID, userAnswer.DrinkID)
		if assessmentHeader.ID > 0 {
			assessmentHeader.DrinkCount = userAnswer.DrinkCount
			assessmentHeader.Quantity = userAnswer.Quantity
			assessmentHeader.QuantityUnitID = userAnswer.QuantityUnitID
			assessmentHeader.Cost = userAnswer.Cost
			assessmentHeader.Calories = userAnswer.Calories

			assessmentHeader.BeforeUpdate()

			errDrinkUpt := s.saveAssessmentDao.UpdateDrinkProfileAssessment(assessmentHeader)
			if errDrinkUpt != nil {
				s.l.Error("DrinkProfileAssessment Error--", errDrinkUpt)
				sentryaccounts.SentryLogExceptions(errDrinkUpt)
				return errDrinkUpt
			}
			drinkProfileLog.AldDrinkProfileHeaderID = assessmentHeader.ID
		} else {
			assessmentHeader.UserID = userID
			assessmentHeader.DrinkID = userAnswer.DrinkID
			assessmentHeader.Name = userAnswer.Name
			assessmentHeader.DrinkCount = userAnswer.DrinkCount
			assessmentHeader.Quantity = userAnswer.Quantity
			assessmentHeader.QuantityUnitID = userAnswer.QuantityUnitID
			assessmentHeader.Cost = userAnswer.Cost
			assessmentHeader.Calories = userAnswer.Calories
			assessmentHeader.CountryID = userAnswer.CountryID

			assessmentHeader.BeforeInsert()

			drinkProfileIns, errDrinkInsert := s.saveAssessmentDao.SaveDrinkProfileAssessment(assessmentHeader)
			if errDrinkInsert != nil {
				s.l.Error("DrinkHabitAssessment Error--", errDrinkInsert)
				sentryaccounts.SentryLogExceptions(errDrinkInsert)
				return errDrinkInsert
			}
			drinkProfileLog.AldDrinkProfileHeaderID = drinkProfileIns.ID
		}

		drinkProfileLog.DrinkID = assessmentHeader.DrinkID
		drinkProfileLog.Name = assessmentHeader.Name
		drinkProfileLog.DrinkCount = assessmentHeader.DrinkCount
		drinkProfileLog.UserID = assessmentHeader.UserID
		drinkProfileLog.Quantity = assessmentHeader.Quantity
		drinkProfileLog.QuantityUnitID = assessmentHeader.QuantityUnitID
		drinkProfileLog.Cost = assessmentHeader.Cost
		drinkProfileLog.Calories = assessmentHeader.Calories
		drinkProfileLog.CountryID = userAnswer.CountryID

		drinkProfileLog.BeforeInsert()
		errDrinkInsertLog := s.saveAssessmentDao.SaveDrinkProfileAssessmentLog(drinkProfileLog)
		if errDrinkInsertLog != nil {
			s.l.Error("DrinkHabitAssessment Error--", errDrinkInsertLog)
			sentryaccounts.SentryLogExceptions(errDrinkInsertLog)
			return errDrinkInsertLog
		}

	}

	return nil
}

func (s *SaveAssessment) DrinkHabitAssessment(req dtos.DrinkHabitAssessmentAnswer, userID int64) error {

	drinkHabitLog := models.AldDrinkHabitAssessmentLog{}

	for _, userAnswer := range req.DrinkHabitAssessment.UserAnswer {
		assessmentHeader := s.saveAssessmentDao.IsExistsDrinkHabitAssessment(userID, userAnswer.QuestionID)
		if assessmentHeader.ID > 0 {
			assessmentHeader.AldDrinkHabitAssessmentOptionID = userAnswer.OptionID
			assessmentHeader.Points = userAnswer.Points
			assessmentHeader.MaxPoints = userAnswer.MaxPoints
			assessmentHeader.AvgPoints = float64((userAnswer.Points / float64(userAnswer.MaxPoints)) * 100)

			assessmentHeader.BeforeUpdate()

			errDrinkUpt := s.saveAssessmentDao.UpdateDrinkHabitAssessment(assessmentHeader)
			if errDrinkUpt != nil {
				s.l.Error("DrinkHabitAssessment Error--", errDrinkUpt)
				sentryaccounts.SentryLogExceptions(errDrinkUpt)
				return errDrinkUpt
			}
			drinkHabitLog.AldDrinkHabitAssessmentHeaderID = assessmentHeader.ID
		} else {
			assessmentHeader.UserID = userID
			assessmentHeader.AldDrinkHabitAssessmentQuestionID = userAnswer.QuestionID
			assessmentHeader.AldDrinkHabitAssessmentOptionID = userAnswer.OptionID
			assessmentHeader.Points = userAnswer.Points
			assessmentHeader.MaxPoints = userAnswer.MaxPoints
			assessmentHeader.AvgPoints = float64((userAnswer.Points / float64(userAnswer.MaxPoints)) * 100)

			assessmentHeader.BeforeInsert()

			drinkHabitIns, errDrinkInsert := s.saveAssessmentDao.SaveDrinkHabitAssessment(assessmentHeader)
			if errDrinkInsert != nil {
				s.l.Error("DrinkHabitAssessment Error--", errDrinkInsert)
				sentryaccounts.SentryLogExceptions(errDrinkInsert)
				return errDrinkInsert
			}
			drinkHabitLog.AldDrinkHabitAssessmentHeaderID = drinkHabitIns.ID

		}

		drinkHabitLog.UserID = assessmentHeader.UserID
		drinkHabitLog.AldDrinkHabitAssessmentQuestionID = assessmentHeader.AldDrinkHabitAssessmentQuestionID
		drinkHabitLog.AldDrinkHabitAssessmentOptionID = assessmentHeader.AldDrinkHabitAssessmentOptionID
		drinkHabitLog.Points = assessmentHeader.Points
		drinkHabitLog.AvgPoints = assessmentHeader.AvgPoints
		drinkHabitLog.MaxPoints = assessmentHeader.MaxPoints

		drinkHabitLog.BeforeInsert()

		errDrinkInsertLog := s.saveAssessmentDao.SaveDrinkHabitAssessmentLog(drinkHabitLog)
		if errDrinkInsertLog != nil {
			s.l.Error("DrinkHabitAssessment Error--", errDrinkInsertLog)
			sentryaccounts.SentryLogExceptions(errDrinkInsertLog)
			return errDrinkInsertLog
		}

	}

	return nil

}

// SupportiveContactAssessment will insert the data into DB
func (s *SaveAssessment) SupportiveContactAssessment(req dtos.SupportiveContact, userID int64) error {

	supContactLog := models.AldSupportiveContactLog{}
	supContactHeader := models.AldSupportiveContactHeader{}

	supContactHeader.UserID = userID

	isExist := s.saveAssessmentDao.IsExistsSupportiveContact(userID)
	if isExist {
		errSCUpt := s.saveAssessmentDao.DeleteSupportiveContact(userID)
		if errSCUpt != nil {
			s.l.Error("SupportiveContactAssessment Error--", errSCUpt)
			sentryaccounts.SentryLogExceptions(errSCUpt)
			return errSCUpt
		}
	}

	for _, contact := range req.Contacts {
		supContactHeader.Name = contact.Name
		supContactHeader.ContactNumber = contact.ContactNumber
		supContactHeader.AldRelationShipID = contact.RelationShipID

		supContactHeader.BeforeInsert()

		supportiveContactIns, errSCInsert := s.saveAssessmentDao.SaveSupportiveContact(supContactHeader)
		if errSCInsert != nil {
			s.l.Error("SupportiveContactAssessment Error--", errSCInsert)
			sentryaccounts.SentryLogExceptions(errSCInsert)
			return errSCInsert
		}
		supContactLog.AldSupportiveContactHeaderID = supportiveContactIns.ID

		supContactLog.Name = supContactHeader.Name
		supContactLog.ContactNumber = supContactHeader.ContactNumber
		supContactLog.AldRelationShipID = supContactHeader.AldRelationShipID
		supContactLog.UserID = supContactHeader.UserID

		supContactLog.BeforeInsert()

		errSCLog := s.saveAssessmentDao.SaveSupportiveContactLog(supContactLog)
		if errSCLog != nil {
			s.l.Error("SupportiveContactAssessment Error--", errSCLog)
			sentryaccounts.SentryLogExceptions(errSCLog)
			return errSCLog
		}

	}

	return nil

}
