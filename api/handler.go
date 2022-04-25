package api

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/zweimach/xendit-test/db"
)

type Handler struct {
	ctx context.Context
	db  db.Querier
}

func NewHandler(ctx context.Context, db db.Querier) *Handler {
	return &Handler{ctx, db}
}

func (h *Handler) Register(e *echo.Echo) {
	orgs := e.Group("/orgs")
	orgs.GET("/:login/comments", h.ListComments)
	orgs.POST("/:login/comments", h.CreateComment)
	orgs.DELETE("/:login/comments", h.DeleteComments)
	orgs.GET("/:login/members", h.ListMembers)
}
