package user

import (
	"database/sql"
	"errors"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/finexblock-dev/gofinexblock/finexblock/user/dto"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"math"
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

func (u *userRepository) FindUserMetadata(tx *gorm.DB, id uint) (result *types.Metadata, err error) {
	var _user *entity.User
	var _profile *entity.UserProfile
	var metaverseSSO *entity.UserSingleSignOnInfo
	var appleSSO *entity.UserSingleSignOnInfo
	var googleSSO *entity.UserSingleSignOnInfo
	var dormant *entity.UserDormant
	var memo *entity.UserMemo

	var btc []decimal.Decimal
	var btcTotal decimal.Decimal

	_user, err = u.FindUserByID(tx, id)
	if err != nil {
		return nil, err
	}

	_profile, err = u.FindUserProfileByUserID(tx, id)
	if err != nil {
		return nil, err
	}

	metaverseSSO, err = u.FindUserSingleSignOnInfoByCond(tx, id, entity.Metaverse)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	appleSSO, err = u.FindUserSingleSignOnInfoByCond(tx, id, entity.Apple)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	googleSSO, err = u.FindUserSingleSignOnInfoByCond(tx, id, entity.Google)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	_, err = u.FindUserEmailSignUpByUserID(tx, id)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	memo, err = u.FindUserMemoByUserID(tx, id)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	dormant, err = u.FindUserDormantByUserID(tx, id)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// Calculate the BTC balance
	if err := tx.Table("coin_transfer").
		Select("amount as btc").
		Joins("JOIN wallet ON coin_transfer.wallet_id = wallet.id").
		Where("wallet.user_id = ? AND wallet.coin_id = 1", id).
		Find(&btc).Error; !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		return nil, err
	}

	btcTotal = decimal.Zero
	DivInto := decimal.NewFromFloat(math.Pow10(8))
	if len(btc) != 0 {
		for _, v := range btc {
			btcTotal.Add(v.Div(DivInto))
		}
	}

	//btc = btc.Div(decimal.NewFromFloat(math.Pow10(8)))

	info := &types.Metadata{
		ID:                _user.ID,
		UUID:              _user.UUID,
		UserType:          _user.UserType,
		Nickname:          _profile.Nickname,
		Fullname:          _profile.Fullname,
		PhoneNumber:       _profile.PhoneNumber,
		BTC:               btcTotal,
		IsBlock:           _user.IsBlock,
		IsDormant:         dormant.ID != 0,
		IsMetaverseUser:   metaverseSSO.SSOType == entity.Metaverse,
		IsAppleUser:       appleSSO.SSOType == entity.Apple,
		IsGoogleUser:      googleSSO.SSOType == entity.Google,
		IsEmailSignUpUser: metaverseSSO.SSOType != entity.Metaverse && appleSSO.SSOType != entity.Apple && googleSSO.SSOType != entity.Google,
		CreatedAt:         _user.CreatedAt,
		UpdatedAt:         _user.UpdatedAt,
		UserMemo:          memo,
	}

	return info, nil
}

func (u *userRepository) SearchUser(tx *gorm.DB, input *dto.SearchUserInput) (result []*types.Metadata, err error) {
	var _user *entity.User
	var users []*entity.User
	var metadata *types.Metadata

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

	for _, v := range users {
		metadata, err = u.FindUserMetadata(tx, v.ID)
		if err != nil {
			return nil, err
		}
		result = append(result, metadata)
	}

	return result, nil
}

func (u *userRepository) CreateMemo(tx *gorm.DB, id uint, desc string) (err error) {
	var _memo = &entity.UserMemo{UserID: id, Description: desc}

	if err = tx.Table(_memo.TableName()).Create(_memo).Error; err != nil {
		return err
	}

	return nil
}

func (u *userRepository) FindUserByUUID(tx *gorm.DB, uuid string) (*entity.User, error) {
	var _user *entity.User
	var err error

	if err = tx.Table(_user.TableName()).Where("uuid = ?", uuid).First(&_user).Error; err != nil {
		return nil, err
	}

	return _user, nil
}

func (u *userRepository) FindUserByUUIDs(tx *gorm.DB, uuids []string) ([]*entity.User, error) {
	var users []*entity.User
	var table *entity.User
	var err error

	if err = tx.Table(table.TableName()).Where("uuid IN (?)", uuids).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) FindUserByID(tx *gorm.DB, id uint) (*entity.User, error) {
	var _user *entity.User
	var err error

	if err = tx.Table(_user.TableName()).Where("id = ?", id).First(&_user).Error; err != nil {
		return nil, err
	}

	return _user, nil
}

func (u *userRepository) BlockUser(tx *gorm.DB, id uint) error {
	var table *entity.User
	var err error

	if err = tx.Table(table.TableName()).Where("id = ?", id).Update("is_block", true).Error; err != nil {
		return err
	}

	return nil
}

func (u *userRepository) UnBlockUser(tx *gorm.DB, id uint) error {
	var table *entity.User
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
