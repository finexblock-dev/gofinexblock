package user

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/user"
	"gorm.io/gorm"
)

func (u *userService) FindUserByUUID(tx *gorm.DB, uuid string) (*user.User, error) {
	var _user *user.User
	var err error

	if err = tx.Table(_user.TableName()).Where("uuid = ?", uuid).First(&_user).Error; err != nil {
		return nil, err
	}

	return _user, nil
}

func (u *userService) FindUserByUUIDs(tx *gorm.DB, uuids []string) ([]*user.User, error) {
	var users []*user.User
	var table *user.User
	var err error

	if err = tx.Table(table.TableName()).Where("uuid IN (?)", uuids).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userService) FindUserByID(tx *gorm.DB, id uint) (*user.User, error) {
	var _user *user.User
	var err error

	if err = tx.Table(_user.TableName()).Where("id = ?", id).First(&_user).Error; err != nil {
		return nil, err
	}

	return _user, nil
}

func (u *userService) BlockUser(tx *gorm.DB, id uint) error {
	var table *user.User
	var err error

	if err = tx.Table(table.TableName()).Where("id = ?", id).Update("is_block", true).Error; err != nil {
		return err
	}

	return nil
}

func (u *userService) UnBlockUser(tx *gorm.DB, id uint) error {
	var table *user.User
	var err error

	if err = tx.Table(table.TableName()).Where("id = ?", id).Update("is_block", false).Error; err != nil {
		return err
	}

	return nil
}