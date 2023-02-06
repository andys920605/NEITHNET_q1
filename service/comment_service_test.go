package service_test

import (
	config "q1/infras/configs"
	models_rep "q1/models/repository"
	svc "q1/service"
	svc_interface "q1/service/interface"
	"q1/test/mock"
	"q1/utils"
	"q1/utils/errs"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/xorcare/pointer"
)

var (
	cfg *config.Config
)

func init() {
	cfg = &config.Config{
		Server: config.ServerConfig{
			Debug: true,
		},
	}
	utils.ConfigPath = "example"
}

// create mock repositories
// @param mockCtl controller
// @result MockReps model
func getMockReps(mockCtl *gomock.Controller) mockReps {
	return mockReps{
		mockiCommentRep: mock.NewMockICommentRep(mockCtl),
	}
}

// create new Comment service
// @param cfg config
// @param mockReps mock reps
// @result CommentService model
func newCommentService(cfg *config.Config, mockReps mockReps) svc_interface.ICommentSvc {
	// iCommentSvc := svc.NewCommentSvc(mockReps.mockiCommentRep)
	return svc.NewCommentSvc(mockReps.mockiCommentRep)
}

// mock repositories struct
type mockReps struct {
	mockiCommentRep *mock.MockICommentRep
}

func TestCommentService_CreateComment(t *testing.T) {
	t.Parallel()
	model := createRandomComment()
	type args struct {
		arg *models_rep.Comment
	}
	tests := []struct {
		name    string
		prepare func(f *mockReps)
		args    args
		want    *models_rep.Comment
	}{
		{
			name: "Create Comment",
			prepare: func(f *mockReps) {
				gomock.InOrder(
					f.mockiCommentRep.EXPECT().Insert(gomock.Any(), model).Return(nil),
				)
			},
			args: args{arg: model},
			want: model,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			f := getMockReps(mockCtl)
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			commentSvc := newCommentService(cfg, f)
			if got, _ := commentSvc.CreateComment(tt.args.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommentService.CreateComment() = %v, want = %v", got, tt.want)
			}
		})
	}
}
func TestCommentService_GetComment(t *testing.T) {
	t.Parallel()
	model := createRandomComment()
	type args struct {
		arg string
	}
	tests := []struct {
		name    string
		prepare func(f *mockReps)
		args    args
		want    *models_rep.Comment
	}{
		{
			name: "Get Comment",
			prepare: func(f *mockReps) {
				gomock.InOrder(
					f.mockiCommentRep.EXPECT().Find(gomock.Any(), model.Uuid).Return(model, nil),
				)
			},
			args: args{arg: model.Uuid},
			want: model,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			f := getMockReps(mockCtl)
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			commentSvc := newCommentService(cfg, f)
			if got, _ := commentSvc.GetComment(tt.args.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommentService.GetComment() = %v, want = %v", got, tt.want)
			}
		})
	}
}
func TestCommentService_UpdateComment(t *testing.T) {
	t.Parallel()
	model := createRandomComment()
	type args struct {
		arg *models_rep.Comment
	}
	tests := []struct {
		name    string
		prepare func(f *mockReps)
		args    args
		want    *models_rep.Comment
	}{
		{
			name: "Update Comment",
			prepare: func(f *mockReps) {
				gomock.InOrder(
					f.mockiCommentRep.EXPECT().Updates(gomock.Any(), model).Return(nil),
					f.mockiCommentRep.EXPECT().Find(gomock.Any(), model.Uuid).Return(model, nil),
				)
			},
			args: args{arg: model},
			want: model,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			f := getMockReps(mockCtl)
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			commentSvc := newCommentService(cfg, f)
			if got, _ := commentSvc.UpdateComment(tt.args.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommentService.UpdateComment() = %v, want = %v", got, tt.want)
			}
		})
	}
}
func TestCommentService_DeleteComment(t *testing.T) {
	t.Parallel()
	model := createRandomComment()
	type args struct {
		arg string
	}
	tests := []struct {
		name    string
		prepare func(f *mockReps)
		args    args
		want    *errs.ErrorResponse
	}{
		{
			name: "Delete Comment",
			prepare: func(f *mockReps) {
				gomock.InOrder(
					f.mockiCommentRep.EXPECT().Delete(gomock.Any(), model.Uuid).Return(nil),
				)
			},
			args: args{arg: model.Uuid},
			want: nil,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()
			f := getMockReps(mockCtl)
			if tt.prepare != nil {
				tt.prepare(&f)
			}
			commentSvc := newCommentService(cfg, f)
			got := commentSvc.DeleteComment(tt.args.arg)
			if !assert.Equal(t, got, tt.want) {
				t.Errorf("CommentService.DeleteComment() = %v, want = %v", got, tt.want)
			}
		})
	}
}

// create Random omment
func createRandomComment() *models_rep.Comment {
	return &models_rep.Comment{
		Uuid:     uuid.NewV4().String(),
		ParentId: uuid.NewV4().String(),
		Comment:  utils.RandomString(50),
		Author:   utils.RandomString(10),
		Favorite: false,
		CreateAt: utils.ToLocalDate(pointer.Time(time.Now())),
		UpdateAt: utils.ToLocalDate(pointer.Time(time.Now())),
	}
}
