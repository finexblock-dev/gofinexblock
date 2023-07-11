package user

import (
	"database/sql"
	"errors"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/user"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/finexblock-dev/gofinexblock/finexblock/user/dto"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"math"
)

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) FindUserProfileByUserID(tx *gorm.DB, userID uint) (result *user.UserProfile, err error) {
	if err = tx.Table(result.TableName()).First(&result, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userRepository) FindUserMetadata(tx *gorm.DB, id uint) (result *types.Metadata, err error) {
	var _user *user.User
	var _profile *user.UserProfile
	var metaverseSSO *user.UserSingleSignOnInfo
	var appleSSO *user.UserSingleSignOnInfo
	var googleSSO *user.UserSingleSignOnInfo
	var emailSignUp *user.UserEmailSignUp
	var dormant *user.UserDormant
	var memo *user.UserMemo

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

	if err = tx.Table(metaverseSSO.TableName()).Where(&user.UserSingleSignOnInfo{UserID: id, SSOType: user.Metaverse}).First(&metaverseSSO, "user_id = ?", id).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err = tx.Table(emailSignUp.TableName()).Where(&user.UserEmailSignUp{UserID: id}).First(&emailSignUp).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err = tx.Table(googleSSO.TableName()).Where(&user.UserSingleSignOnInfo{UserID: id, SSOType: user.Google}).First(&googleSSO, "user_id = ?", id).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err = tx.Table(appleSSO.TableName()).Where(&user.UserSingleSignOnInfo{UserID: id, SSOType: user.Apple}).First(&appleSSO, "user_id = ?", id).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err = tx.Table(memo.TableName()).First(&memo, "user_id = ?", id).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err = tx.Table(dormant.TableName()).First(&dormant, "user_id = ?", id).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
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
		IsDormant:         !errors.Is(err, gorm.ErrRecordNotFound),
		IsMetaverseUser:   metaverseSSO.SSOType == "METAVERSE",
		IsAppleUser:       appleSSO.SSOType == "APPLE",
		IsGoogleUser:      googleSSO.SSOType == "GOOGLE",
		IsEmailSignUpUser: metaverseSSO.SSOType != "METAVERSE" && appleSSO.SSOType != "APPLE" && googleSSO.SSOType != "GOOGLE",
		CreatedAt:         _user.CreatedAt,
		UpdatedAt:         _user.UpdatedAt,
		UserMemo:          memo,
	}

	return info, nil
}

func (u *userRepository) SearchUser(tx *gorm.DB, input *dto.SearchUserInput) ([]*types.Metadata, error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) CreateMemo(tx *gorm.DB, id uint, desc string) (err error) {
	//TODO implement me
	panic("implement me")
}

func (u *userRepository) FindUserByUUID(tx *gorm.DB, uuid string) (*user.User, error) {
	var _user *user.User
	var err error

	if err = tx.Table(_user.TableName()).Where("uuid = ?", uuid).First(&_user).Error; err != nil {
		return nil, err
	}

	return _user, nil
}

func (u *userRepository) FindUserByUUIDs(tx *gorm.DB, uuids []string) ([]*user.User, error) {
	var users []*user.User
	var table *user.User
	var err error

	if err = tx.Table(table.TableName()).Where("uuid IN (?)", uuids).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) FindUserByID(tx *gorm.DB, id uint) (*user.User, error) {
	var _user *user.User
	var err error

	if err = tx.Table(_user.TableName()).Where("id = ?", id).First(&_user).Error; err != nil {
		return nil, err
	}

	return _user, nil
}

func (u *userRepository) BlockUser(tx *gorm.DB, id uint) error {
	var table *user.User
	var err error

	if err = tx.Table(table.TableName()).Where("id = ?", id).Update("is_block", true).Error; err != nil {
		return err
	}

	return nil
}

func (u *userRepository) UnBlockUser(tx *gorm.DB, id uint) error {
	var table *user.User
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