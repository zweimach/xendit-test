package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zweimach/xendit-test/db"
	"github.com/zweimach/xendit-test/utils"
)

type Organization struct {
	ID    int32  `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
}

type ListOrganizationsRequest struct {
	utils.PaginationMeta
}

type ListOrganizationsResponse []Organization

func (h *Handler) ListOrganizations(c echo.Context) error {
	errRes := utils.NewError()

	req := new(ListOrganizationsRequest)
	if err := c.Bind(req); err != nil {
		errRes.AddBadRequest()
		return c.JSON(http.StatusBadRequest, errRes)
	}
	if err := c.Validate(req); err != nil {
		errRes.AddValidationError(err)
		return c.JSON(http.StatusBadRequest, errRes)
	}

	orgs, err := h.db.ListOrganizations(h.ctx, db.ListOrganizationsParams{
		Limit:  utils.PaginationLimit(req.Limit),
		Offset: utils.PaginationOffset(req.Offset),
	})
	if err != nil || orgs == nil {
		errRes.AddNotFound("organization")
		return c.JSON(http.StatusNotFound, errRes)
	}

	var res ListOrganizationsResponse
	for _, v := range orgs {
		res = append(res, Organization{v.ID, v.Login, v.Name})
	}

	return c.JSON(http.StatusOK, res)
}

type GetOrganizationRequest struct {
	Login string `param:"login"`
}

type GetOrganizationResponse Organization

func (h *Handler) GetOrganization(c echo.Context) error {
	errRes := utils.NewError()

	req := new(GetOrganizationRequest)
	if err := c.Bind(req); err != nil {
		errRes.AddBadRequest()
		return c.JSON(http.StatusBadRequest, errRes)
	}
	if err := c.Validate(req); err != nil {
		errRes.AddValidationError(err)
		return c.JSON(http.StatusBadRequest, errRes)
	}

	org, err := h.db.GetOrganizationByLogin(h.ctx, req.Login)
	if err != nil {
		errRes.AddNotFound("organization")
		return c.JSON(http.StatusNotFound, errRes)
	}

	return c.JSON(http.StatusOK, GetOrganizationResponse{org.ID, org.Login, org.Name})
}

type DeleteOrganizationRequest struct {
	Login string `param:"login"`
}

type DeleteOrganizationResponse struct {
	Login string `json:"login"`
}

func (h *Handler) DeleteOrganization(c echo.Context) error {
	errRes := utils.NewError()

	req := new(DeleteOrganizationRequest)
	if err := c.Bind(req); err != nil {
		errRes.AddBadRequest()
		return c.JSON(http.StatusBadRequest, errRes)
	}
	if err := c.Validate(req); err != nil {
		errRes.AddValidationError(err)
		return c.JSON(http.StatusBadRequest, errRes)
	}

	if err := h.db.DeleteOrganizationByLogin(h.ctx, req.Login); err != nil {
		errRes.AddNotFound("organization")
		return c.JSON(http.StatusNotFound, errRes)
	}

	return c.JSON(http.StatusOK, DeleteOrganizationResponse{req.Login})
}
