package admin

import (
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"
	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

var _admin = &admin.Admin{}

func (a *AdminRepository) FindAdminByID(tx *gorm.DB, id uint) (*admin.Admin, error) {
	var result *admin.Admin
	var err error

	if err = tx.Table(_admin.TableName()).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (a *AdminRepository) FindAdminByEmail(tx *gorm.DB, email string) (*admin.Admin, error) {
	var result *admin.Admin
	var err error

	if err = tx.Table(_admin.TableName()).Where("email = ?", email).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (a *AdminRepository) FindAdminCredentialsByID(tx *gorm.DB, id uint) (*admin.Admin, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AdminRepository) FindAccessToken(tx *gorm.DB, limit, offset int) ([]*admin.AdminAccessToken, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AdminRepository) FindAdminByGrade(tx *gorm.DB, grade admin.GradeType, limit, offset int) ([]*admin.Admin, error) {
	var result []*admin.Admin
	var err error

	if err = tx.Table(_admin.TableName()).Where("grade = ?", grade).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *AdminRepository) FindAllAdmin(tx *gorm.DB, limit, offset int) ([]*admin.Admin, error) {
	var result []*admin.Admin
	var err error

	if err = tx.Table(_admin.TableName()).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *AdminRepository) CreateAdmin(tx *gorm.DB, email, password string) (*admin.Admin, error) {
	var err error
	var input *admin.Admin

	input = &admin.Admin{
		Email:    email,
		Password: password,
	}

	if err = tx.Table(_admin.TableName()).Create(input).Error; err != nil {
		return nil, err
	}

	return input, nil
}

func (a *AdminRepository) UpdateAdminByID(tx *gorm.DB, id uint, admin *admin.Admin) error {
	var err error

	if err = tx.Table(_admin.TableName()).Where("id = ?", id).Updates(admin).Error; err != nil {
		return err
	}

	return nil
}

func (a *AdminRepository) DeleteAdminByID(tx *gorm.DB, id uint) error {
	var err error

	if err = tx.Table(_admin.TableName()).Where("id = ?", id).Delete(&admin.Admin{}).Error; err != nil {
		return err
	}

	return nil
}

func NewAdminRepository(db *gorm.DB) RepositoryImpl {
	return &AdminRepository{db: db}
}

func (a *AdminRepository) Tx(level sql.IsolationLevel) *gorm.DB {
	return a.db.Begin(&sql.TxOptions{Isolation: level})
}