package announcement

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/announcement"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository interface {
	types.Repository
	FindAnnouncementByID(id uint) (*announcement.Announcement, error)
	ScanAnnouncement(limit, offset int) ([]*announcement.Announcement, error) // find announcement by limit, offset
	SearchAnnouncement(title, word string, visible, pinned bool, limit, offset int) ([]*announcement.Announcement, error)
	CreateAnnouncement(categoryID uint, kot, ent, cnt, kor, eng, cn string, visible, pinned bool) (*announcement.Announcement, error)     // create announcement
	UpdateAnnouncement(id, categoryID uint, kot, ent, cnt, kor, eng, cn string, visible, pinned bool) (*announcement.Announcement, error) // update announcement
	DeleteAnnouncement(id uint) error

	CreateCategory(ko, en, cn string) (*announcement.AnnouncementCategory, error) // create announcement category
	FindAllCategory() ([]*announcement.AnnouncementCategory, error)
	UpdateCategory(id uint, ko, en, cn string) error
	DeleteCategory(id uint) error
}

type Service interface {
	types.Service
	FindAnnouncementByID(c *fiber.Ctx, id uint) (*announcement.Announcement, error)
	FindAllAnnouncement(c *fiber.Ctx, limit, offset int) ([]*announcement.Announcement, error) // find announcement by limit, offset
	// SearchAnnouncement FIXME: parameter 묶어서 struct로 처리
	SearchAnnouncement(c *fiber.Ctx, title, word string, visible, pinned bool, limit, offset int) ([]*announcement.Announcement, error)
	// CreateAnnouncement FIXME: parameter 묶어서 struct로 처리
	CreateAnnouncement(c *fiber.Ctx, categoryID uint, kot, ent, cnt, kor, eng, cn string, visible, pinned bool) (*announcement.Announcement, error) // create announcement
	// UpdateAnnouncement FIXME: parameter 묶어서 struct로 처리
	UpdateAnnouncement(c *fiber.Ctx, id, categoryID uint, kot, ent, cnt, kor, eng, cn string, visible, pinned bool) (*announcement.Announcement, error) // update announcement
	DeleteAnnouncement(c *fiber.Ctx, id uint) error

	CreateCategory(c *fiber.Ctx, ko, en, cn string) (*announcement.AnnouncementCategory, error) // create announcement category
	FindAllCategory(c *fiber.Ctx) ([]*announcement.AnnouncementCategory, error)
	UpdateCategory(c *fiber.Ctx, id uint, ko, en, cn string) error
	DeleteCategory(c *fiber.Ctx, id uint) error
}

func NewRepository(db *gorm.DB) Repository {
	return newAnnouncementRepository(db)
}

func NewService(repo Repository) Service {
	return newAnnouncementService(repo)
}