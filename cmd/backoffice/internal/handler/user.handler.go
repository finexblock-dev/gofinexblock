package handler

import (
	"errors"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler/dto"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/presenter"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/types"
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
	"github.com/finexblock-dev/gofinexblock/pkg/user"
	"github.com/finexblock-dev/gofinexblock/pkg/user/structs"
	"github.com/gofiber/fiber/v2"
)

// FindUserByID @FindUserByID
//
//	@description	Find user by user id.
//	@security		BearerAuth
//	@tags			User
//	@accept			json
//	@produce		json
//	@param			FindUserByIDInput	query		dto.FindUserByIDInput	true	"FindUserByIDInput"
//	@success		200					{object}	entity.UserMetadata		"Success"
//	@failure		400					{object}	presenter.ErrResponse	"Failed"
//	@router			/user [get]
func FindUserByID(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var query = new(dto.FindUserByIDInput)
		var result = new(entity.UserMetadata)

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		result, err = service.FindUserMetadata(query.UserID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToFindUser, err)))
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// SearchUser @SearchUser
// @description	Search user by given condition.
// @security		BearerAuth
// @tags			User
// @accept			json
// @produce		json
// @param			SearchUserInput	query		dto.SearchUserInput		true	"SearchUserInput"
// @success		200				{object}	[]entity.UserMetadata	"Success"
// @failure		400				{object}	presenter.ErrResponse	"Failed"
// @router			/user/search [get]
func SearchUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var query = new(dto.SearchUserInput)
		var input = new(structs.SearchUserInput)
		var result []*entity.UserMetadata

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		input = &structs.SearchUserInput{
			ID:              query.ID,
			GradeID:         query.GradeID,
			UUID:            query.UUID,
			Email:           query.Email,
			Nickname:        query.Nickname,
			Fullname:        query.Fullname,
			PhoneNumber:     query.PhoneNumber,
			UserType:        query.UserType,
			IsAdult:         query.IsAdult,
			IsBlock:         query.IsBlock,
			IsDormant:       query.IsDormant,
			IsMetaverseUser: query.IsMetaverseUser,
			IsDropOutUser:   query.IsDropOutUser,
			Description:     query.Description,
			Limit:           query.Limit,
			Offset:          query.Offset,
		}

		result, err = service.SearchUser(input)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToSearchUser, err)))
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// BlockUser @BlockUser
// @description	Force block user.
// @security		BearerAuth
// @tags			User
// @accept			json
// @produce		json
// @param			BlockUserInput	body		dto.BlockUserInput		true	"BlockUserInput"
// @success		200				{object}	presenter.MsgResponse	"Success"
// @failure		400				{object}	presenter.ErrResponse	"Failed"
// @router			/user/block [patch]
func BlockUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var body = new(dto.BlockUserInput)
		if err = c.BodyParser(body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseBody, err)))
		}

		err = service.BlockUser(body.UserID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToBlockUser, err)))
		}
		return c.Status(fiber.StatusOK).JSON(presenter.UserMsgResponse(fiber.StatusOK, "Successfully updated"))
	}
}

// UnblockUser @UnblockUser
//
//	@description	Force block user.
//	@security		BearerAuth
//	@tags			User
//	@accept			json
//	@produce		json
//	@param			UnblockUserInput	body		dto.UnblockUserInput	true	"UnblockUserInput"
//	@success		200					{object}	presenter.MsgResponse	"Success"
//	@failure		400					{object}	presenter.ErrResponse	"Failed"
//	@router			/user/unblock [patch]
func UnblockUser(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var body = new(dto.BlockUserInput)
		if err = c.BodyParser(body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseBody, err)))
		}

		err = service.UnBlockUser(body.UserID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToBlockUser, err)))
		}
		return c.Status(fiber.StatusOK).JSON(presenter.UserMsgResponse(fiber.StatusOK, "Successfully updated"))
	}
}

// CreateMemo @CreateMemo
//
//	@description	Force block user.
//	@security		BearerAuth
//	@tags			User
//	@accept			json
//	@produce		json
//	@param			CreateMemoInput	body		dto.CreateMemoInput		true	"CreateMemoInput"
//	@success		200				{object}	presenter.MsgResponse	"Success"
//	@failure		400				{object}	presenter.ErrResponse	"Failed"
//	@router			/user/memo [post]
func CreateMemo(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var body = new(dto.CreateMemoInput)

		if err = c.BodyParser(body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseBody, err)))
		}

		err = service.CreateMemo(body.UserID, body.Description)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToCreateMemo, err)))
		}
		return c.Status(fiber.StatusCreated).JSON(presenter.UserMsgResponse(fiber.StatusOK, "Successfully created memo"))
	}
}