package router_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	models_rep "q1/models/repository"
	web "q1/router"
	"q1/test/mock"
	"q1/utils"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
)

var (
	platformPath   string = "/quiz/v1"
	mockCommentSvc *mock.MockICommentSvc
	commentRouter  web.IRouter
	testAll        bool   = false // default is false
	jsonType       string = "application/json"
)

// region setup
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Log("setup test case")
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mockCommentSvc = mock.NewMockICommentSvc(mockCtl)
	commentRouter = web.NewRouter(mockCommentSvc)
	return func(t *testing.T) {
		// tear-down code here
		mockCtl.Finish()
		t.Log("teardown test case")
	}
}

// endregion

func TestAccountRouter_CreateComment(t *testing.T) {
	var model = createRandomComment()
	tests := []struct {
		name   string
		method string
		path   string
		want   *models_rep.Comment
	}{
		{
			name:   "Create Comment",
			method: http.MethodPost,
			path:   platformPath + "/comment",
			want:   model,
		},
	}
	if !testAll {
		teardownTestCase := setupTestCase(t)
		defer teardownTestCase(t)
	}
	for _, test := range tests {
		path := strings.ReplaceAll(test.path, "/", "_")
		t.Run(test.method+path, func(t *testing.T) {
			mockCommentSvc.EXPECT().CreateComment(model).Return(model, nil)
			bodyBuf, _ := json.Marshal(model)
			req, err := http.NewRequest(test.method, test.path, bytes.NewReader(bodyBuf))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Add("Content-Type", jsonType)
			rec := httptest.NewRecorder()
			commentRouter.InitRouter().ServeHTTP(rec, req)
			if status := rec.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
			// Check the response body is what we expect.
			expected, _ := json.Marshal(test.want)
			if rec.Body.String() != string(expected) {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rec.Body.String(), expected)
			}
		})
	}
}

func TestAccountRouter_GetComment(t *testing.T) {
	var model = createRandomComment()
	tests := []struct {
		name   string
		method string
		path   string
		want   *models_rep.Comment
	}{
		{
			name:   "Get Comment",
			method: http.MethodGet,
			path:   platformPath + "/comment/" + model.Uuid,
			want:   model,
		},
	}
	if !testAll {
		teardownTestCase := setupTestCase(t)
		defer teardownTestCase(t)
	}
	for _, test := range tests {
		path := strings.ReplaceAll(test.path, "/", "_")
		t.Run(test.method+path, func(t *testing.T) {
			mockCommentSvc.EXPECT().GetComment(model.Uuid).Return(model, nil)
			req, err := http.NewRequest(test.method, test.path, nil)
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Add("Content-Type", jsonType)
			rec := httptest.NewRecorder()
			commentRouter.InitRouter().ServeHTTP(rec, req)
			if status := rec.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
			// Check the response body is what we expect.
			expected, _ := json.Marshal(test.want)
			if rec.Body.String() != string(expected) {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rec.Body.String(), expected)
			}
		})
	}
}

func TestAccountRouter_UpdateComment(t *testing.T) {
	var model = createRandomComment()
	tests := []struct {
		name   string
		method string
		path   string
		want   *models_rep.Comment
	}{
		{
			name:   "Update Comment",
			method: http.MethodPut,
			path:   platformPath + "/comment/" + model.Uuid,
			want:   model,
		},
	}
	if !testAll {
		teardownTestCase := setupTestCase(t)
		defer teardownTestCase(t)
	}
	for _, test := range tests {
		path := strings.ReplaceAll(test.path, "/", "_")
		t.Run(test.method+path, func(t *testing.T) {
			mockCommentSvc.EXPECT().UpdateComment(model).Return(model, nil)
			bodyBuf, _ := json.Marshal(model)
			req, err := http.NewRequest(test.method, test.path, bytes.NewReader(bodyBuf))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Add("Content-Type", jsonType)
			rec := httptest.NewRecorder()
			commentRouter.InitRouter().ServeHTTP(rec, req)
			if status := rec.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}
			// Check the response body is what we expect.
			expected, _ := json.Marshal(test.want)
			if rec.Body.String() != string(expected) {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rec.Body.String(), expected)
			}
		})
	}
}

func TestAccountRouter_DeleteComment(t *testing.T) {
	var model = createRandomComment()
	tests := []struct {
		name   string
		method string
		path   string
		want   int
	}{
		{
			name:   "Delete Comment",
			method: http.MethodDelete,
			path:   platformPath + "/comment/" + model.Uuid,
			want:   http.StatusOK,
		},
	}
	if !testAll {
		teardownTestCase := setupTestCase(t)
		defer teardownTestCase(t)
	}
	for _, test := range tests {
		path := strings.ReplaceAll(test.path, "/", "_")
		t.Run(test.method+path, func(t *testing.T) {
			mockCommentSvc.EXPECT().DeleteComment(model.Uuid).Return(nil)
			req, err := http.NewRequest(test.method, test.path, nil)
			if err != nil {
				t.Fatal(err)
			}
			rec := httptest.NewRecorder()
			commentRouter.InitRouter().ServeHTTP(rec, req)
			if status := rec.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
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
	}
}
