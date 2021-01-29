package auditAssessmentService

import (
	"ecargoware/alcochange-dtx/dtos"
	"ecargoware/alcochange-dtx/internals/daos"
	"ecargoware/alcochange-dtx/sentryaccounts"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type AuditAssessment struct {
	dbConn             *pg.DB
	l                  *log.Logger
	auditAssessmentDao daos.AuditAssessmentDao
}

func NewAuditAssessment(l *log.Logger, dbConn *pg.DB) *AuditAssessment {
	return &AuditAssessment{
		l:                  l,
		dbConn:             dbConn,
		auditAssessmentDao: daos.NewAuditAssessmentDB(l, dbConn),
	}
}

// GetAuditAssessmentMessage service for logic
func (a *AuditAssessment) GetAuditAssessmentMessage() (*dtos.AuditAssessmentResponse, error) {
	auditIns := dtos.AuditAssessmentResponse{}
	options := dtos.AuditAssessmentOption{}

	auditQuestionResponse, err := a.auditAssessmentDao.AuditAssessmentQuestion()
	if err != nil {
		a.l.Error("GetAuditAssessmentMessage Error - ", err)
		sentryaccounts.SentryLogExceptions(err)
		return nil, err
	}

	for _, auditQuestion := range auditQuestionResponse {

		auditIns.ID = auditQuestion.ID
		auditIns.Question = auditQuestion.Question
		auditIns.QuestionNo = auditQuestion.QuestionNo
		auditIns.OptionType = auditQuestion.OptionType
		auditIns.OptionTypeLabel = auditQuestion.OptionTypeLabel
		auditIns.SequenceOrder = auditQuestion.SequenceOrder

		auditOptionResponse, err := a.auditAssessmentDao.AuditAssessmentOption(auditIns.ID)
		if err != nil {
			a.l.Error("GetAuditAssessmentMessage Error - ", err)
			sentryaccounts.SentryLogExceptions(err)
			return nil, err
		}

		for _, auditOption := range auditOptionResponse {
			options.ID = auditOption.ID
			options.Name = auditOption.Name
			options.Points = auditOption.Points
			options.QuestionID = auditOption.AldAuditAssessmentQuestionID
			options.SequenceOrder = auditOption.SequenceOrder
			auditIns.Options = append(auditIns.Options, options)
		}
	}

	return &auditIns, nil
}
