package announcement

import (
	"context"
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/announcement"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type announcementService struct {
	announcementRepository Repository
}

func (a *announcementService) Conn() *gorm.DB {
	return a.announcementRepository.Conn()
}

func (a *announcementService) Tx(level sql.IsolationLevel) *gorm.DB {
	return a.announcementRepository.Tx(level)
}

func (a *announcementService) FindAnnouncementByID(c *fiber.Ctx, id uint) (*announcement.Announcement, error) {
	//TODO implement me
	panic("implement me")
}

func (a *announcementService) FindAllAnnouncement(c *fiber.Ctx, limit, offset int) ([]*announcement.Announcement, error) {
	//TODO implement me
	panic("implement me")
}

func (a *announcementService) SearchAnnouncement(c *fiber.Ctx, title, word string, visible, pinned bool, limit, offset int) ([]*announcement.Announcement, error) {
	//TODO implement me
	panic("implement me")
}

func (a *announcementService) CreateAnnouncement(c *fiber.Ctx, categoryID uint, kot, ent, cnt, kor, eng, cn string, visible, pinned bool) (*announcement.Announcement, error) {
	//TODO implement me
	panic("implement me")
}

func (a *announcementService) UpdateAnnouncement(c *fiber.Ctx, id, categoryID uint, kot, ent, cnt, kor, eng, cn string, visible, pinned bool) (*announcement.Announcement, error) {
	//TODO implement me
	panic("implement me")
}

func (a *announcementService) DeleteAnnouncement(c *fiber.Ctx, id uint) error {
	//TODO implement me
	panic("implement me")
}

func (a *announcementService) CreateCategory(c *fiber.Ctx, ko, en, cn string) (*announcement.AnnouncementCategory, error) {
	//TODO implement me
	panic("implement me")
}

func (a *announcementService) FindAllCategory(c *fiber.Ctx) ([]*announcement.AnnouncementCategory, error) {
	//TODO implement me
	panic("implement me")
}

func (a *announcementService) UpdateCategory(c *fiber.Ctx, id uint, ko, en, cn string) error {
	//TODO implement me
	panic("implement me")
}

func (a *announcementService) DeleteCategory(c *fiber.Ctx, id uint) error {
	//TODO implement me
	panic("implement me")
}

func newAnnouncementService(announcementRepository Repository) *announcementService {
	return &announcementService{announcementRepository: announcementRepository}
}

func (a *announcementService) Ctx() context.Context {
	return context.Background()
}

func (a *announcementService) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}