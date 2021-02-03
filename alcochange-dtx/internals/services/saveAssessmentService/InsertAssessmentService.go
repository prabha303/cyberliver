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

		errGoalInsertLog := s.saveAssessmentDao.SaveGoalSettingAssessmentLog(goalSettingLog)
		if errGoalInsertLog != nil {
			s.l.Error("GoalSettingAssessment Error--", errGoalInsertLog)
			sentryaccounts.SentryLogExceptions(errGoalInsertLog)
			return errGoalInsertLog
		}

	}

	return nil

}
