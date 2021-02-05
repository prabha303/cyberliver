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
	SaveAuditAssessment(auditAssessmentResp models.AldAuditAssessmentHeader) (models.AldAuditAssessmentHeader, error)
	UpdateAuditAssessment(auditAssessmentResp models.AldAuditAssessmentHeader) error
	SaveAuditAssessmentLog(auditAssessmentResp models.AldAuditAssessmentLog) error
	IsExistsGoalSettingAssessment(userID int64, questionID int64) models.AldGoalSettingAssessmentHeader
	SaveGoalSettingAssessment(goalAssessmentResp models.AldGoalSettingAssessmentHeader) (models.AldGoalSettingAssessmentHeader, error)
	UpdateGoalSettingAssessment(goalAssessmentResp models.AldGoalSettingAssessmentHeader) error
	SaveGoalSettingAssessmentLog(goalAssessmentResp models.AldGoalSettingAssessmentLog) error
	IsExistsSupportiveContact(userID int64) bool
	SaveSupportiveContact(supportiveContactResp models.AldSupportiveContactHeader) (models.AldSupportiveContactHeader, error)
	DeleteSupportiveContact(userID int64) error
	SaveSupportiveContactLog(supportiveContactResp models.AldSupportiveContactLog) error
	IsExistsDrinkHabitAssessment(userID int64, questionID int64) models.AldDrinkHabitAssessmentHeader
	UpdateDrinkHabitAssessment(drinkAssessmentResp models.AldDrinkHabitAssessmentHeader) error
	SaveDrinkHabitAssessment(drinkAssessmentResp models.AldDrinkHabitAssessmentHeader) (models.AldDrinkHabitAssessmentHeader, error)
	SaveDrinkHabitAssessmentLog(drinkAssessmentResp models.AldDrinkHabitAssessmentLog) error

	IsExistsDrinkProfileAssessment(userID int64, drinkID int64) models.AldDrinkProfileHeader
	UpdateDrinkProfileAssessment(drinkProfileResp models.AldDrinkProfileHeader) error
	SaveDrinkProfileAssessment(drinkProfileResp models.AldDrinkProfileHeader) (models.AldDrinkProfileHeader, error)
	SaveDrinkProfileAssessmentLog(drinkProfileResp models.AldDrinkProfileLog) error
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

func (s *SaveAssessment) IsExistsSupportiveContact(userID int64) bool {
	assessmentHeader := models.AldSupportiveContactHeader{}
	s.dbConn.Model(&assessmentHeader).Where("user_id = ?", userID).Select()
	if assessmentHeader.ID > 0 {
		return true
	}
	return false
}

func (s *SaveAssessment) SaveSupportiveContact(supportiveContactResp models.AldSupportiveContactHeader) (models.AldSupportiveContactHeader, error) {
	_, insErr := s.dbConn.Model(&supportiveContactResp).Insert()

	if insErr != nil {
		s.l.Error("SaveSupportiveContact Error--", insErr)
		return supportiveContactResp, insErr
	}

	dataBytes, _ := json.Marshal(supportiveContactResp)
	s.l.Debug("User table json : %q", string(dataBytes))

	return supportiveContactResp, nil
}

func (s *SaveAssessment) DeleteSupportiveContact(userID int64) error {
	supportiveContactResp := models.AldSupportiveContactHeader{}
	_, uptErr := s.dbConn.Model(&supportiveContactResp).Where("user_id=?", supportiveContactResp.UserID).Delete()
	if uptErr != nil {
		s.l.Error("DeleteSupportiveContact Error--", uptErr)
		return uptErr
	}

	return nil
}

func (s *SaveAssessment) SaveSupportiveContactLog(supportiveContactResp models.AldSupportiveContactLog) error {

	_, insErr := s.dbConn.Model(&supportiveContactResp).Insert()

	if insErr != nil {
		s.l.Error("SaveSupportiveContactLog Error--", insErr)
		return insErr
	}

	dataBytes, _ := json.Marshal(supportiveContactResp)
	s.l.Debug("User table json : %q", string(dataBytes))

	return nil
}

func (s *SaveAssessment) IsExistsDrinkHabitAssessment(userID int64, questionID int64) models.AldDrinkHabitAssessmentHeader {
	assessmentHeader := models.AldDrinkHabitAssessmentHeader{}
	s.dbConn.Model(&assessmentHeader).Where("user_id = ? AND ald_drink_habit_assessment_question_id = ? ", userID, questionID).Select()
	return assessmentHeader
}

func (s *SaveAssessment) UpdateDrinkHabitAssessment(drinkAssessmentResp models.AldDrinkHabitAssessmentHeader) error {
	_, uptErr := s.dbConn.Model(&drinkAssessmentResp).Column("ald_drink_habit_assessment_option_id", "points", "max_points").Where("id=?", drinkAssessmentResp.ID).Update()
	if uptErr != nil {
		s.l.Error("UpdateDrinkHabitAssessment Error--", uptErr)
		return uptErr
	}

	dataBytes, _ := json.Marshal(drinkAssessmentResp)
	s.l.Debug("User table json : %q", string(dataBytes))

	return nil
}

func (s *SaveAssessment) SaveDrinkHabitAssessment(drinkAssessmentResp models.AldDrinkHabitAssessmentHeader) (models.AldDrinkHabitAssessmentHeader, error) {
	_, insErr := s.dbConn.Model(&drinkAssessmentResp).Insert()

	if insErr != nil {
		s.l.Error("SaveDrinkHabitAssessment Error--", insErr)
		return drinkAssessmentResp, insErr
	}

	dataBytes, _ := json.Marshal(drinkAssessmentResp)
	s.l.Debug("User table json : %q", string(dataBytes))

	return drinkAssessmentResp, nil
}

func (s *SaveAssessment) SaveDrinkHabitAssessmentLog(drinkAssessmentResp models.AldDrinkHabitAssessmentLog) error {

	_, insErr := s.dbConn.Model(&drinkAssessmentResp).Insert()

	if insErr != nil {
		s.l.Error("SaveDrinkHabitAssessmentLog Error--", insErr)
		return insErr
	}

	dataBytes, _ := json.Marshal(drinkAssessmentResp)
	s.l.Debug("User table json : %q", string(dataBytes))

	return nil
}

func (s *SaveAssessment) IsExistsDrinkProfileAssessment(userID int64, drinkID int64) models.AldDrinkProfileHeader {
	assessmentHeader := models.AldDrinkProfileHeader{}
	s.dbConn.Model(&assessmentHeader).Where("user_id = ? AND drink_id = ? ", userID, drinkID).Select()
	return assessmentHeader
}

func (s *SaveAssessment) UpdateDrinkProfileAssessment(drinkAssessmentResp models.AldDrinkProfileHeader) error {
	_, uptErr := s.dbConn.Model(&drinkAssessmentResp).Column("ald_drink_habit_assessment_option_id", "points", "max_points").Where("id=?", drinkAssessmentResp.ID).Update()
	if uptErr != nil {
		s.l.Error("UpdateDrinkProfileAssessment Error--", uptErr)
		return uptErr
	}

	dataBytes, _ := json.Marshal(drinkAssessmentResp)
	s.l.Debug("User table json : %q", string(dataBytes))

	return nil
}

func (s *SaveAssessment) SaveDrinkProfileAssessment(drinkAssessmentResp models.AldDrinkProfileHeader) (models.AldDrinkProfileHeader, error) {
	_, insErr := s.dbConn.Model(&drinkAssessmentResp).Insert()

	if insErr != nil {
		s.l.Error("SaveDrinkProfileAssessment Error--", insErr)
		return drinkAssessmentResp, insErr
	}

	dataBytes, _ := json.Marshal(drinkAssessmentResp)
	s.l.Debug("User table json : %q", string(dataBytes))

	return drinkAssessmentResp, nil
}

func (s *SaveAssessment) SaveDrinkProfileAssessmentLog(drinkAssessmentResp models.AldDrinkProfileLog) error {

	_, insErr := s.dbConn.Model(&drinkAssessmentResp).Insert()

	if insErr != nil {
		s.l.Error("SaveDrinkProfileAssessmentLog Error--", insErr)
		return insErr
	}

	dataBytes, _ := json.Marshal(drinkAssessmentResp)
	s.l.Debug("User table json : %q", string(dataBytes))

	return nil
}
