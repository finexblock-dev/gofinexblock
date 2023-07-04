package announcement

import (
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/announcement"
	"gorm.io/gorm"
)

type announcementRepository struct {
	db *gorm.DB
}

func (a *announcementRepository) FindAdminByID(tx *gorm.DB, id uint) (*admin.Admin, error) {
	//TODO implement me
	panic("implement me")
}

func (a *announcementRepository) FindAdminByEmail(tx *gorm.DB, email string) (*admin.Admin, error) {
	//TODO implement me
	panic("implement me")
}

func (a *announcementRepository) FindAdminCredentialsByID(tx *gorm.DB, id uint) (*admin.Admin, error) {
	//TODO implement me
	panic("implement me")
}

func (a *announcementRepository) FindAccessToken(tx *gorm.DB, limit, offset int) ([]*admin.AdminAccessToken, error) {
	//TODO implement me
	panic("implement me")
}

func (a *announcementRepository) FindAdminByGrade(tx *gorm.DB, grade admin.GradeType, limit, offset int) ([]*admin.Admin, error) {
	//TODO implement me
	panic("implement me")
}

func (a *announcementRepository) FindAllAdmin(tx *gorm.DB, limit, offset int) ([]*admin.Admin, error) {
	//TODO implement me
	panic("implement me")
}

func (a *announcementRepository) CreateAdmin(tx *gorm.DB, email, password string) (*admin.Admin, error) {
	//TODO implement me
	panic("implement me")
}

func (a *announcementRepository) UpdateAdminByID(tx *gorm.DB, id uint, admin *admin.Admin) error {
	//TODO implement me
	panic("implement me")
}

func (a *announcementRepository) DeleteAdminByID(tx *gorm.DB, id uint) error {
	//TODO implement me
	panic("implement me")
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