package announcement

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/announcement/dto"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/announcement"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"gorm.io/gorm"
)

type Repository interface {
	types.Repository
	FindAnnouncementByID(tx *gorm.DB, id uint) (result *announcement.Announcement, err error)
	FindAllAnnouncement(tx *gorm.DB, limit, offset int) (result []*announcement.Announcement, err error)
	SearchAnnouncement(tx *gorm.DB, input dto.SearchAnnouncementInput) (result []*announcement.Announcement, err error)
	InsertAnnouncement(tx *gorm.DB, _announcement *announcement.Announcement) (result *announcement.Announcement, err error)
	UpdateAnnouncement(tx *gorm.DB, id uint, _announcement *announcement.Announcement) (result *announcement.Announcement, err error)
	DeleteAnnouncement(tx *gorm.DB, id uint) (err error)

	InsertCategory(tx *gorm.DB, ko, en, cn string) (result *announcement.AnnouncementCategory, err error)
	FindAllCategory(tx *gorm.DB) (result []*announcement.AnnouncementCategory, err error)
	UpdateCategory(tx *gorm.DB, id uint, ko, en, cn string) error
	DeleteCategory(tx *gorm.DB, id uint) error
}

type Service interface {
	types.Service
	FindAnnouncementByID(id uint) (result *announcement.Announcement, err error)
	FindAllAnnouncement(limit, offset int) (result []*announcement.Announcement, err error)
	SearchAnnouncement(input dto.SearchAnnouncementInput) (result []*announcement.Announcement, err error)
	InsertAnnouncement(_announcement *announcement.Announcement) (result *announcement.Announcement, err error)
	UpdateAnnouncement(id uint, _announcement *announcement.Announcement) (result *announcement.Announcement, err error)
	DeleteAnnouncement(id uint) (err error)

	InsertCategory(ko, en, cn string) (result *announcement.AnnouncementCategory, err error)
	FindAllCategory(limit, offset int) (result []*announcement.AnnouncementCategory, err error)
	UpdateCategory(id uint, ko, en, cn string) (err error)
	DeleteCategory(id uint) (err error)
}

func NewRepository(db *gorm.DB) Repository {
	return newAnnouncementRepository(db)
}

func NewService(repo Repository) Service {
	return newAnnouncementService(repo)
}