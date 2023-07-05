package admin

import (
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"
	"gorm.io/gorm"
)

var table = &admin.Admin{}

type adminRepository struct {
	db *gorm.DB
}

func (a *adminRepository) Conn() *gorm.DB {
	return a.db
}

func (a *adminRepository) FindAdminByID(tx *gorm.DB, id uint) (*admin.Admin, error) {
	var result *admin.Admin
	var err error

	if err = tx.Table(table.TableName()).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (a *adminRepository) FindAdminByEmail(tx *gorm.DB, email string) (*admin.Admin, error) {
	var result *admin.Admin
	var err error

	if err = tx.Table(table.TableName()).Where("email = ?", email).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (a *adminRepository) FindAdminCredentialsByID(tx *gorm.DB, id uint) (*admin.Admin, error) {
	//TODO implement me
	panic("implement me")
}

func (a *adminRepository) FindAccessToken(tx *gorm.DB, limit, offset int) ([]*admin.AdminAccessToken, error) {
	//TODO implement me
	panic("implement me")
}

func (a *adminRepository) FindAdminByGrade(tx *gorm.DB, grade admin.GradeType, limit, offset int) ([]*admin.Admin, error) {
	var result []*admin.Admin
	var err error

	if err = tx.Table(table.TableName()).Where("grade = ?", grade).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) FindAllAdmin(tx *gorm.DB, limit, offset int) ([]*admin.Admin, error) {
	var result []*admin.Admin
	var err error

	if err = tx.Table(table.TableName()).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminRepository) CreateAdmin(tx *gorm.DB, email, password string) (*admin.Admin, error) {
	var err error
	var input *admin.Admin

	input = &admin.Admin{
		Email:    email,
		Password: password,
	}

	if err = tx.Table(table.TableName()).Create(input).Error; err != nil {
		return nil, err
	}

	return input, nil
}

func (a *adminRepository) UpdateAdminByID(tx *gorm.DB, id uint, admin *admin.Admin) error {
	var err error

	if err = tx.Table(table.TableName()).Where("id = ?", id).Updates(admin).Error; err != nil {
		return err
	}

	return nil
}

func (a *adminRepository) DeleteAdminByID(tx *gorm.DB, id uint) error {
	var err error

	if err = tx.Table(table.TableName()).Where("id = ?", id).Delete(&admin.Admin{}).Error; err != nil {
		return err
	}

	return nil
}

func newAdminRepository(db *gorm.DB) *adminRepository {
	return &adminRepository{db: db}
}

func (a *adminRepository) Tx(level sql.IsolationLevel) *gorm.DB {
	return a.db.Begin(&sql.TxOptions{Isolation: level})
}