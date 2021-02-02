package saveAssessmentService

import (
	"cyberliver/alcochange-dtx/dtos"
	"cyberliver/alcochange-dtx/sentryaccounts"
)

// HealthConditionAssessment will insert the data into DB
func (s *SaveAssessment) HealthConditionAssessment(req dtos.HealthConditionAssessmentAnswer, userID int64) error {

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
		} else {
			assessmentHeader.UserID = userID
			assessmentHeader.AldHealthConditionQuestionID = healthCondUserAnswer.QuestionID
			assessmentHeader.AldHealthConditionOptionID = healthCondUserAnswer.OptionID
			assessmentHeader.Points = healthCondUserAnswer.Points
			assessmentHeader.MaxPoints = healthCondUserAnswer.MaxPoints
			errHealthInsert := s.saveAssessmentDao.SaveHealthConditionAssessment(assessmentHeader)
			if errHealthInsert != nil {
				s.l.Error("HealthConditionAssessment Error--", errHealthInsert)
				sentryaccounts.SentryLogExceptions(errHealthInsert)

				return errHealthInsert
			}
		}

		errHealthInsertLog := s.saveAssessmentDao.SaveHealthConditionAssessmentLog(assessmentHeader)
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

	for _, userAnswer := range req.UserAnswer {
		assessmentHeader := s.saveAssessmentDao.IsExistsAuditAssessment(userID, userAnswer.QuestionID)
		if assessmentHeader.ID > 0 {
			assessmentHeader.AldAuditAssessmentOptionID = userAnswer.OptionID
			assessmentHeader.Points = userAnswer.Points
			assessmentHeader.MaxPoints = userAnswer.MaxPoints
			assessmentHeader.AvgPonits = float64((userAnswer.Points / float64(userAnswer.MaxPoints)) * 100)
			errAuditUpt := s.saveAssessmentDao.UpdateAuditAssessment(assessmentHeader)
			if errAuditUpt != nil {
				s.l.Error("AuditAssessment Error--", errAuditUpt)
				sentryaccounts.SentryLogExceptions(errAuditUpt)
				return errAuditUpt
			}
		} else {
			assessmentHeader.UserID = userID
			assessmentHeader.AldAuditAssessmentQuestionID = userAnswer.QuestionID
			assessmentHeader.AldAuditAssessmentOptionID = userAnswer.OptionID
			assessmentHeader.Points = userAnswer.Points
			assessmentHeader.MaxPoints = userAnswer.MaxPoints
			assessmentHeader.AvgPonits = float64((userAnswer.Points / float64(userAnswer.MaxPoints)) * 100)
			errAuditInsert := s.saveAssessmentDao.SaveAuditAssessment(assessmentHeader)
			if errAuditInsert != nil {
				s.l.Error("AuditAssessment Error--", errAuditInsert)
				sentryaccounts.SentryLogExceptions(errAuditInsert)
				return errAuditInsert
			}
		}

		errAuditInsertLog := s.saveAssessmentDao.SaveAuditAssessmentLog(assessmentHeader)
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

	for _, userAnswer := range req.UserAnswer {
		assessmentHeader := s.saveAssessmentDao.IsExistsGoalSettingAssessment(userID, userAnswer.QuestionID)
		if assessmentHeader.ID > 0 {
			assessmentHeader.AldGoalSettingAssessmentOptionID = userAnswer.OptionID
			assessmentHeader.Points = userAnswer.Points
			assessmentHeader.MaxPoints = userAnswer.MaxPoints
			assessmentHeader.AvgPonits = float64((userAnswer.Points / float64(userAnswer.MaxPoints)) * 100)
			errGoalUpt := s.saveAssessmentDao.UpdateGoalSettingAssessment(assessmentHeader)
			if errGoalUpt != nil {
				s.l.Error("GoalSettingAssessment Error--", errGoalUpt)
				sentryaccounts.SentryLogExceptions(errGoalUpt)
				return errGoalUpt
			}
		} else {
			assessmentHeader.UserID = userID
			assessmentHeader.AldGoalSettingAssessmentQuestionID = userAnswer.QuestionID
			assessmentHeader.AldGoalSettingAssessmentOptionID = userAnswer.OptionID
			assessmentHeader.Points = userAnswer.Points
			assessmentHeader.MaxPoints = userAnswer.MaxPoints
			assessmentHeader.AvgPonits = float64((userAnswer.Points / float64(userAnswer.MaxPoints)) * 100)
			errGoalInsert := s.saveAssessmentDao.SaveGoalSettingAssessment(assessmentHeader)
			if errGoalInsert != nil {
				s.l.Error("GoalSettingAssessment Error--", errGoalInsert)
				sentryaccounts.SentryLogExceptions(errGoalInsert)
				return errGoalInsert
			}
		}

		errGoalInsertLog := s.saveAssessmentDao.SaveGoalSettingAssessmentLog(assessmentHeader)
		if errGoalInsertLog != nil {
			s.l.Error("GoalSettingAssessment Error--", errGoalInsertLog)
			sentryaccounts.SentryLogExceptions(errGoalInsertLog)
			return errGoalInsertLog
		}

	}

	return nil

}
