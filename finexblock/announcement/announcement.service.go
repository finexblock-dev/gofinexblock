package announcement

import (
	"context"
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/announcement/dto"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/announcement"
	"gorm.io/gorm"
)

type announcementService struct {
	repo Repository
}

func (a *announcementService) FindAnnouncementByID(id uint) (result *announcement.Announcement, err error) {
	err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindAnnouncementByID(tx, id)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *announcementService) FindAllAnnouncement(limit, offset int) (result []*announcement.Announcement, err error) {
	err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindAllAnnouncement(tx, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *announcementService) SearchAnnouncement(input dto.SearchAnnouncementInput) (result []*announcement.Announcement, err error) {
	err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.SearchAnnouncement(tx, input)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *announcementService) InsertAnnouncement(_announcement *announcement.Announcement) (result *announcement.Announcement, err error) {
	err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.InsertAnnouncement(tx, _announcement)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *announcementService) UpdateAnnouncement(id uint, _announcement *announcement.Announcement) (result *announcement.Announcement, err error) {
	err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.UpdateAnnouncement(tx, id, _announcement)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *announcementService) DeleteAnnouncement(id uint) (err error) {
	return a.Conn().Transaction(func(tx *gorm.DB) error {
		return a.repo.DeleteAnnouncement(tx, id)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func (a *announcementService) InsertCategory(ko, en, cn string) (result *announcement.AnnouncementCategory, err error) {
	err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.InsertCategory(tx, ko, en, cn)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *announcementService) FindAllCategory(limit, offset int) (result []*announcement.AnnouncementCategory, err error) {
	err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindAllCategory(tx)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *announcementService) UpdateCategory(id uint, ko, en, cn string) (err error) {
	return a.Conn().Transaction(func(tx *gorm.DB) error {
		return a.repo.UpdateCategory(tx, id, ko, en, cn)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func (a *announcementService) DeleteCategory(id uint) (err error) {
	return a.Conn().Transaction(func(tx *gorm.DB) error {
		return a.repo.DeleteCategory(tx, id)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func (a *announcementService) Conn() *gorm.DB {
	return a.repo.Conn()
}

func (a *announcementService) Tx(level sql.IsolationLevel) *gorm.DB {
	return a.repo.Tx(level)
}

func (a *announcementService) Ctx() context.Context {
	return context.Background()
}

func (a *announcementService) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}

func newAnnouncementService(announcementRepository Repository) *announcementService {
	return &announcementService{repo: announcementRepository}
}