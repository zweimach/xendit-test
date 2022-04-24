package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zweimach/xendit-test/db"
	"github.com/zweimach/xendit-test/utils"
)

type Comment struct {
	ID    int32  `json:"id"`
	Login string `json:"login"`
	Text  string `json:"text"`
}

type CreateCommentRequest struct {
	Comment string `json:"comment" validate:"required"`
	Login   string `json:"-" param:"login" validate:"required"`
}

type CreateCommentResponse Comment

// CreateComment godoc
// @Summary Create a comment
// @Description Create a comment
// @ID create-comment
// @Tags Organizations
// @Accept  json
// @Produce  json
// @Param login path string true "Login of the organization"
// @Param comment body CreateCommentRequest true "Comment to create"
// @Success 201 {object} CreateCommentResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /orgs/{login}/comments [post]
func (h *Handler) CreateComment(c echo.Context) error {
	errRes := utils.NewError()

	req := new(CreateCommentRequest)
	if err := c.Bind(req); err != nil {
		errRes.AddBadRequest()
		return c.JSON(http.StatusBadRequest, errRes)
	}
	if err := c.Validate(req); err != nil {
		errRes.AddValidationError(err)
		return c.JSON(http.StatusBadRequest, errRes)
	}

	comment, err := h.db.CreateCommentByOrganizationLogin(h.ctx, db.CreateCommentByOrganizationLoginParams{
		Text:  req.Comment,
		Login: req.Login,
	})
	if err != nil {
		errRes.AddNotFound("organization")
		return c.JSON(http.StatusNotFound, errRes)
	}

	return c.JSON(http.StatusCreated, CreateCommentResponse{comment.ID, req.Login, comment.Text})
}

type ListCommentsRequest struct {
	utils.PaginationMeta
	Login string `param:"login" validate:"required"`
}

type ListCommentsResponse []Comment

// ListComments godoc
// @Summary List all comments
// @Description List all comments
// @ID list-comments
// @Tags Organizations
// @Accept  json
// @Produce  json
// @Param login path string true "Login of the organization"
// @Success 200 {object} ListCommentsResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /orgs/{login}/comments [get]
func (h *Handler) ListComments(c echo.Context) error {
	errRes := utils.NewError()

	req := new(ListCommentsRequest)
	if err := c.Bind(req); err != nil {
		errRes.AddBadRequest()
		return c.JSON(http.StatusBadRequest, errRes)
	}
	if err := c.Validate(req); err != nil {
		errRes.AddValidationError(err)
		return c.JSON(http.StatusBadRequest, errRes)
	}

	comments, err := h.db.ListCommentsByOrganizationLogin(h.ctx, db.ListCommentsByOrganizationLoginParams{
		Limit:  utils.PaginationLimit(req.Limit),
		Offset: utils.PaginationOffset(req.Offset),
		Login:  req.Login,
	})
	if err != nil || comments == nil {
		errRes.AddNotFound("organization or comment")
		return c.JSON(http.StatusNotFound, errRes)
	}

	var res ListCommentsResponse
	for _, v := range comments {
		res = append(res, Comment{v.ID, req.Login, v.Text})
	}

	return c.JSON(http.StatusOK, res)
}

type DeleteCommentsRequest struct {
	Login string `param:"login" validate:"required"`
}

// DeleteComments godoc
// @Summary Delete all comments
// @Description Delete all comments
// @ID delete-comments
// @Tags Organizations
// @Accept  json
// @Produce  json
// @Param login path string true "Login of the organization"
// @Success 200 "OK"
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /orgs/{login}/comments [delete]
func (h *Handler) DeleteComments(c echo.Context) error {
	errRes := utils.NewError()

	req := new(DeleteCommentsRequest)
	if err := c.Bind(req); err != nil {
		errRes.AddBadRequest()
		return c.JSON(http.StatusBadRequest, errRes)
	}
	if err := c.Validate(req); err != nil {
		errRes.AddValidationError(err)
		return c.JSON(http.StatusBadRequest, errRes)
	}

	if err := h.db.DeleteCommentsByOrganizationLogin(h.ctx, req.Login); err != nil {
		errRes.AddNotFound("organization or comment")
		return c.JSON(http.StatusNotFound, errRes)
	}

	return c.NoContent(http.StatusOK)
}
