package api

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/zweimach/xendit-test/db"
	"github.com/zweimach/xendit-test/db/mock"
	"github.com/zweimach/xendit-test/utils"
)

func TestCreateComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	login := "xendit"
	id := int32(1)
	text := "Hello, World!"

	reqBody := fmt.Sprintf(`{ "comment": "%s" }`, text)
	resBody := fmt.Sprintf(`{"id":%d,"login":"%s","text":"%s"}%s`, id, login, text, "\n")

	e := echo.New()
	e.Validator = utils.NewCustomValidator(validator.New())
	q := mock.NewMockQuerier(ctrl)
	q.EXPECT().CreateCommentByOrganizationLogin(gomock.Any(), gomock.Any()).Times(1).Return(db.Comment{
		ID:   id,
		Text: text,
	}, nil)
	h := NewHandler(context.Background(), q)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c := e.NewContext(req, rec)
	c.SetPath("/orgs/:login/comments")
	c.SetParamNames("login")
	c.SetParamValues(login)

	if assert.NoError(t, h.CreateComment(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, resBody, rec.Body.String())
	}
}

func TestListComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	login := "xendit"
	id := int32(1)
	text := "Hello, World!"

	resBody := fmt.Sprintf(`[{"id":%d,"login":"%s","text":"%s"}]%s`, id, login, text, "\n")

	e := echo.New()
	e.Validator = utils.NewCustomValidator(validator.New())
	q := mock.NewMockQuerier(ctrl)
	q.EXPECT().ListCommentsByOrganizationLogin(gomock.Any(), gomock.Any()).Times(1).Return([]db.Comment{{
		ID:   id,
		Text: text,
	}}, nil)
	h := NewHandler(context.Background(), q)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	c := e.NewContext(req, rec)
	c.SetPath("/orgs/:login/comments")
	c.SetParamNames("login")
	c.SetParamValues(login)

	if assert.NoError(t, h.ListComments(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, resBody, rec.Body.String())
	}
}

func TestDeleteComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	login := "xendit"

	e := echo.New()
	e.Validator = utils.NewCustomValidator(validator.New())
	q := mock.NewMockQuerier(ctrl)
	q.EXPECT().DeleteCommentsByOrganizationLogin(gomock.Any(), gomock.Any()).Times(1).Return(nil)
	h := NewHandler(context.Background(), q)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	c := e.NewContext(req, rec)
	c.SetPath("/orgs/:login/comments")
	c.SetParamNames("login")
	c.SetParamValues(login)

	if assert.NoError(t, h.DeleteComments(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
