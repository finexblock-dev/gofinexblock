package announcement

import (
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/announcement/dto"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type announcementRepository struct {
	db *gorm.DB
}

func (a *announcementRepository) FindAnnouncementByID(tx *gorm.DB, id uint) (result *entity.Announcement, err error) {

	if err = tx.Table(result.TableName()).Where("id = ?", id).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *announcementRepository) FindAllAnnouncement(tx *gorm.DB, limit, offset int) (result []*entity.Announcement, err error) {
	var _announcement *entity.Announcement

	if err = tx.Table(_announcement.TableName()).Preload(clause.Associations).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (a *announcementRepository) SearchAnnouncement(tx *gorm.DB, input dto.SearchAnnouncementInput) (result []*entity.Announcement, err error) {
	var _announcement *entity.Announcement

	query := tx.Table(_announcement.TableName())
	if input.Title != "" {
		query.Where("title = ?", input.Title)
	}

	if input.Word != "" {
		query = query.Where("kor LIKE ? OR eng LIKE ? OR cn LIKE ?", "%"+input.Word+"%", "%"+input.Word+"%", "%"+input.Word+"%")
	}

	query = query.Where("visible = ?", input.Visible)
	query = query.Where("pinned = ?", input.Pinned)

	if input.Limit > 0 {
		query.Limit(input.Limit)
	}

	if input.Offset > 0 {
		query.Offset(input.Offset)
	}

	if err = query.Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *announcementRepository) InsertAnnouncement(tx *gorm.DB, _announcement *entity.Announcement) (result *entity.Announcement, err error) {
	if err = tx.Table(result.TableName()).Create(_announcement).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *announcementRepository) UpdateAnnouncement(tx *gorm.DB, id uint, _announcement *entity.Announcement) (result *entity.Announcement, err error) {
	if err = tx.Table(result.TableName()).Where("id = ?", id).Updates(_announcement).Error; err != nil {
		return nil, err
	}

	result = _announcement
	return result, nil
}

func (a *announcementRepository) DeleteAnnouncement(tx *gorm.DB, id uint) (err error) {

	var _announcement = &entity.Announcement{ID: id}

	if err = tx.Table(_announcement.TableName()).Where("id = ?", id).Delete(_announcement).Error; err != nil {
		return err
	}

	return nil
}

func (a *announcementRepository) InsertCategory(tx *gorm.DB, ko, en, cn string) (result *entity.AnnouncementCategory, err error) {
	result = &entity.AnnouncementCategory{
		KoreanType:  ko,
		EnglishType: en,
		ChineseType: cn,
	}
	if err = tx.Table(result.TableName()).Create(result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *announcementRepository) FindAllCategory(tx *gorm.DB) (result []*entity.AnnouncementCategory, err error) {
	var _category *entity.AnnouncementCategory

	if err = tx.Table(_category.TableName()).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *announcementRepository) UpdateCategory(tx *gorm.DB, id uint, ko, en, cn string) (err error) {
	var _category = &entity.AnnouncementCategory{
		KoreanType:  ko,
		EnglishType: en,
		ChineseType: cn,
	}

	if err = tx.Table(_category.TableName()).Where("id = ?", id).Updates(_category).Error; err != nil {
		return err
	}

	return nil
}

func (a *announcementRepository) DeleteCategory(tx *gorm.DB, id uint) (err error) {
	var _category *entity.AnnouncementCategory

	if err = tx.Table(_category.TableName()).Where("id = ?", id).Delete(&entity.AnnouncementCategory{ID: id}).Error; err != nil {
		return err
	}

	return nil
}

func newAnnouncementRepository(db *gorm.DB) *announcementRepository {
	return &announcementRepository{db: db}
}

func (a *announcementRepository) Conn() *gorm.DB {
	return a.db
}

func (a *announcementRepository) Tx(level sql.IsolationLevel) *gorm.DB {
	return a.db.Begin(&sql.TxOptions{Isolation: level})
}
