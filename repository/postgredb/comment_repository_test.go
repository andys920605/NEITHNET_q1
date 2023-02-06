package postgredb_test

import (
	"context"
	"log"
	"regexp"
	"testing"
	"time"

	models_rep "q1/models/repository"
	"q1/repository/postgredb"
	"q1/utils"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-test/deep"
	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xorcare/pointer"
)

func TestMemberRep_Insert(t *testing.T) {
	t.Parallel()
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	db, mock := newMemberMock()
	repository := postgredb.NewCommentRep(db)
	defer func() {
		repository.Close()
	}()
	want := createRandomComment()
	query := `INSERT INTO comments(uuid, parentid, comment, author, favorite, updated_at) VALUES ($1,$2,$3,$4,$5,$6)`
	mock.ExpectExec(regexp.QuoteMeta(query)).
		WithArgs(want.Uuid, want.ParentId, want.Comment, want.Author, want.Favorite, want.UpdateAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := repository.Insert(context.Background(), want)
	assert.NoError(t, err)
}

func TestMemberRep_Find(t *testing.T) {
	t.Parallel()
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	db, mock := newMemberMock()
	repository := postgredb.NewCommentRep(db)
	defer func() {
		repository.Close()
	}()
	want := createRandomComment()
	query := `SELECT uuid, parentid, comment, author, favorite, created_at, updated_at FROM comments WHERE uuid = $1`
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(want.Uuid).
		WillReturnRows(sqlmock.NewRows([]string{"uuid", "parentid", "comment", "author", "favorite", "created_at", "updated_at"}).
			AddRow(want.Uuid, want.ParentId, want.Comment, want.Author, want.Favorite, want.CreateAt, want.UpdateAt))
	got, err := repository.Find(context.Background(), want.Uuid)
	require.NoError(t, err)
	require.Nil(t, deep.Equal(want, got))
}

func TestMemberRep_Updates(t *testing.T) {
	t.Parallel()
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	db, mock := newMemberMock()
	repository := postgredb.NewCommentRep(db)
	defer func() {
		repository.Close()
	}()
	want := createRandomComment()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE comments SET comment =$2, updated_at=$3 WHERE uuid = $1`)).
		WithArgs(want.Uuid, want.Comment, want.UpdateAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := repository.Updates(context.Background(), want)
	require.NoError(t, err)
}

func TestMemberRep_Delete(t *testing.T) {
	t.Parallel()
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	db, mock := newMemberMock()
	repository := postgredb.NewCommentRep(db)
	defer func() {
		repository.Close()
	}()
	want := createRandomComment()
	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM comments WHERE uuid = $1`)).
		WithArgs(want.Uuid).
		WillReturnResult(sqlmock.NewResult(1, 1))
	err := repository.Delete(context.Background(), want.Uuid)
	require.NoError(t, err)
}

// region private methods
func newMemberMock() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New() // mock sql.DB
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, err := gorm.Open("postgres", db) // open gorm db
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb.LogMode(true)
	return gdb, mock
}

func createRandomComment() *models_rep.Comment {
	return &models_rep.Comment{
		Uuid:     uuid.NewV4().String(),
		ParentId: uuid.NewV4().String(),
		Comment:  utils.RandomString(50),
		Author:   utils.RandomString(10),
		Favorite: false,
		CreateAt: pointer.Time(time.Now()),
		UpdateAt: pointer.Time(time.Now()),
	}
}

// endregion
