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
		result = new(entity.UserMetadata)

		_user, err = u.repo.FindUserByID(tx, id)
		if err != nil {
			return err
		}

		result.ID = _user.ID
		result.UUID = _user.UUID
		result.UserType = _user.UserType
		result.CreatedAt = _user.CreatedAt
		result.UpdatedAt = _user.UpdatedAt

		_profile, err = u.repo.FindUserProfileByUserID(tx, id)
		if err != nil {
			return err
		}

		result.Nickname = _profile.Nickname
		result.Fullname = _profile.Fullname
		result.PhoneNumber = _profile.PhoneNumber

		metaverseSSO, err = u.repo.FindUserSingleSignOnInfoByCond(tx, id, entity.Metaverse)
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		if metaverseSSO != nil {
			result.IsMetaverseUser = metaverseSSO != nil && metaverseSSO.SSOType == entity.Metaverse
		} else {
			result.IsMetaverseUser = false
		}

		appleSSO, err = u.repo.FindUserSingleSignOnInfoByCond(tx, id, entity.Apple)
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		if appleSSO != nil {
			result.IsAppleUser = appleSSO != nil && appleSSO.SSOType == entity.Apple
		} else {
			result.IsAppleUser = false
		}

		googleSSO, err = u.repo.FindUserSingleSignOnInfoByCond(tx, id, entity.Google)
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		if googleSSO != nil {
			result.IsGoogleUser = googleSSO != nil && googleSSO.SSOType == entity.Google
		} else {
			result.IsGoogleUser = false
		}

		memo, err = u.repo.FindUserMemoByUserID(tx, id)
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		if memo != nil {
			result.UserMemo = memo
		} else {
			result.UserMemo = nil
		}

		dormant, err = u.repo.FindUserDormantByUserID(tx, id)
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		if dormant != nil {
			result.IsDormant = true
		} else {
			result.IsDormant = false
		}

		// Calculate the BTC balance
		btcTotal, err = u.manager.GetBalance(_user.UUID, grpc_order.Currency_BTC.String())
		if err != nil && !errors.Is(err, redis.Nil) {
			return err
		}

		result.BTC = btcTotal

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

func newUserService(db *gorm.DB, cluster *redis.ClusterClient) *userService {
	return &userService{repo: NewRepository(db), manager: trade.New(cluster)}
}