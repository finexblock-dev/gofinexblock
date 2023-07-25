package user

import (
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/user/structs"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) FindUserDormantByUserID(tx *gorm.DB, userID uint) (result *entity.UserDormant, err error) {
	if err = tx.Table(result.TableName()).Where("user_id = ?", userID).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userRepository) FindUserMemoByUserID(tx *gorm.DB, userID uint) (result *entity.UserMemo, err error) {
	if err = tx.Table(result.TableName()).Where("user_id = ?", userID).First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userRepository) FindUserEmailSignUpByUserID(tx *gorm.DB, userID uint) (result *entity.UserEmailSignUp, err error) {
	if err = tx.Table(result.TableName()).Where("user_id = ?", userID).First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userRepository) FindUserSingleSignOnInfoByCond(tx *gorm.DB, userID uint, ssoType entity.SSOType) (result *entity.UserSingleSignOnInfo, err error) {
	if err = tx.Table(result.TableName()).Where("user_id = ? and sso_type = ?", userID, ssoType).First(&result, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userRepository) FindUserProfileByUserID(tx *gorm.DB, userID uint) (result *entity.UserProfile, err error) {
	if err = tx.Table(result.TableName()).First(&result, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userRepository) SearchUser(tx *gorm.DB, input *structs.SearchUserInput) (result []*entity.User, err error) {
	var users []*entity.User
	var _user = new(entity.User)

	query := tx.Table(_user.TableName())

	if input.ID != 0 {
		query = query.Where("user.id = ?", input.ID)
	}

	if input.GradeID != 0 {
		query = query.Where("user.grade_id = ?", input.GradeID)
	}

	if input.Description != "" {
		query = query.Joins("JOIN user_memo on user_memo.user_id = user.id").Where("user_memo.description LIKE ?", "%"+input.Description+"%")
	}

	if input.UUID != "" {
		query = query.Where("user.uuid = ?", input.UUID)
	}

	if input.UserType != "" {
		query = query.Where("user.user_type = ?", input.UserType)
	}

	if input.Email != "" {
		query = query.Joins("JOIN user_email_signup on user_email_signup.user_id = user.id").Where("user_email_signup.email = ?", input.Email)
	}

	if input.Nickname != "" {
		query = query.Joins("JOIN user_profile on user_profile.user_id = user.id").
			Where("user_profile.nickname = ?", input.Nickname)
	}

	if input.Fullname != "" {
		query = query.Joins("JOIN user_profile on user_profile.user_id = user.id").
			Where("user_profile.fullname = ?", input.Fullname)
	}
	if input.PhoneNumber != "" {
		query = query.Joins("JOIN user_profile on user_profile.user_id = user.id").
			Where("user_profile.phone_number = ?", input.PhoneNumber)
	}

	if input.IsMetaverseUser {
		query = query.Joins("JOIN user_sso_info on user_sso_info.user_id = user.id").
			Where("user_sso_info.sso_type = ?", entity.Metaverse)
	}

	if input.IsDormant {
		query = query.Joins("JOIN user_dormant on user_dormant.user_id = user.id")
	} else {
		query = query.Joins("LEFT JOIN user_dormant on user_dormant.user_id = user.id").
			Where("user_dormant.id IS NULL")
	}

	if input.IsDropOutUser {
		query = query.Where("user.deleted_at IS NOT NULL")
	} else {
		query = query.Where("user.deleted_at IS NULL")
	}

	if input.IsBlock {
		query = query.Where("user.is_block = ?", input.IsBlock)
	}

	if input.Offset != 0 {
		query = query.Offset(input.Offset)
	}

	if input.Limit != 0 {
		query = query.Limit(input.Limit)
	}

	if err = query.Group("user.id").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) CreateMemo(tx *gorm.DB, id uint, desc string) (err error) {
	var _memo = &entity.UserMemo{UserID: id, Description: desc}

	if err = tx.Table(_memo.TableName()).Create(&_memo).Error; err != nil {
		return err
	}

	return nil
}

func (u *userRepository) FindUserByUUID(tx *gorm.DB, uuid string) (*entity.User, error) {
	var _user = new(entity.User)
	var err error

	if err = tx.Table(_user.TableName()).Where("uuid = ?", uuid).First(&_user).Error; err != nil {
		return nil, err
	}

	return _user, nil
}

func (u *userRepository) FindManyUserByUUID(tx *gorm.DB, uuids []string) ([]*entity.User, error) {
	var users []*entity.User
	var table = new(entity.User)
	var err error

	if err = tx.Table(table.TableName()).Where("uuid IN (?)", uuids).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) FindUserByID(tx *gorm.DB, id uint) (*entity.User, error) {
	var _user = new(entity.User)
	var err error

	if err = tx.Table(_user.TableName()).Where("id = ?", id).First(&_user).Error; err != nil {
		return nil, err
	}

	return _user, nil
}

func (u *userRepository) BlockUser(tx *gorm.DB, id uint) error {
	var table = new(entity.User)
	var err error

	if err = tx.Table(table.TableName()).Where("id = ?", id).Update("is_block", true).Error; err != nil {
		return err
	}

	return nil
}

func (u *userRepository) UnBlockUser(tx *gorm.DB, id uint) error {
	var table = new(entity.User)
	var err error

	if err = tx.Table(table.TableName()).Where("id = ?", id).Update("is_block", false).Error; err != nil {
		return err
	}

	return nil
}

func (u *userRepository) Tx(level sql.IsolationLevel) *gorm.DB {
	return u.db.Begin(&sql.TxOptions{Isolation: level})
}

func (u *userRepository) Conn() *gorm.DB {
	return u.db
}

func newUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}
