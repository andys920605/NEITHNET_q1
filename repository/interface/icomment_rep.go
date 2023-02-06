package repository_interface

import (
	"context"
	models_rep "q1/models/repository"
)

//go:generate mockgen -destination=../../test/mock/icomment_mock_repository.go -package=mock q1/repository/interface ICommentRep
type ICommentRep interface {
	Insert(context.Context, *models_rep.Comment) error
	Find(context.Context, string) (*models_rep.Comment, error)
	Updates(context.Context, *models_rep.Comment) error
	Delete(context.Context, string) error
	Close()
}
