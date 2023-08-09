package handler

import (
	"errors"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler/dto"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/presenter"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/types"
	"github.com/finexblock-dev/gofinexblock/pkg/goredis"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

// XRange @XRange
//
//	@description	XRange command.
//	@tags			Redis
//	@accept			json
//	@produce		json
//	@param			XRangeInput	query		dto.XRangeInput			true	"XRangeInput"
//	@success		200			{object}	interface{}				"Success"
//	@failure		400			{object}	presenter.ErrResponse	"Failed"
//	@router			/redis/xrange [get]
func XRange(service goredis.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var messages []redis.XMessage
		var query = new(dto.XRangeInput)
		var err error

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		messages, err = service.XRange(query.Stream, query.Start, query.End)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, err))
		}

		return c.JSON(messages)
	}
}

// XInfoStream @XInfoStream
//
//	@description	XInfoStream command.
//	@tags			Redis
//	@accept			json
//	@produce		json
//	@param			XInfoStreamInput	query		dto.XInfoStreamInput	true	"XInfoStreamInput"
//	@success		200					{object}	interface{}				"Success"
//	@failure		400					{object}	presenter.ErrResponse	"Failed"
//	@router			/redis/xinfostream [get]
func XInfoStream(service goredis.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var info *redis.XInfoStream
		var query = new(dto.XInfoStreamInput)
		var err error

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		info, err = service.XInfoStream(query.Stream)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, err))
		}
		return c.JSON(info)
	}
}

// Get @Get
//
//	@description	Get command.
//	@tags			Redis
//	@accept			json
//	@produce		json
//	@param			GetInput	query		dto.GetInput			true	"GetInput"
//	@success		200			{object}	interface{}				"Success"
//	@failure		400			{object}	presenter.ErrResponse	"Failed"
//	@router			/redis/get [get]
func Get(service goredis.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query = new(dto.GetInput)
		var value string
		var err error

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		value, err = service.Get(query.Key)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, err))
		}

		return c.JSON(value)
	}
}

// Set @Set
//
//	@description	Set command.
//	@tags			Redis
//	@accept			json
//	@produce		json
//	@param			SetInput	body		dto.SetInput			true	"SetInput"
//	@success		200			{object}	interface{}				"Success"
//	@failure		400			{object}	presenter.ErrResponse	"Failed"
//	@router			/redis/set [post]
func Set(service goredis.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body = new(dto.SetInput)
		var err error

		if err = c.BodyParser(body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		err = service.Set(body.Key, body.Value, 0)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, err))
		}

		return c.SendString("OK")
	}
}

// Del @Del
//
//	@description	Del command.
//	@tags			Redis
//	@accept			json
//	@produce		json
//	@param			DelInput	query		dto.DelInput			true	"DelInput"
//	@success		200			{object}	interface{}				"Success"
//	@failure		400			{object}	presenter.ErrResponse	"Failed"
//	@router			/redis/del [delete]
func Del(service goredis.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query = new(dto.DelInput)
		var err error

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		err = service.Del(query.Key)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, err))
		}

		return c.SendString("OK")
	}
}

// Keys @Keys
// @description	Keys command.
// @tags			Redis
// @accept			json
// @produce		json
// @param			KeysInput	query		dto.KeysInput			true	"KeysInput"
// @success		200			{object}	interface{}				"Success"
// @failure		400			{object}	presenter.ErrResponse	"Failed"
// @router			/redis/keys [get]
func Keys(service goredis.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query = new(dto.KeysInput)
		var keys []string
		var err error

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		keys, err = service.Keys(query.Pattern)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, err))
		}

		return c.JSON(keys)
	}
}

// DeleteAllRefreshTokens @DeleteAllRefreshTokens
// @description	DeleteAllRefreshTokens command.
// @tags			Redis
// @accept			json
// @produce		json
// @success		200			{object}	interface{}				"Success"
// @failure		400			{object}	presenter.ErrResponse	"Failed"
// @router			/redis/deleteRefreshToken [delete]
func DeleteAllRefreshTokens(service goredis.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var keys []string

		keys, err = service.Keys("finexblock:refreshToken:*")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, err))
		}

		for _, key := range keys {
			err = service.Del(key)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, err))
			}
		}

		return c.SendString("OK")

	}
}