package daos

import (
	"cyberliver/alcochange-dtx/models"
	"encoding/json"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type SaveAssessment struct {
	l      *log.Logger
	dbConn *pg.DB
}

func NewSaveAssessmentDB(l *log.Logger, dbConn *pg.DB) *SaveAssessment {
	return &SaveAssessment{
		l:      l,
		dbConn: dbConn,
	}
}

type SaveAssessmentDao interface {
	IsExistsHealthConditionAssessment(userID int64, questionID int64) models.AldHealthAssessmentHeader
	SaveHealthConditionAssessment(healthAssessmentResp models.AldHealthAssessmentHeader) (models.AldHealthAssessmentHeader, error)
	UpdateHealthConditionAssessment(healthAssessmentResp models.AldHealthAssessmentHeader) error
	SaveHealthConditionAssessmentLog(healthAssessmentResp models.AldHealthAssessmentLog) error
	IsExistsAuditAssessment(userID int64, questionID int64) models.AldAuditAssessmentHeader
	SaveAuditAssessment(healthAssessmentResp models.AldAuditAssessmentHeader) (models.AldAuditAssessmentHeader, error)
	UpdateAuditAssessment(healthAssessmentResp models.AldAuditAssessmentHeader) error
	SaveAuditAssessmentLog(healthAssessmentResp models.AldAuditAssessmentLog) error
	IsExistsGoalSettingAssessment(userID int64, questionID int64) models.AldGoalSettingAssessmentHeader
	SaveGoalSettingAssessment(goalAssessmentResp models.AldGoalSettingAssessmentHeader) (models.AldGoalSettingAssessmentHeader, error)
	UpdateGoalSettingAssessment(goalAssessmentResp models.AldGoalSettingAssessmentHeader) error
	SaveGoalSettingAssessmentLog(goalAssessmentResp models.AldGoalSettingAssessmentLog) error
}

func (s *SaveAssessment) IsExistsHealthConditionAssessment(userID int64, questionID int64) models.AldHealthAssessmentHeader {
	assessmentHeader := models.AldHealthAssessmentHeader{}
	s.dbConn.Model(&assessmentHeader).Where("user_id = ? AND ald_health_condition_question_id = ? ", userID, questionID).Select()
	return assessmentHeader
}

func (s *SaveAssessment) SaveHealthConditionAssessment(healthAssessmentResp models.AldHealthAssessmentHeader) (models.AldHealthAssessmentHeader, error) {
	_, insErr := s.dbConn.Model(&healthAssessmentResp).Insert()

	if insErr != nil {
		s.l.Error("SaveHealthConditionAssessment Error--", insErr)
		return healthAssessmentResp, insErr
	}

	dataBytes, _ := json.Marshal(healthAssessmentResp)
	s.l.Debug("User table json : %q", string(dataBytes))

	return healthAssessmentResp, nil
}

func (s *SaveAssessment) UpdateHealthConditionAssessment(healthAssessmentResp models.AldHealthAssessmentHeader) error {
	_, uptErr := s.dbConn.Model(&healthAssessmentResp).Column("ald_health_condition_option_id", "points", "max_points").Where("id=?", healthAssessmentResp.ID).Update()
	if uptErr != nil {
		s.l.Error("UpdateHealthConditionAssessment Error--", uptErr)
		return uptErr
	}

	dataBytes, _ := json.Marshal(healthAssessmentResp)
	s.l.Debug("User table json : %q", string(dataBytes))

	return nil
}

func (s *SaveAssessment) SaveHealthConditionAssessmentLog(healthAssessmentResp models.AldHealthAssessmentLog) error {

	_, insErr := s.dbConn.Model(&healthAssessmentResp).Insert()

	if insErr != nil {
		s.l.Error("SaveHealthConditionAssessmentLog Error--", insErr)
		return insErr
	}

	dataBytes, _ := json.Marshal(healthAssessmentResp)
	s.l.Debug("User table json : %q", string(dataBytes))

	return nil
}

func (s *SaveAssessment) IsExistsAuditAssessment(userID int64, questionID int64) models.AldAuditAssessmentHeader {
	assessmentHeader := models.AldAuditAssessmentHeader{}
	s.dbConn.Model(&assessmentHeader).Where("user_id = ? AND ald_audit_assessment_question_id = ? ", userID, questionID).Select()
	return assessmentHeader
}

func (s *SaveAssessment) SaveAuditAssessment(auditAssessmentResp models.AldAuditAssessmentHeader) (models.AldAuditAssessmentHeader, error) {
	_, insErr := s.dbConn.Model(&auditAssessmentResp).Insert()

	if insErr != nil {
		s.l.Error("SaveAuditAssessment Error--", insErr)
		return auditAssessmentResp, insErr
	}

	dataBytes, _ := json.Marshal(auditAssessmentResp)
	s.l.Debug("User table json : %q", string(dataBytes))

	return auditAssessmentResp, nil
}

func (s *SaveAssessment) UpdateAuditAssessment(auditAssessmentResp models.AldAuditAssessmentHeader) error {
	_, uptErr := s.dbConn.Model(&auditAssessmentResp).Column("ald_audit_assessment_option_id", "points", "max_points").Where("id=?", auditAssessmentResp.ID).Update()
	if uptErr != nil {
		s.l.Error("UpdateAuditAssessment Error--", uptErr)
		return uptErr
	}

	dataBytes, _ := json.Marshal(auditAssessmentResp)
	s.l.Debug("User table json : %q", string(dataBytes))

	return nil
}

func (s *SaveAssessment) SaveAuditAssessmentLog(auditAssessmentResp models.AldAuditAssessmentLog) error {
	_, insErr := s.dbConn.Model(&auditAssessmentResp).Insert()

	if insErr != nil {
		s.l.Error("SaveAuditAssessmentLog Error--", insErr)
		return insErr
	}

	dataBytes, _ := json.Marshal(auditAssessmentResp)
	s.l.Debug("User table json : %q", string(dataBytes))

	return nil
}

func (s *SaveAssessment) IsExistsGoalSettingAssessment(userID int64, questionID int64) models.AldGoalSettingAssessmentHeader {
	assessmentHeader := models.AldGoalSettingAssessmentHeader{}
	s.dbConn.Model(&assessmentHeader).Where("user_id = ? AND ald_goal_setting_assessment_question_id = ? ", userID, questionID).Select()
	return assessmentHeader
}

func (s *SaveAssessment) SaveGoalSettingAssessment(goalAssessmentResp models.AldGoalSettingAssessmentHeader) (models.AldGoalSettingAssessmentHeader, error) {
	_, insErr := s.dbConn.Model(&goalAssessmentResp).Insert()

	if insErr != nil {
		s.l.Error("SaveGoalSettingAssessment Error--", insErr)
		return goalAssessmentResp, insErr
	}

	dataBytes, _ := json.Marshal(goalAssessmentResp)
	s.l.Debug("User table json : %q", string(dataBytes))

	return goalAssessmentResp, nil
}

func (s *SaveAssessment) UpdateGoalSettingAssessment(goalAssessmentResp models.AldGoalSettingAssessmentHeader) error {
	_, uptErr := s.dbConn.Model(&goalAssessmentResp).Column("ald_goal_setting_assessment_option_id", "points", "max_points").Where("id=?", goalAssessmentResp.ID).Update()
	if uptErr != nil {
		s.l.Error("UpdateAuditAssessment Error--", uptErr)
		return uptErr
	}

	dataBytes, _ := json.Marshal(goalAssessmentResp)
	s.l.Debug("User table json : %q", string(dataBytes))

	return nil
}

func (s *SaveAssessment) SaveGoalSettingAssessmentLog(goalAssessmentResp models.AldGoalSettingAssessmentLog) error {

	_, insErr := s.dbConn.Model(&goalAssessmentResp).Insert()

	if insErr != nil {
		s.l.Error("SaveGoalSettingAssessmentLog Error--", insErr)
		return insErr
	}

	dataBytes, _ := json.Marshal(goalAssessmentResp)
	s.l.Debug("User table json : %q", string(dataBytes))

	return nil
}
