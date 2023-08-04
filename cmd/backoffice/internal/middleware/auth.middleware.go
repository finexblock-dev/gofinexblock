package middleware

import (
	"errors"
	"fmt"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler/dto"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/presenter"
	"github.com/finexblock-dev/gofinexblock/pkg/admin"
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"log"
	"os"
	"strings"
	"time"
)

func LoginMiddleware(adminService admin.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var body = new(dto.LoginInput)
		var _admin *entity.Admin

		if err = c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AuthErrResponse(errors.Join(errors.New("invalid body"), err)))
		}

		_admin, err = adminService.FindAdminByEmail(body.Email)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(presenter.AuthErrResponse(errors.Join(errors.New("invalid email"), err)))
		}

		err = c.Next()

		statusCode := c.Response().StatusCode()
		if statusCode == fiber.StatusOK {
			_, err = adminService.InsertLoginHistory(_admin.ID)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenter.AuthErrResponse(errors.Join(errors.New("failed to insert login history"), err)))
			}
			_, err = adminService.InsertAccessToken(_admin.ID, time.Now().Add(time.Hour*8))
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenter.AuthErrResponse(errors.Join(errors.New("failed to insert access token"), err)))
			}
		} else {
			_, err = adminService.InsertLoginFailedLog(_admin.ID)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(presenter.AuthErrResponse(errors.Join(errors.New("failed to insert login failed log"), err)))
			}
		}

		return c.Status(fiber.StatusOK).Send(c.Response().Body())
	}
}

func BearerTokenMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") && len(authHeader) > 0 {
			// Add "Bearer " prefix if it's missing
			c.Request().Header.Set("Authorization", fmt.Sprintf("Bearer %v", authHeader))
			log.Println("Authorization", c.Get("Authorization"))
		}
		return JwtMiddleware(os.Getenv("JWT_SECRET"))(c)
	}
}

func JwtMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(presenter.AuthErrResponse(errors.Join(errors.New("invalid token"), err)))
		},
		SigningKey: []byte(secret),
		ContextKey: "admin",
		AuthScheme: "Bearer",
	})
}