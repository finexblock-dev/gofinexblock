package announcement

import (
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/announcement"
	"gorm.io/gorm"
)

type announcementRepository struct {
	db *gorm.DB
}

func (a *announcementRepository) FindAnnouncementByID(id uint) (*announcement.Announcement, error) {
	//TODO implement me
	panic("implement me")
}

func (a *announcementRepository) ScanAnnouncement(limit, offset int) ([]*announcement.Announcement, error) {
	//TODO implement me
	panic("implement me")
}

func (a *announcementRepository) SearchAnnouncement(title, word string, visible, pinned bool, limit, offset int) ([]*announcement.Announcement, error) {
	//TODO implement me
	panic("implement me")
}

func (a *announcementRepository) CreateAnnouncement(categoryID uint, kot, ent, cnt, kor, eng, cn string, visible, pinned bool) (*announcement.Announcement, error) {
	//TODO implement me
	panic("implement me")
}

func (a *announcementRepository) UpdateAnnouncement(id, categoryID uint, kot, ent, cnt, kor, eng, cn string, visible, pinned bool) (*announcement.Announcement, error) {
	//TODO implement me
	panic("implement me")
}

func (a *announcementRepository) DeleteAnnouncement(id uint) error {
	//TODO implement me
	panic("implement me")
}

func (a *announcementRepository) CreateCategory(ko, en, cn string) (*announcement.AnnouncementCategory, error) {
	//TODO implement me
	panic("implement me")
}

func (a *announcementRepository) FindAllCategory() ([]*announcement.AnnouncementCategory, error) {
	//TODO implement me
	panic("implement me")
}

func (a *announcementRepository) UpdateCategory(id uint, ko, en, cn string) error {
	//TODO implement me
	panic("implement me")
}

func (a *announcementRepository) DeleteCategory(id uint) error {
	//TODO implement me
	panic("implement me")
}

func newAnnouncementRepository(db *gorm.DB) *announcementRepository {
	return &announcementRepository{db: db}
}

func (a *announcementRepository) Tx(level sql.IsolationLevel) *gorm.DB {
	return a.db.Begin(&sql.TxOptions{Isolation: level})
}
