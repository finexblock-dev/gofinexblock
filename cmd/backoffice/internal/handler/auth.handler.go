package handler

import (
	"errors"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler/dto"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/presenter"
	"github.com/finexblock-dev/gofinexblock/pkg/auth"
	"github.com/gofiber/fiber/v2"
)

type AuthAPI interface {
	Register() fiber.Handler
	Login() fiber.Handler
}

type AuthHandler struct {
	authService auth.Service
}

func NewAuthHandler(authService auth.Service) AuthAPI {
	return &AuthHandler{authService: authService}
}

// Register @Register
//
//	@description	Register new admin user, only superuser can call this api
//	@tags			Auth
//	@accept			json
//	@produce		json
//	@param			RegisterInput	body		dto.AdminRegisterInput	true	"RegisterInput"
//	@success		201				{object}	presenter.ErrResponse	"Success"
//	@failure		400				{object}	presenter.ErrResponse	"Failed"
//	@router			/auth/register [post]
func (a *AuthHandler) Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body = new(dto.AdminRegisterInput)
		var err error

		if err = c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AuthErrResponse(errors.Join(errors.New("invalid request body"), err)))
		}

		if _, err = a.authService.AdminRegister(body.Email, body.Password); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AuthErrResponse(errors.Join(errors.New("failed to register admin"), err)))
		}

		return c.Status(fiber.StatusCreated).JSON(presenter.AuthMsgResponse(fiber.StatusCreated, "Successfully registered"))
	}
}

// Login @Login
//
//	@description	Login API for backoffice admin user
//	@tags			Auth
//	@accept			json
//	@produce		json
//	@param			LoginInput	body		dto.LoginInput			true	"LoginInput"
//	@success		200			{object}	dto.LoginOutput			"Success"
//	@failure		400			{object}	presenter.ErrResponse	"Failed"
//	@router			/auth/login [post]
func (a *AuthHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body = new(dto.LoginInput)
		var err error
		var result string

		if err = c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AuthErrResponse(errors.Join(errors.New("invalid request body"), err)))
		}

		result, err = a.authService.AdminLogin(body.Email, body.Password)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AuthErrResponse(errors.Join(errors.New("invalid email or password"), err)))
		}

		return c.Status(fiber.StatusOK).JSON(&dto.LoginOutput{AccessToken: result})
	}
}
