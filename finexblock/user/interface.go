package user

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/finexblock-dev/gofinexblock/finexblock/user/dto"
	"gorm.io/gorm"
)

type Repository interface {
	types.Repository
	FindUserByUUID(tx *gorm.DB, uuid string) (result *entity.User, err error)
	FindUserByUUIDs(tx *gorm.DB, uuids []string) (result []*entity.User, err error)
	FindUserByID(tx *gorm.DB, id uint) (result *entity.User, err error)
	FindUserMetadata(tx *gorm.DB, id uint) (result *types.Metadata, err error)
	SearchUser(tx *gorm.DB, input *dto.SearchUserInput) (result []*types.Metadata, err error)

	BlockUser(tx *gorm.DB, id uint) (err error)
	UnBlockUser(tx *gorm.DB, id uint) (err error)

	CreateMemo(tx *gorm.DB, id uint, desc string) (err error)

	FindUserProfileByUserID(tx *gorm.DB, userID uint) (result *entity.UserProfile, err error)
	FindUserSingleSignOnInfoByCond(tx *gorm.DB, userID uint, ssoType entity.SSOType) (result *entity.UserSingleSignOnInfo, err error)
	FindUserEmailSignUpByUserID(tx *gorm.DB, userID uint) (result *entity.UserEmailSignUp, err error)

	FindUserMemoByUserID(tx *gorm.DB, userID uint) (result *entity.UserMemo, err error)

	FindUserDormantByUserID(tx *gorm.DB, userID uint) (result *entity.UserDormant, err error)
}

type Service interface {
	types.Service
	FindUserByUUID(uuid string) (result *entity.User, err error)
	FindUserByUUIDs(uuids []string) (result []*entity.User, err error)
	FindUserByID(id uint) (result *entity.User, err error)
	FindUserMetadata(id uint) (result *types.Metadata, err error)
	SearchUser(input *dto.SearchUserInput) (result []*types.Metadata, err error)

	BlockUser(id uint) (err error)
	UnBlockUser(id uint) (err error)

	CreateMemo(id uint, desc string) (err error)
}

func NewRepository(db *gorm.DB) Repository {
	return newUserRepository(db)
}

func NewService(db *gorm.DB) Service {
	return newUserService(NewRepository(db))
}
