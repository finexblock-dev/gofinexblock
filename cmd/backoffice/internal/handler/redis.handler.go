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

type RedisAPI interface {
	XRange() fiber.Handler
	XInfoStream() fiber.Handler
	Get() fiber.Handler
	Set() fiber.Handler
	Del() fiber.Handler
	Keys() fiber.Handler
	DeleteAllRefreshTokens() fiber.Handler
}

type RedisHandler struct {
	redisService goredis.Service
}

func NewRedisHandler(redisService goredis.Service) RedisAPI {
	return &RedisHandler{redisService: redisService}
}

// XRange @XRange
// @description	XRange command.
// @tags			Redis
// @accept			json
// @produce		json
// @param			XRangeInput	query		dto.XRangeInput			true	"XRangeInput"
// @success		200			{object}	interface{}				"Success"
// @failure		400			{object}	presenter.ErrResponse	"Failed"
// @router			/redis/xrange [get]
func (r *RedisHandler) XRange() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var messages []redis.XMessage
		var query = new(dto.XRangeInput)
		var err error

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		messages, err = r.redisService.XRange(query.Stream, query.Start, query.End)
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
func (r *RedisHandler) XInfoStream() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var info *redis.XInfoStream
		var query = new(dto.XInfoStreamInput)
		var err error

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		info, err = r.redisService.XInfoStream(query.Stream)
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
func (r *RedisHandler) Get() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query = new(dto.GetInput)
		var value string
		var err error

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		value, err = r.redisService.Get(query.Key)
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
func (r *RedisHandler) Set() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body = new(dto.SetInput)
		var err error

		if err = c.BodyParser(body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		err = r.redisService.Set(body.Key, body.Value, 0)
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
func (r *RedisHandler) Del() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query = new(dto.DelInput)
		var err error

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		err = r.redisService.Del(query.Key)
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
func (r *RedisHandler) Keys() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query = new(dto.KeysInput)
		var keys []string
		var err error

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		keys, err = r.redisService.Keys(query.Pattern)
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
func (r *RedisHandler) DeleteAllRefreshTokens() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var keys []string

		keys, err = r.redisService.Keys("finexblock:refreshToken:*")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, err))
		}

		for _, key := range keys {
			err = r.redisService.Del(key)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenter.UserErrResponse(fiber.StatusBadRequest, err))
			}
		}

		return c.SendString("OK")

	}
}
