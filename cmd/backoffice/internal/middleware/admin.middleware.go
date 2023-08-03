package middleware

import (
	"errors"
	"fmt"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler/dto"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/presenter"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/types"
	"github.com/finexblock-dev/gofinexblock/pkg/admin"
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AdminDeleteLogMiddleware(service admin.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var token *jwt.Token
		var claims jwt.MapClaims
		var adminID uint
		var _admin *entity.Admin
		var err error
		var ok bool
		var query = new(dto.DeleteAdminInput)

		if token, ok = c.Locals("admin").(*jwt.Token); !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(presenter.AuthErrResponse(errors.New("token not found")))
		}

		if claims, ok = token.Claims.(jwt.MapClaims); !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(presenter.AuthErrResponse(errors.New("jwt payload not found")))
		}

		adminID = uint(claims["adminId"].(float64))

		_admin, err = service.FindAdminByID(adminID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(err)
		}

		if _admin.Grade != entity.SUPERUSER {
			return c.Status(fiber.StatusUnauthorized).JSON(presenter.AuthErrResponse(errors.New("permission denied")))
		}

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		err = c.Next()

		statusCode := c.Response().StatusCode()
		if statusCode == fiber.StatusOK {
			// FIXME: error handling for failed to insert delete log
			_, _ = service.InsertDeleteLog(_admin.ID, query.AdminID)
		}

		return c.Status(fiber.StatusOK).Send(c.Response().Body())
	}
}

func AdminGradeUpdateLogMiddleware(service admin.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var token *jwt.Token
		var claims jwt.MapClaims
		var adminID uint
		var executor *entity.Admin
		var target *entity.Admin
		var err error
		var ok bool
		var query = new(dto.UpdateGradeInput)

		if token, ok = c.Locals("admin").(*jwt.Token); !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(presenter.AuthErrResponse(errors.New("token not found")))
		}

		if claims, ok = token.Claims.(jwt.MapClaims); !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(presenter.AuthErrResponse(errors.New("jwt payload not found")))
		}

		adminID = uint(claims["adminId"].(float64))

		executor, err = service.FindAdminByID(adminID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(err)
		}

		if err = c.BodyParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		switch {
		case entity.GradeType(query.Grade).Validate() != nil:
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AdminErrResponse(fiber.StatusBadRequest, errors.New("invalid grade")))
		case executor.Grade != entity.SUPERUSER && entity.GradeType(query.Grade) == entity.SUPERUSER:
			return c.Status(fiber.StatusUnauthorized).JSON(presenter.AuthErrResponse(errors.New("permission denied")))
		case executor.Grade != entity.SUPERUSER && executor.Grade != entity.MAINTAINER:
			return c.Status(fiber.StatusUnauthorized).JSON(presenter.AuthErrResponse(errors.New("permission denied")))
		}

		target, err = service.FindAdminByID(query.AdminID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(presenter.AdminErrResponse(fiber.StatusInternalServerError, errors.Join(types.ErrFailedToFindAdmin, err)))
		}

		err = c.Next()

		statusCode := c.Response().StatusCode()
		if statusCode == fiber.StatusOK {
			// FIXME: error handling for failed to insert grade update log
			_, _ = service.InsertGradeUpdateLog(executor.ID, target.ID, target.Grade, entity.GradeType(query.Grade))
		}

		return c.Status(fiber.StatusOK).Send(c.Response().Body())
	}
}

func AdminApiLogMiddleware(service admin.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var token *jwt.Token
		var claims jwt.MapClaims
		var adminID uint
		var _admin *entity.Admin
		var err error
		var ok bool

		if token, ok = c.Locals("admin").(*jwt.Token); !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(presenter.AuthErrResponse(errors.New("token not found")))
		}

		if claims, ok = token.Claims.(jwt.MapClaims); !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(presenter.AuthErrResponse(errors.New("jwt payload not found")))
		}

		adminID = uint(claims["adminId"].(float64))

		_admin, err = service.FindAdminByID(adminID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(err)
		}

		if _, err = service.InsertApiLog(_admin.ID, entity.ApiMethod(c.Method()), c.IP(), fmt.Sprintf("%v%v", c.BaseURL(), c.Path())); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}

		return c.Next()
	}
}

func SuperUserGuard(service admin.Service) fiber.Handler {
	var ret = errors.New("superuser guard")
	return func(c *fiber.Ctx) error {
		var token *jwt.Token
		var claims jwt.MapClaims
		var adminID uint
		var _admin *entity.Admin
		var err error
		var ok bool

		if token, ok = c.Locals("admin").(*jwt.Token); !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(presenter.AuthErrResponse(errors.New("token not found")))
		}

		if claims, ok = token.Claims.(jwt.MapClaims); !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(presenter.AuthErrResponse(errors.New("jwt payload not found")))
		}

		adminID = uint(claims["adminId"].(float64))

		_admin, err = service.FindAdminByID(adminID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(ret)
		}

		if _admin.Grade != entity.SUPERUSER {
			return c.Status(fiber.StatusUnauthorized).JSON(ret)
		}

		return c.Next()
	}
}

func MaintainerGuard(service admin.Service) fiber.Handler {
	var ret = errors.New("maintainer guard")
	return func(c *fiber.Ctx) error {
		var token *jwt.Token
		var claims jwt.MapClaims
		var adminID uint
		var _admin *entity.Admin
		var err error
		var ok bool

		if token, ok = c.Locals("admin").(*jwt.Token); !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(presenter.AuthErrResponse(errors.New("token not found")))
		}

		if claims, ok = token.Claims.(jwt.MapClaims); !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(presenter.AuthErrResponse(errors.New("jwt payload not found")))
		}

		adminID = uint(claims["adminId"].(float64))

		_admin, err = service.FindAdminByID(adminID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(ret)
		}

		if _admin.Grade != entity.MAINTAINER && _admin.Grade != entity.SUPERUSER {
			return c.Status(fiber.StatusUnauthorized).JSON(ret)
		}

		return c.Next()
	}
}

func SupportGuard(service admin.Service) fiber.Handler {
	var ret = errors.New("support guard")
	return func(c *fiber.Ctx) error {
		var token *jwt.Token
		var claims jwt.MapClaims
		var adminID uint
		var err error
		var ok bool

		if token, ok = c.Locals("admin").(*jwt.Token); !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(presenter.AuthErrResponse(errors.New("token not found")))
		}

		if claims, ok = token.Claims.(jwt.MapClaims); !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(presenter.AuthErrResponse(errors.New("jwt payload not found")))
		}

		adminID = uint(claims["adminId"].(float64))

		_, err = service.FindAdminByID(adminID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(ret)
		}

		return c.Next()
	}
}

func InitialLoginGuard(service admin.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Locals("admin").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		adminID := uint(claims["adminId"].(float64))
		_admin, err := service.FindAdminByID(adminID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(presenter.AdminErrResponse(fiber.StatusUnauthorized, errors.Join(errors.New("user not found"), err)))
		}
		if _admin.InitialLogin {
			return c.Status(fiber.StatusUnauthorized).JSON(presenter.AdminErrResponse(fiber.StatusUnauthorized, errors.New("initial login")))
		}

		return c.Next()
	}
}