package postgredb

import (
	"context"
	"fmt"
	models_rep "q1/models/repository"
	rep_interface "q1/repository/interface"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
)

type CommentRep struct {
	mutex sync.Mutex
	db    *gorm.DB
}

func NewCommentRep(db *gorm.DB) rep_interface.ICommentRep {
	return &CommentRep{
		db: db,
	}
}

// Close attaches the provider and close the connection
func (rep *CommentRep) Close() {
	rep.db.Close()
}

func (rep *CommentRep) Insert(ctx context.Context, param *models_rep.Comment) error {
	rep.mutex.Lock()
	defer rep.mutex.Unlock()
	db := rep.db.DB()
	query := `INSERT INTO comments(uuid, parentid, comment, author, favorite, updated_at) VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := db.ExecContext(ctx, query, param.Uuid, param.ParentId, param.Comment, param.Author, param.Favorite, time.Now())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	return nil
}
func (rep *CommentRep) Find(ctx context.Context, uuid string) (*models_rep.Comment, error) {
	rep.mutex.Lock()
	defer rep.mutex.Unlock()
	db := rep.db.DB()
	result := &models_rep.Comment{}
	query := `SELECT uuid, parentid, comment, author, favorite, created_at, updated_at FROM comments WHERE uuid = $1`
	row := db.QueryRowContext(ctx, query, uuid)
	if err := row.Scan(&result.Uuid, &result.ParentId, &result.Comment, &result.Author, &result.Favorite, &result.CreateAt, &result.UpdateAt); err != nil {
		return nil, err
	}
	return result, nil
}
func (rep *CommentRep) Updates(ctx context.Context, param *models_rep.Comment) error {
	rep.mutex.Lock()
	defer rep.mutex.Unlock()
	db := rep.db.DB()
	query := `UPDATE comments SET comment =$2, updated_at=$3 WHERE uuid = $1`
	_, err := db.ExecContext(ctx, query, param.Uuid, param.Comment, time.Now())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	return nil
}
func (rep *CommentRep) Delete(ctx context.Context, uuid string) error {
	rep.mutex.Lock()
	defer rep.mutex.Unlock()
	db := rep.db.DB()
	query := `DELETE FROM comments WHERE uuid = $1`
	_, err := db.ExecContext(ctx, query, uuid)
	return err
}
