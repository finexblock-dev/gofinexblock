package handler

import (
	"errors"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/types"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler/dto"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/presenter"
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
	"github.com/finexblock-dev/gofinexblock/pkg/wallet"
	"github.com/gofiber/fiber/v2"
)

// ScanWithdrawalRequestByStatus @ScanWithdrawalRequestByStatus
//
//	@description	Find all withdrawal request using limit, offset.
//	@security		BearerAuth
//	@tags			Withdraw
//	@accept			json
//	@produce		json
//	@param			FindAllWithdrawalRequestsInput	query		dto.ScanWithdrawalRequestByStatusInput	true	"FindAllWithdrawalRequestsInput"
//	@success		200								{object}	[]entity.WithdrawalRequest				"Success"
//	@failure		400								{object}	presenter.MsgResponse					"Failed"
//	@router			/withdraw [get]
func ScanWithdrawalRequestByStatus(service wallet.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query = new(dto.ScanWithdrawalRequestByStatusInput)
		var result []*entity.WithdrawalRequest
		var err error

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.WithdrawalErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		var status = entity.WithdrawalStatus(query.Status)

		if err = status.Validate(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.WithdrawalErrResponse(fiber.StatusBadRequest, err))
		}

		result, err = service.ScanWithdrawalRequestByStatusWithLimitOffset(status, query.Limit, query.Offset)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.WithdrawalErrResponse(fiber.StatusBadRequest, err))
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// FindWithdrawalRequestsByUserID @FindWithdrawalRequestsByUserID
//
//	@description	Find withdrawal request by user id using limit, offset.
//	@security		BearerAuth
//	@tags			Withdraw
//	@accept			json
//	@produce		json
//	@param			FindWithdrawalRequestsByUserIDInput	query		dto.FindWithdrawalRequestsByUserIDInput	true	"FindWithdrawalRequestsByUserIDInput"
//	@success		200									{object}	[]entity.WithdrawalRequest				"Success"
//	@failure		400									{object}	presenter.ErrResponse					"Failed"
//	@router			/withdraw/user [get]
func FindWithdrawalRequestsByUserID(service wallet.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var result []*entity.WithdrawalRequest
		var query = new(dto.FindWithdrawalRequestsByUserIDInput)
		var err error

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.WithdrawalErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		result, err = service.ScanWithdrawalRequestByUser(query.UserID, query.Limit, query.Offset)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.WithdrawalErrResponse(fiber.StatusBadRequest, err))
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// RejectWithdrawalRequests @RejectWithdrawalRequests
//
//	@description	Reject withdrawal requests
//	@security		BearerAuth
//	@tags			Withdraw
//	@accept			json
//	@produce		json
//	@param			RejectWithdrawalRequestsInput	query		dto.RejectWithdrawalRequestsInput	true	"RejectWithdrawalRequestsInput"
//	@success		200								{object}	presenter.MsgResponse				"Success"
//	@failure		400								{object}	presenter.ErrResponse				"Failed"
//	@router			/withdraw/reject [patch]
func RejectWithdrawalRequests(service wallet.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query = new(dto.RejectWithdrawalRequestsInput)
		var err error

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.WithdrawalErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		if _, err = service.UpdateWithdrawalRequest(query.ID, entity.REJECTED); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.WithdrawalErrResponse(fiber.StatusBadRequest, err))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.WithdrawalMsgResponse(fiber.StatusOK, "Successfully rejected withdrawal request"))
	}
}

// ApproveWithdrawalRequests @ApproveWithdrawalRequests
//
//	@description	Approve withdrawal requests
//	@security		BearerAuth
//	@tags			Withdraw
//	@accept			json
//	@produce		json
//	@param			ApproveWithdrawalRequestsInput	query		dto.ApproveWithdrawalRequestsInput	true	"ApproveWithdrawalRequestsInput"
//	@success		200								{object}	presenter.MsgResponse				"Success"
//	@failure		400								{object}	presenter.ErrResponse				"Failed"
//	@router			/withdraw/approve [patch]
func ApproveWithdrawalRequests(service wallet.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query = new(dto.ApproveWithdrawalRequestsInput)
		var err error

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.WithdrawalErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		if _, err = service.UpdateWithdrawalRequest(query.ID, entity.APPROVED); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.WithdrawalErrResponse(fiber.StatusBadRequest, err))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.WithdrawalMsgResponse(fiber.StatusOK, "Successfully approved withdrawal request"))
	}
}