package handler

import (
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler/dto"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/presenter"
	"github.com/finexblock-dev/gofinexblock/pkg/announcement"
	"github.com/finexblock-dev/gofinexblock/pkg/announcement/structs"
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2"
)

type AnnouncementAPI interface {
	FindAllAnnouncement() fiber.Handler
	FindAnnouncementByID() fiber.Handler
	SearchAnnouncement() fiber.Handler
	CreateAnnouncement() fiber.Handler
	UpdateAnnouncement() fiber.Handler
	DeleteAnnouncement() fiber.Handler
	FindAllCategory() fiber.Handler
	CreateCategory() fiber.Handler
	UpdateCategory() fiber.Handler
	DeleteCategory() fiber.Handler
}

type AnnouncementHandler struct {
	announcementService announcement.Service
}

func NewAnnouncementHandler(announcementService announcement.Service) AnnouncementAPI {
	return &AnnouncementHandler{announcementService: announcementService}
}

// FindAllAnnouncement @FindAllAnnouncement
//
//	@description	Find all announcement.
//	@security		BearerAuth
//	@tags			Announcement
//	@accept			json
//	@produce		json
//	@param			FindAllAnnouncementInput	query		dto.FindAllAnnouncementInput	true	"FindAllAnnouncementInput"
//	@success		200							{object}	[]entity.Announcement			"Success"
//	@failure		400							{object}	presenter.ErrResponse			"Failed"
//	@router			/announcement/all [get]
func (a *AnnouncementHandler) FindAllAnnouncement() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var result []*entity.Announcement
		var query = new(dto.FindAllAnnouncementInput)

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AnnouncementErrResponse(fiber.StatusBadRequest, err))
		}

		if result, err = a.announcementService.FindAllAnnouncement(query.Limit, query.Offset); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AnnouncementErrResponse(fiber.StatusBadRequest, err))
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// FindAnnouncementByID @FindAnnouncementByID
//
//	@security		BearerAuth
//	@description	Find single announcement by id.
//	@tags			Announcement
//	@accept			json
//	@produce		json
//	@param			FindAnnouncementByIDInput	query		dto.FindAnnouncementByIDInput	true	"FindAnnouncementByIDInput"
//	@success		200							{object}	entity.Announcement				"Success"
//	@failure		400							{object}	presenter.ErrResponse			"Failed"
//	@router			/announcement [get]
func (a *AnnouncementHandler) FindAnnouncementByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var query = new(dto.FindAnnouncementByIDInput)
		var result *entity.Announcement

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AnnouncementErrResponse(fiber.StatusBadRequest, err))
		}

		if result, err = a.announcementService.FindAnnouncementByID(query.ID); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AnnouncementErrResponse(fiber.StatusBadRequest, err))
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// SearchAnnouncement @SearchAnnouncement
//
//	@security		BearerAuth
//	@description	Search announcement.
//	@tags			Announcement
//	@accept			json
//	@produce		json
//	@param			SearchAnnouncementInput	query		dto.SearchAnnouncementInput	true	"SearchAnnouncementInput"
//	@success		200						{object}	[]entity.Announcement		"Success"
//	@failure		400						{object}	presenter.ErrResponse		"Failed"
//	@router			/announcement/search [get]
func (a *AnnouncementHandler) SearchAnnouncement() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var query = new(dto.SearchAnnouncementInput)
		var result []*entity.Announcement

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AnnouncementErrResponse(fiber.StatusBadRequest, err))
		}

		input := &structs.SearchAnnouncementInput{
			Word:       query.Word,
			Title:      query.Title,
			Visible:    query.Visible,
			Pinned:     query.Pinned,
			CategoryID: query.CategoryID,
			Limit:      query.Limit,
			Offset:     query.Offset,
		}

		if result, err = a.announcementService.SearchAnnouncement(input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AnnouncementErrResponse(fiber.StatusBadRequest, err))
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// CreateAnnouncement @CreateAnnouncement
//
//	@security		BearerAuth
//	@description	Insert announcement.
//	@tags			Announcement
//	@accept			json
//	@produce		json
//	@param			CreateAnnouncementInput	body		dto.CreateAnnouncementInput	true	"CreateAnnouncementInput"
//	@success		200						{object}	presenter.MsgResponse		"Success"
//	@failure		400						{object}	presenter.ErrResponse		"Failed"
//	@router			/announcement [post]
func (a *AnnouncementHandler) CreateAnnouncement() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var body = new(dto.CreateAnnouncementInput)
		var _announcement = new(entity.Announcement)

		if err = c.BodyParser(body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AnnouncementErrResponse(fiber.StatusBadRequest, err))
		}

		_announcement = &entity.Announcement{
			CategoryID:   body.CategoryID,
			KoreanTitle:  body.KoreanTitle,
			EnglishTitle: body.EnglishTitle,
			ChineseTitle: body.ChineseTitle,
			Korean:       body.Korean,
			English:      body.English,
			Chinese:      body.Chinese,
			Visible:      body.Visible,
			Pinned:       body.Pinned,
		}

		if _, err = a.announcementService.InsertAnnouncement(_announcement); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AnnouncementErrResponse(fiber.StatusBadRequest, err))
		}

		return c.Status(fiber.StatusCreated).JSON(presenter.AnnouncementMsgResponse(fiber.StatusCreated, "Successfully created"))
	}
}

// UpdateAnnouncement @UpdateAnnouncement
//
//	@security		BearerAuth
//	@description	Update announcement.
//	@tags			Announcement
//	@accept			json
//	@produce		json
//	@param			UpdateAnnouncementInput	body		dto.UpdateAnnouncementInput	true	"UpdateAnnouncementInput"
//	@success		200						{object}	presenter.MsgResponse		"Success"
//	@failure		400						{object}	presenter.ErrResponse		"Failed"
//	@router			/announcement [patch]
func (a *AnnouncementHandler) UpdateAnnouncement() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var body = new(dto.UpdateAnnouncementInput)
		var input = new(entity.Announcement)

		if err = c.BodyParser(body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AnnouncementErrResponse(fiber.StatusBadRequest, err))
		}

		input = &entity.Announcement{
			CategoryID:   body.CategoryID,
			KoreanTitle:  body.KoreanTitle,
			EnglishTitle: body.EnglishTitle,
			ChineseTitle: body.ChineseTitle,
			Korean:       body.Korean,
			English:      body.English,
			Chinese:      body.Chinese,
			Visible:      body.Visible,
			Pinned:       body.Pinned,
		}

		if _, err = a.announcementService.UpdateAnnouncement(body.ID, input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AnnouncementErrResponse(fiber.StatusBadRequest, err))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.AnnouncementMsgResponse(fiber.StatusOK, "Successfully updated"))

	}
}

// DeleteAnnouncement @DeleteAnnouncement
//
//	@security		BearerAuth
//	@description	Delete announcement.
//	@tags			Announcement
//	@accept			json
//	@produce		json
//	@param			DeleteAnnouncementInput	query		dto.DeleteAnnouncementInput	true	"DeleteAnnouncementInput"
//	@success		200						{object}	presenter.MsgResponse		"Success"
//	@failure		400						{object}	presenter.ErrResponse		"Failed"
//	@router			/announcement [delete]
func (a *AnnouncementHandler) DeleteAnnouncement() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var query = new(dto.DeleteAnnouncementInput)

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AnnouncementErrResponse(fiber.StatusBadRequest, err))
		}

		if err = a.announcementService.DeleteAnnouncement(query.AnnouncementID); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AnnouncementErrResponse(fiber.StatusBadRequest, err))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.AnnouncementMsgResponse(fiber.StatusOK, "Successfully deleted"))
	}
}

// FindAllCategory @FindAllCategory
//
//	@security		BearerAuth
//	@description	Find all categories.
//	@tags			Announcement
//	@accept			json
//	@produce		json
//	@param			FindAllCategoryInput	query		dto.FindAllCategoryInput		true	"FindAllCategoryInput"
//	@success		200						{object}	[]entity.AnnouncementCategory	"Success"
//	@failure		400						{object}	presenter.ErrResponse			"Failed"
//	@router			/announcement/category [get]
func (a *AnnouncementHandler) FindAllCategory() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var result []*entity.AnnouncementCategory
		var query = new(dto.FindAllCategoryInput)

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AnnouncementErrResponse(fiber.StatusBadRequest, err))
		}

		if result, err = a.announcementService.FindAllCategory(query.Limit, query.Offset); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AnnouncementErrResponse(fiber.StatusBadRequest, err))
		}

		return c.Status(fiber.StatusOK).JSON(result)
	}
}

// CreateCategory @CreateCategory
//
//	@description	Create announcement category.
//	@security		BearerAuth
//	@tags			Announcement
//	@accept			json
//	@produce		json
//	@param			CreateCategoryInput	body		dto.CreateCategoryInput	true	"CreateCategoryInput"
//	@success		200					{object}	presenter.MsgResponse	"Success"
//	@failure		400					{object}	presenter.ErrResponse	"Failed"
//	@router			/announcement/category [post]
func (a *AnnouncementHandler) CreateCategory() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var body = new(dto.CreateCategoryInput)

		if err = c.BodyParser(body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AnnouncementErrResponse(fiber.StatusBadRequest, err))
		}

		if _, err = a.announcementService.InsertCategory(body.KoreanType, body.EnglishType, body.ChineseType); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AnnouncementErrResponse(fiber.StatusBadRequest, err))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.AnnouncementMsgResponse(fiber.StatusOK, "Successfully created"))

	}
}

// UpdateCategory @UpdateCategory
//
//	@description	Update announcement category.
//	@security		BearerAuth
//	@tags			Announcement
//	@accept			json
//	@produce		json
//	@param			UpdateCategoryInput	body		dto.UpdateCategoryInput	true	"UpdateCategoryInput"
//	@success		200					{object}	presenter.MsgResponse	"Success"
//	@failure		400					{object}	presenter.ErrResponse	"Failed"
//	@router			/announcement/category [patch]
func (a *AnnouncementHandler) UpdateCategory() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var body = new(dto.UpdateCategoryInput)

		if err = c.BodyParser(body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AnnouncementErrResponse(fiber.StatusBadRequest, err))
		}

		if err = a.announcementService.UpdateCategory(body.ID, body.KoreanType, body.EnglishType, body.ChineseType); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AnnouncementErrResponse(fiber.StatusBadRequest, err))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.AnnouncementMsgResponse(fiber.StatusOK, "Successfully updated"))
	}
}

// DeleteCategory @DeleteCategory
//
//	@security		BearerAuth
//	@description	Delete announcement category.
//	@tags			Announcement
//	@accept			json
//	@produce		json
//	@param			DeleteCategoryInput	query		dto.DeleteCategoryInput	true	"DeleteCategoryInput"
//	@success		200					{object}	presenter.MsgResponse	"Success"
//	@failure		400					{object}	presenter.ErrResponse	"Failed"
//	@router			/announcement/category [delete]
func (a *AnnouncementHandler) DeleteCategory() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		var query = new(dto.DeleteCategoryInput)

		if err = c.QueryParser(query); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AnnouncementErrResponse(fiber.StatusBadRequest, err))
		}

		if err = a.announcementService.DeleteCategory(query.ID); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.AnnouncementErrResponse(fiber.StatusBadRequest, err))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.AnnouncementMsgResponse(fiber.StatusOK, "Successfully deleted"))
	}
}
