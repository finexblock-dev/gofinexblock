package handler

import (
	"errors"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler/dto"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/presenter"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/types"
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
	"github.com/finexblock-dev/gofinexblock/pkg/wallet"
	"github.com/finexblock-dev/gofinexblock/pkg/wallet/structs"
	"github.com/gofiber/fiber/v2"
)

// FindUserAssets @FindAssetByUserID
// @description	Find asset by user id.
// @security		BearerAuth
// @tags			Asset
// @accept			json
// @produce		json
// @param			wallet.FindUserAssetsInput	query		dto.FindUserAssetsInput	true	"FindUserAssetsInput"
// @success		200							{object}	[]structs.Asset			"Success"
// @failure		400							{object}	presenter.MsgResponse	"Failed"
// @router			/asset [get]
func FindUserAssets(service wallet.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query = new(dto.FindUserAssetsInput)
		var assets []*structs.Asset
		var err error

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AssetErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		assets, err = service.FindAllUserAssets(query.UserID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AssetErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToScanWallet, err)))
		}

		return c.Status(fiber.StatusOK).JSON(assets)
	}
}

// FindUserBalanceUpdateLog @FindUserBalanceUpdateLog
// @description	Find user balance update log
// @security		BearerAuth
// @tags			Asset
// @accept			json
// @produce		json
// @param			wallet.FindUserBalanceUpdateLogInput	query		dto.FindUserBalanceUpdateLogInput	true	"FindUserBalanceUpdateLogInput"
// @success		200										{object}	[]entity.CoinTransfer				"Success"
// @failure		400										{object}	presenter.MsgResponse				"Failed"
// @router			/asset/balance/log [get]
func FindUserBalanceUpdateLog(service wallet.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query = new(dto.FindUserBalanceUpdateLogInput)
		var coinTransfers []*entity.CoinTransfer
		var err error

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AssetErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		coinTransfers, err = service.ScanCoinTransferByCond(query.UserID, query.CoinID, query.Limit, query.Offset)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AssetErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToScanWallet, err)))
		}

		return c.Status(fiber.StatusOK).JSON(coinTransfers)
	}
}

// FindUserAssetsByCond @FindUserAssetsByCond
// @description	Find asset by cond.
// @security		BearerAuth
// @tags			Asset
// @accept			json
// @produce		json
// @param			wallet.FindUserAssetsByCondInputInput	query		dto.FindUserAssetsByCondInput	true	"FindUserAssetsByCondInputInput"
// @success		200							{object}	structs.Asset			"Success"
// @failure		400							{object}	presenter.MsgResponse	"Failed"
// @router			/asset/search [get]
func FindUserAssetsByCond(service wallet.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var query = new(dto.FindUserAssetsByCondInput)
		var assets = new(structs.Asset)
		var err error

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AssetErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToParseQuery, err)))
		}

		assets, err = service.FindUserAssetsByCond(query.UserID, query.CoinID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AssetErrResponse(fiber.StatusBadRequest, errors.Join(types.ErrFailedToScanWallet, err)))
		}

		return c.Status(fiber.StatusOK).JSON(assets)
	}
}
