package user

import (
	"context"
	"database/sql"
	"errors"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/trade"
	"github.com/finexblock-dev/gofinexblock/finexblock/user/structs"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type userService struct {
	repo    Repository
	manager trade.Manager
}

func newUserService(db *gorm.DB, cluster *redis.ClusterClient) *userService {
	return &userService{repo: NewRepository(db), manager: trade.New(cluster)}
}

func (u *userService) Ctx() context.Context {
	return context.Background()
}

func (u *userService) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}

func (u *userService) Tx(level sql.IsolationLevel) *gorm.DB {
	return u.repo.Tx(level)
}

func (u *userService) Conn() *gorm.DB {
	return u.repo.Conn()
}

func (u *userService) FindUserByUUID(uuid string) (result *entity.User, err error) {
	if err = u.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = u.repo.FindUserByUUID(tx, uuid)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userService) FindUserByUUIDs(uuids []string) (result []*entity.User, err error) {
	if err = u.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = u.repo.FindManyUserByUUID(tx, uuids)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userService) FindUserByID(id uint) (result *entity.User, err error) {
	if err = u.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = u.repo.FindUserByID(tx, id)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userService) FindUserMetadata(id uint) (result *entity.UserMetadata, err error) {
	if err = u.Conn().Transaction(func(tx *gorm.DB) error {
		var _user = new(entity.User)
		var _profile = new(entity.UserProfile)
		var metaverseSSO = new(entity.UserSingleSignOnInfo)
		var appleSSO = new(entity.UserSingleSignOnInfo)
		var googleSSO = new(entity.UserSingleSignOnInfo)
		var dormant = new(entity.UserDormant)
		var memo = new(entity.UserMemo)
		var btcTotal = decimal.Zero

		_user, err = u.repo.FindUserByID(tx, id)
		if err != nil {
			return err
		}

		_profile, err = u.repo.FindUserProfileByUserID(tx, id)
		if err != nil {
			return err
		}

		metaverseSSO, err = u.repo.FindUserSingleSignOnInfoByCond(tx, id, entity.Metaverse)
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		appleSSO, err = u.repo.FindUserSingleSignOnInfoByCond(tx, id, entity.Apple)
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		googleSSO, err = u.repo.FindUserSingleSignOnInfoByCond(tx, id, entity.Google)
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		_, err = u.repo.FindUserEmailSignUpByUserID(tx, id)
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		memo, err = u.repo.FindUserMemoByUserID(tx, id)
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		dormant, err = u.repo.FindUserDormantByUserID(tx, id)
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		// Calculate the BTC balance
		btcTotal, err = u.manager.GetBalance(_user.UUID, grpc_order.Currency_BTC.String())
		if err != nil && !errors.Is(err, redis.Nil) {
			return err
		}

		result = &entity.UserMetadata{
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

		return nil
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userService) SearchUser(input *structs.SearchUserInput) (result []*entity.UserMetadata, err error) {
	if err = u.Conn().Transaction(func(tx *gorm.DB) error {
		var users []*entity.User
		var metadata *entity.UserMetadata

		users, err = u.repo.SearchUser(tx, input)
		if err != nil {
			return err
		}

		for _, user := range users {
			metadata, err = u.FindUserMetadata(user.ID)
			if err != nil {
				return err
			}
			result = append(result, metadata)
		}
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userService) BlockUser(id uint) (err error) {
	return u.Conn().Transaction(func(tx *gorm.DB) error {
		return u.repo.BlockUser(tx, id)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func (u *userService) UnBlockUser(id uint) (err error) {
	return u.Conn().Transaction(func(tx *gorm.DB) error {
		return u.repo.UnBlockUser(tx, id)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func (u *userService) CreateMemo(id uint, desc string) (err error) {
	return u.Conn().Transaction(func(tx *gorm.DB) error {
		return u.repo.CreateMemo(tx, id, desc)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}
