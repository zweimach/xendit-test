package api

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/zweimach/xendit-test/db"
	"github.com/zweimach/xendit-test/db/mock"
	"github.com/zweimach/xendit-test/utils"
)

func TestListMembers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	login := "xendit"
	id := int32(1)
	name := "budi"

	resBody := fmt.Sprintf(`[{"id":%d,"login":"%s","name":"%s","followers":0,"follows":0,"avatar_url":null}]%s`, id, name, name, "\n")

	e := echo.New()
	e.Validator = utils.NewCustomValidator(validator.New())
	q := mock.NewMockQuerier(ctrl)
	q.EXPECT().ListUsersByOrganizationLogin(gomock.Any(), gomock.Any()).Times(1).Return([]db.User{{
		ID:    id,
		Name:  name,
		Login: name,
	}}, nil)
	h := NewHandler(context.Background(), q)
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	c := e.NewContext(req, rec)
	c.SetPath("/orgs/:login/members")
	c.SetParamNames("login")
	c.SetParamValues(login)

	if assert.NoError(t, h.ListMembers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, resBody, rec.Body.String())
	}
}
