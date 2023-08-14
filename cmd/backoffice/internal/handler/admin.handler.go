package handler

import (
	"errors"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler/dto"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/parser"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/presenter"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/types"
	"github.com/finexblock-dev/gofinexblock/pkg/admin"
	"github.com/finexblock-dev/gofinexblock/pkg/admin/structs"
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
	"github.com/finexblock-dev/gofinexblock/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type AdminAPI interface {
	FindAllAdmin() fiber.Handler
	FindAdminByGrade() fiber.Handler
	FindLoginFailedLog() fiber.Handler
	FindLoginHistory() fiber.Handler
	SearchApiLog() fiber.Handler
	SearchGradeUpdateLog() fiber.Handler
	SearchPasswordUpdateLog() fiber.Handler
	SearchDeleteLog() fiber.Handler
	DeleteAdmin() fiber.Handler
	BlockAdmin() fiber.Handler
	UnblockAdmin() fiber.Handler
	UpdatePassword() fiber.Handler
	UpdateEmail() fiber.Handler
	UpdateGrade() fiber.Handler
	FindOnlineAdmin() fiber.Handler
}

type AdminHandler struct {
	adminService admin.Service
}

func NewAdminHandler(adminService admin.Service) AdminAPI {
	return &AdminHandler{adminService: adminService}
}

// FindAllAdmin @FindAllAdmin
// @security		BearerAuth
// @description	Find admin list.
// @tags			Admin
// @accept			json
// @produce		json
// @param			FindAllAdminInput	query		dto.FindAllAdminInput	true	"FindAllAdminInput"
// @success		200					{object}	[]entity.PartialAdmin	"Success"
// @failure		400					{object}	presenter.ErrResponse	"Failed"
// @router			/admin [get]
func (a *AdminHandler) FindAllAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var admins []*entity.Admin
		var result []*entity.PartialAdmin
		var query = new(dto.FindAllAdminInput)

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		admins, err = a.adminService.FindAllAdmin(query.Limit, query.Offset)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToFindAdmin, err)))
		}

		result = parser.AdminToPartial(admins)

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// FindAdminByGrade @FindAdminByGrade
//
//	@security		BearerAuth
//	@description	Find admin list by grade.
//	@tags			Admin
//	@accept			json
//	@produce		json
//	@param			FindAdminByGradeInput	query		dto.FindAdminByGradeInput	true	"FindAdminByGradeInput"
//	@success		200						{object}	[]entity.PartialAdmin		"Success"
//	@failure		400						{object}	presenter.ErrResponse		"Failed"
//	@router			/admin/grade [get]
func (a *AdminHandler) FindAdminByGrade() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var admins []*entity.Admin
		var result []*entity.PartialAdmin
		var query = new(dto.FindAdminByGradeInput)

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		admins, err = a.adminService.FindAdminByGrade(entity.GradeType(query.Grade), query.Limit, query.Offset)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToFindAdmin, err)))
		}

		result = parser.AdminToPartial(admins)
		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// FindLoginFailedLog @FindLoginFailedLog
//
//	@security		BearerAuth
//	@description	Find login failed log of entity.
//	@tags			Admin
//	@accept			json
//	@produce		json
//	@param			FindLoginFailedLogInput	query		dto.FindLoginFailedLogInput		true	"FindLoginFailedLogInput"
//	@success		200						{object}	[]entity.AdminLoginFailedLog	"Success"
//	@failure		400						{object}	presenter.ErrResponse			"Failed"
//	@router			/admin/log/failed [get]
func (a *AdminHandler) FindLoginFailedLog() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var query = new(dto.FindLoginFailedLogInput)
		var result []*entity.AdminLoginFailedLog

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		result, err = a.adminService.FindLoginFailedLogOfAdmin(query.AdminID, query.Limit, query.Offset)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToFindLoginFailedLog, err)))
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// FindLoginHistory @FindLoginHistoryOfAdmin
//
//	@security		BearerAuth
//	@description	Find login history of admin user
//	@tags			Admin
//	@accept			json
//	@produce		json
//	@param			FindLoginHistoryOfAdminInput	query		dto.FindLoginHistoryOfAdminInput	true	"FindLoginHistoryOfAdminInput"
//	@success		200								{object}	[]entity.AdminLoginHistory			"Success"
//	@failure		400								{object}	presenter.ErrResponse				"Failed"
//	@router			/admin/log/login [get]
func (a *AdminHandler) FindLoginHistory() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var query = new(dto.FindLoginHistoryOfAdminInput)
		var result []*entity.AdminLoginHistory

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		result, err = a.adminService.FindLoginHistoryOfAdmin(query.AdminID, query.Limit, query.Offset)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToFindLoginHistory, err)))
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// SearchApiLog @SearchApiLog
//
//	@security		BearerAuth
//	@description	Search api log for matching condition.
//	@tags			Admin
//	@accept			json
//	@produce		json
//	@param			SearchApiLogInput	query		dto.SearchApiLogInput	true	"SearchApiLogInput"
//	@success		200					{object}	[]entity.AdminApiLog	"Success"
//	@failure		400					{object}	presenter.ErrResponse	"Failed"
//	@router			/admin/log/api/search [get]
func (a *AdminHandler) SearchApiLog() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var query = new(dto.SearchApiLogInput)
		var result []*entity.AdminApiLog
		var input *structs.SearchApiLogInput

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		input = &structs.SearchApiLogInput{
			Limit:     query.Limit,
			Offset:    query.Offset,
			AdminID:   query.AdminID,
			StartTime: query.StartTime,
			EndTime:   query.EndTime,
			Method:    entity.ApiMethod(query.Method),
			Endpoint:  query.Endpoint,
		}

		result, err = a.adminService.SearchApiLog(input)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToSearchApiLog, err)))
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// SearchGradeUpdateLog @SearchGradeUpdateLog
//
//	@security		BearerAuth
//	@description	Search grade update log for matching condition.
//	@tags			Admin
//	@accept			json
//	@produce		json
//	@param			SearchGradeUpdateLogInput	query		dto.SearchGradeUpdateLogInput	true	"SearchGradeUpdateLogInput"
//	@success		200							{object}	[]entity.AdminGradeUpdateLog	"Success"
//	@failure		400							{object}	presenter.ErrResponse			"Failed"
//	@router			/admin/log/grade/search [get]
func (a *AdminHandler) SearchGradeUpdateLog() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var query = new(dto.SearchGradeUpdateLogInput)
		var result []*entity.AdminGradeUpdateLog
		var input *structs.SearchGradeUpdateLogInput

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		input = &structs.SearchGradeUpdateLogInput{
			Executor:  query.Executor,
			Target:    query.Target,
			StartTime: query.StartTime,
			EndTime:   query.EndTime,
			Limit:     query.Limit,
			Offset:    query.Offset,
		}

		result, err = a.adminService.SearchGradeUpdateLog(input)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToSearchGradeUpdateLog, err)))
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// SearchPasswordUpdateLog @SearchPasswordUpdateLog
//
//	@security		BearerAuth
//	@description	Search password update log for matching condition.
//	@tags			Admin
//	@accept			json
//	@produce		json
//	@param			SearchPasswordUpdateLogInput	query		dto.SearchPasswordUpdateLogInput	true	"SearchPasswordUpdateLogInput"
//	@success		200								{object}	[]entity.AdminPasswordLog			"Success"
//	@failure		400								{object}	presenter.ErrResponse				"Failed"
//	@router			/admin/log/password/search [get]
func (a *AdminHandler) SearchPasswordUpdateLog() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var query = new(dto.SearchPasswordUpdateLogInput)
		var result []*entity.AdminPasswordLog
		var input = new(structs.SearchPasswordUpdateLogInput)

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		input = &structs.SearchPasswordUpdateLogInput{
			Executor:  query.Executor,
			Target:    query.Target,
			StartTime: query.StartTime,
			EndTime:   query.EndTime,
			Limit:     query.Limit,
			Offset:    query.Offset,
		}

		result, err = a.adminService.SearchPasswordUpdateLog(input)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToSearchPasswordUpdateLog, err)))
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// SearchDeleteLog @SearchDeleteLog
//
//	@security		BearerAuth
//	@description	Search admin delete log for matching condition.
//	@tags			Admin
//	@accept			json
//	@produce		json
//	@param			SearchDeleteLogInput	query		dto.SearchDeleteLogInput	true	"SearchDeleteLogInput"
//	@success		200						{object}	[]entity.AdminDeleteLog		"Success"
//	@failure		400						{object}	presenter.ErrResponse		"Failed"
//	@router			/admin/log/delete/search [get]
func (a *AdminHandler) SearchDeleteLog() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var query = new(dto.SearchDeleteLogInput)
		var result []*entity.AdminDeleteLog
		var input = new(structs.SearchDeleteLogInput)

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		input = &structs.SearchDeleteLogInput{
			Executor:  query.Executor,
			Target:    query.Target,
			StartTime: query.StartTime,
			EndTime:   query.EndTime,
			Limit:     query.Limit,
			Offset:    query.Offset,
		}

		result, err = a.adminService.SearchDeleteLog(input)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToSearchDeleteLog, err)))
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// DeleteAdmin @DeleteAdmin
//
//	@security		BearerAuth
//	@description	Delete admin user.
//	@tags			Admin
//	@accept			json
//	@produce		json
//	@param			DeleteAdminInput	query		dto.DeleteAdminInput	true	"DeleteAdminInput"
//	@success		200					{object}	presenter.MsgResponse	"Success"
//	@failure		400					{object}	presenter.ErrResponse	"Failed"
//	@router			/admin [delete]
func (a *AdminHandler) DeleteAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var query = new(dto.DeleteAdminInput)

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		if err = a.adminService.DeleteAdmin(query.AdminID); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToDeleteAdmin, err)))
		}

		// TODO: Should be implemented, log delete admin

		return c.Status(fiber.StatusOK).JSON(presenter.AdminMsgResponse(fiber.StatusOK, "Successfully deleted"))
	}
}

// BlockAdmin @BlockAdmin
//
//	@security		BearerAuth
//	@description	Block entity.
//	@tags			Admin
//	@accept			json
//	@produce		json
//	@param			BlockAdminInput	body		dto.BlockAdminInput		true	"BlockAdminInput"
//	@success		200				{object}	presenter.MsgResponse	"Success"
//	@failure		400				{object}	presenter.ErrResponse	"Failed"
//	@router			/admin/block [patch]
func (a *AdminHandler) BlockAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var body = new(dto.BlockAdminInput)
		var _admin = new(entity.Admin)

		if err = c.BodyParser(body); err != nil {
			return c.JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseBody, err)))
		}

		_admin, err = a.adminService.FindAdminByID(body.AdminID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToFindAdmin, err)))
		}

		if err = a.adminService.BlockAdmin(_admin.ID); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToBlockAdmin, err)))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.AdminMsgResponse(fiber.StatusOK, "Successfully updated"))
	}
}

// UnblockAdmin @UnblockAdmin
//
//	@security		BearerAuth
//	@description	Unblock entity.
//	@tags			Admin
//	@accept			json
//	@produce		json
//	@param			UnblockAdminInput	body		dto.UnblockAdminInput	true	"UnblockAdminInput"
//	@success		200					{object}	presenter.MsgResponse	"Success"
//	@failure		400					{object}	presenter.ErrResponse	"Failed"
//	@router			/admin/unblock [patch]
func (a *AdminHandler) UnblockAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var body = new(dto.BlockAdminInput)
		var _admin = new(entity.Admin)

		if err = c.BodyParser(body); err != nil {
			return c.JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseBody, err)))
		}

		_admin, err = a.adminService.FindAdminByID(body.AdminID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToFindAdmin, err)))
		}

		if err = a.adminService.UnblockAdmin(_admin.ID); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToUnblockAdmin, err)))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.AdminMsgResponse(fiber.StatusOK, "Successfully updated"))
	}
}

// UpdatePassword @UpdatePassword
//
//	@security		BearerAuth
//	@description	Update password.
//	@tags			Admin
//	@accept			json
//	@produce		json
//	@param			UpdatePasswordInput	body		dto.UpdatePasswordInput	true	"UpdatePasswordInput"
//	@success		200					{object}	presenter.MsgResponse	"Success"
//	@failure		400					{object}	presenter.ErrResponse	"Failed"
//	@router			/admin/password [patch]
func (a *AdminHandler) UpdatePassword() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var body = new(dto.UpdatePasswordInput)

		if err = c.BodyParser(body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseBody, err)))
		}

		if body.NewPassword == body.PrevPassword {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrCheckPassword, err)))
		}

		if !utils.PasswordRegex(body.NewPassword) {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrCheckPassword, err)))
		}

		if err = a.adminService.UpdatePassword(body.AdminID, body.PrevPassword, body.NewPassword); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToUpdatePassword, err)))
		}

		// TODO: Should be implemented, log update password

		return c.Status(fiber.StatusOK).JSON(presenter.AdminMsgResponse(fiber.StatusOK, "Successfully updated"))
	}
}

// UpdateEmail @UpdateEmail
// @security		BearerAuth
// @description	Update email.
// @tags			Admin
// @accept			json
// @produce		json
// @param			UpdateEmailInput	body		dto.UpdateEmailInput	true	"UpdateEmailInput"
// @success		200					{object}	presenter.MsgResponse	"Success"
// @failure		400					{object}	presenter.ErrResponse	"Failed"
// @router			/admin/email [patch]
func (a *AdminHandler) UpdateEmail() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var body = new(dto.UpdateEmailInput)

		if err = c.BodyParser(body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseBody, err)))
		}

		if err = a.adminService.UpdateEmail(body.AdminID, body.Email); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToUpdateEmail, err)))
		}

		// TODO: Should be implemented, log update email
		return c.Status(fiber.StatusOK).JSON(presenter.AdminMsgResponse(fiber.StatusOK, "Successfully updated"))
	}
}

// UpdateGrade @UpdateGrade
//
//	@security		BearerAuth
//	@description	Update grade.
//	@tags			Admin
//	@accept			json
//	@produce		json
//	@param			UpdateGradeInput	body		dto.UpdateGradeInput	true	"UpdateGradeInput"
//	@success		200					{object}	presenter.MsgResponse	"Success"
//	@failure		400					{object}	presenter.ErrResponse	"Failed"
//	@router			/admin/grade [patch]
func (a *AdminHandler) UpdateGrade() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var body = new(dto.UpdateGradeInput)

		if err = c.BodyParser(body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseBody, err)))
		}

		if err = a.adminService.UpdateGrade(body.AdminID, entity.GradeType(body.Grade)); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToUpdateGrade, err)))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.AdminMsgResponse(fiber.StatusOK, "Successfully updated"))
	}
}

// FindOnlineAdmin @FindOnlineAdmin
//
//	@description	Find online admin user for now.
//	@security		BearerAuth
//	@tags			Admin
//	@accept			json
//	@produce		json
//	@param			FindOnlineAdminInput	query		dto.FindOnlineAdminInput	true	"FindOnlineAdminInput"
//	@success		200						{object}	[]entity.Admin				"Success"
//	@failure		400						{object}	presenter.ErrResponse		"Failed"
//	@router			/admin/online [get]
func (a *AdminHandler) FindOnlineAdmin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var query = new(dto.FindOnlineAdminInput)
		var result []*entity.Admin

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		result, err = a.adminService.FindOnlineAdmin(query.Limit, query.Offset)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToFindAdmin, err)))
		}

		return c.Status(fiber.StatusOK).JSON(parser.AdminToPartial(result))
	}
}
