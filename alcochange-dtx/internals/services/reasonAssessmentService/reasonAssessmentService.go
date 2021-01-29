package reasonAssessmentService

import (
	"ecargoware/alcochange-dtx/dtos"
	"ecargoware/alcochange-dtx/internals/daos"
	"ecargoware/alcochange-dtx/sentryaccounts"

	"github.com/FenixAra/go-util/log"
	"github.com/go-pg/pg"
)

type ReasonAssessment struct {
	dbConn              *pg.DB
	l                   *log.Logger
	reasonAssessmentDao daos.ReasonAssessmentDao
}

func NewReasonAssessment(l *log.Logger, dbConn *pg.DB) *ReasonAssessment {
	return &ReasonAssessment{
		l:                   l,
		dbConn:              dbConn,
		reasonAssessmentDao: daos.NewReasonAssessmentDB(l, dbConn),
	}
}

// GetReasonAssessmentMessage service for logic
func (r *ReasonAssessment) GetReasonAssessmentMessage() (*dtos.ReasonAssessmentResponse, error) {
	reasonIns := dtos.ReasonAssessmentResponse{}
	options := dtos.ReasonAssessmentOption{}

	reasonQuestionResponse, err := r.reasonAssessmentDao.ReasonAssessmentQuestion()
	if err != nil {
		r.l.Error("GetReasonAssessmentMessage Error - ", err)
		sentryaccounts.SentryLogExceptions(err)
		return nil, err
	}

	reasonIns.ID = reasonQuestionResponse.ID
	reasonIns.Question = reasonQuestionResponse.Question
	reasonIns.QuestionNo = reasonQuestionResponse.QuestionNo
	reasonIns.OptionType = reasonQuestionResponse.OptionType
	reasonIns.OptionTypeLabel = reasonQuestionResponse.OptionTypeLabel
	reasonIns.SequenceOrder = reasonQuestionResponse.SequenceOrder
	reasonIns.HeaderNote = reasonQuestionResponse.HeaderNote

	reasonOptionResponse, err := r.reasonAssessmentDao.ReasonAssessmentOption()
	if err != nil {
		r.l.Error("GetReasonAssessmentMessage Error - ", err)
		sentryaccounts.SentryLogExceptions(err)
		return nil, err
	}

	for _, reasonOption := range reasonOptionResponse {
		options.ID = reasonOption.ID
		options.Name = reasonOption.Name
		options.Points = reasonOption.Points
		options.QuestionID = reasonOption.AldReasonAssessmentQuestionID
		options.SequenceOrder = reasonOption.SequenceOrder
		reasonIns.Options = append(reasonIns.Options, options)
	}

	return &reasonIns, nil
}
