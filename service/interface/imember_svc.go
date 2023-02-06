package service_interface

import (
	models_rep "q1/models/repository"
	"q1/utils/errs"
)

//go:generate mockgen -destination=../../test/mock/icomment_mock_service.go -package=mock q1/service/interface ICommentSvc
type ICommentSvc interface {
	CreateComment(*models_rep.Comment) (*models_rep.Comment, *errs.ErrorResponse)
	GetComment(string) (*models_rep.Comment, *errs.ErrorResponse)
	UpdateComment(*models_rep.Comment) (*models_rep.Comment, *errs.ErrorResponse)
	DeleteComment(string) *errs.ErrorResponse
}
