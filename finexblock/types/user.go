package types

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/user"
	"github.com/shopspring/decimal"
	"time"
)

type Metadata struct {
	ID                uint            `json:"id" query:"id"`
	UUID              string          `json:"uuid" query:"uuid"`
	UserType          string          `json:"user_type" query:"user_type"`
	Nickname          string          `json:"nickname" query:"nickname"`
	Fullname          string          `json:"fullname" query:"fullname"`
	PhoneNumber       string          `json:"phone_number" query:"phone_number"`
	BTC               decimal.Decimal `json:"btc" query:"btc"`
	IsBlock           bool            `json:"is_block" query:"is_block"`
	IsDormant         bool            `json:"is_dormant" query:"is_dormant"`
	IsMetaverseUser   bool            `json:"is_metaverse_user" query:"is_metaverse_user"`
	IsGoogleUser      bool            `json:"is_google_user" query:"is_google_user"`
	IsAppleUser       bool            `json:"is_apple_user" query:"is_apple_user"`
	IsEmailSignUpUser bool            `json:"is_email_sign_up_user" query:"is_email_sign_up_user"`
	CreatedAt         time.Time       `json:"created_at" query:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at" query:"updated_at"`
	UserMemo          *user.UserMemo  `json:"user_memo" query:"user_memo"`
}