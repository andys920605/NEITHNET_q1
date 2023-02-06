package service

import (
	"context"
	"net/http"
	models_rep "q1/models/repository"
	rep "q1/repository/interface"
	svc_interface "q1/service/interface"
	"q1/utils/errs"
	"time"

	uuid "github.com/satori/go.uuid"
)

var (
	cancelTimeout time.Duration = 3 // default 3 second
)

type CommentSvc struct {
	CommentRep rep.ICommentRep
}

func NewCommentSvc(ICommentRep rep.ICommentRep) svc_interface.ICommentSvc {
	return &CommentSvc{
		CommentRep: ICommentRep,
	}
}

func (svc *CommentSvc) CreateComment(param *models_rep.Comment) (*models_rep.Comment, *errs.ErrorResponse) {
	ctx, cancel := context.WithTimeout(context.Background(), cancelTimeout*time.Second)
	defer cancel()
	param.Uuid = uuid.NewV4().String()
	if errRsp := svc.CommentRep.Insert(ctx, param); errRsp != nil {
		return nil, &errs.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    errRsp.Error(),
		}
	}
	return param, nil
}
func (svc *CommentSvc) GetComment(uuid string) (*models_rep.Comment, *errs.ErrorResponse) {
	ctx, cancel := context.WithTimeout(context.Background(), cancelTimeout*time.Second)
	defer cancel()
	result, errRsp := svc.CommentRep.Find(ctx, uuid)
	if errRsp != nil {
		return nil, &errs.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    errRsp.Error(),
		}
	}
	return result, nil
}
func (svc *CommentSvc) UpdateComment(param *models_rep.Comment) (*models_rep.Comment, *errs.ErrorResponse) {
	ctx, cancel := context.WithTimeout(context.Background(), cancelTimeout*time.Second)
	defer cancel()
	errRsp := svc.CommentRep.Updates(ctx, param)
	if errRsp != nil {
		return nil, &errs.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    errRsp.Error(),
		}
	}
	result, errRsp := svc.CommentRep.Find(ctx, param.Uuid)
	if errRsp != nil {
		return nil, &errs.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    errRsp.Error(),
		}
	}
	return result, nil
}
func (svc *CommentSvc) DeleteComment(uuid string) *errs.ErrorResponse {
	ctx, cancel := context.WithTimeout(context.Background(), cancelTimeout*time.Second)
	defer cancel()
	errRsp := svc.CommentRep.Delete(ctx, uuid)
	if errRsp != nil {
		return &errs.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Message:    errRsp.Error(),
		}
	}
	return nil
}
