package handler

import (
	"errors"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler/dto"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/presenter"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/types"
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
	"github.com/finexblock-dev/gofinexblock/pkg/image"
	"github.com/gofiber/fiber/v2"
	"mime/multipart"
)

type ImageAPI interface {
	UploadImage() fiber.Handler
	ListImage() fiber.Handler
}

type ImageHandler struct {
	imageService image.Service
}

func NewImageHandler(imageService image.Service) ImageAPI {
	return &ImageHandler{imageService: imageService}
}

// UploadImage @UploadImage
//
//	@description	Upload image.
//	@security		BearerAuth
//	@tags			Image
//	@accept			multipart/form-data
//	@produce		json
//	@param			file	formData	[]file					true	"Upload file"
//	@success		200		{object}	[]entity.Image			"Success"
//	@failure		400		{object}	presenter.MsgResponse	"Failed"
//	@router			/image [post]
func (i *ImageHandler) UploadImage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var result []*entity.Image
		var err error
		var files = new(multipart.Form)

		files, err = c.MultipartForm()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ImageErrResponse(errors.Join(types.ErrFailedToParseImages, err)))
		}

		if result, err = i.imageService.UploadFile(files); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ImageErrResponse(errors.Join(types.ErrFailedToUploadFiles, err)))
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// ListImage @ListImage
//
//	@description	List all images.
//	@security		BearerAuth
//	@tags			Image
//	@accept			json
//	@produce		json
//	@param			ListImageInput	query		dto.ListImageInput		true	"ListImageInput"
//	@success		200				{object}	[]entity.Image			"Success"
//	@failure		400				{object}	presenter.MsgResponse	"Failed"
//	@router			/image [get]
func (i *ImageHandler) ListImage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var query = new(dto.ListImageInput)
		var result []*entity.Image

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ImageErrResponse(errors.Join(types.ErrFailedToParseQuery, err)))
		}

		if result, err = i.imageService.FindAllImages(query.Limit, query.Offset); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.ImageErrResponse(errors.Join(err)))
		}

		return c.Status(fiber.StatusOK).JSON(result)

	}
}
