package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zweimach/xendit-test/db"
	"github.com/zweimach/xendit-test/utils"
)

type Membership struct {
	UserID         int32 `json:"user_id"`
	OrganizationID int32 `json:"organization_id"`
}

type ListMembersRequest struct {
	utils.PaginationMeta
	Login string `param:"login" validate:"required"`
}

type ListMembersResponse []User

// ListMembers godoc
// @Summary List all members
// @Description List all members
// @ID list-members
// @Tags Organizations
// @Accept  json
// @Produce  json
// @Param login path string true "Login of the organization"
// @Success 200 {object} ListMembersResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /orgs/{login}/members [get]
func (h *Handler) ListMembers(c echo.Context) error {
	errRes := utils.NewError()

	req := new(ListMembersRequest)
	if err := c.Bind(req); err != nil {
		errRes.AddBadRequest()
		return c.JSON(http.StatusBadRequest, errRes)
	}
	if err := c.Validate(req); err != nil {
		errRes.AddValidationError(err)
		return c.JSON(http.StatusBadRequest, errRes)
	}

	members, err := h.db.ListUsersByOrganizationLogin(h.ctx, db.ListUsersByOrganizationLoginParams{
		Limit:  utils.PaginationLimit(req.Limit),
		Offset: utils.PaginationOffset(req.Offset),
		Login:  req.Login,
	})
	if err != nil || members == nil {
		errRes.AddNotFound("organization or member")
		return c.JSON(http.StatusNotFound, errRes)
	}

	var res ListMembersResponse
	for _, v := range members {
		user := User{v.ID, v.Login, v.Name, v.Followers, v.Follows, nil}
		if v.AvatarUrl.Valid {
			user.AvatarUrl = &v.AvatarUrl.String
		}
		res = append(res, user)
	}

	return c.JSON(http.StatusOK, res)
}
