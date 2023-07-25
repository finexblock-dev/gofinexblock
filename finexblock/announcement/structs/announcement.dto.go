package structs

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
)

type (
	FindAllAnnouncementInput struct {
		Limit  int `json:"limit,required" query:"limit,required" default:"20"`
		Offset int `json:"offset,required" query:"offset,required" default:"0"`
	}

	FindAllAnnouncementOutput struct {
		Result []*entity.Announcement `json:"result,required"`
	}

	FindAllAnnouncementSuccessResponse struct {
		Code int                       `json:"code,required" default:"200"`
		Data FindAllAnnouncementOutput `json:"data,required"`
	}
)

type (
	FindAnnouncementByIDInput struct {
		ID uint `json:"id,required" query:"id,required"`
	}

	FindAnnouncementByIDOutput struct {
		Result *entity.Announcement `json:"result,required"`
	}

	FindAnnouncementByIDSuccessResponse struct {
		Code int                        `json:"code,required" default:"200"`
		Data FindAnnouncementByIDOutput `json:"data,required"`
	}
)

type (
	SearchAnnouncementInput struct {
		Word       string `json:"word,required" query:"word,required"`
		Title      string `json:"title,required" query:"title,required"`
		Visible    bool   `json:"visible,required" query:"visible,required"`
		Pinned     bool   `json:"pinned,required" query:"pinned,required"`
		CategoryID uint   `json:"category_id,required" query:"category_id,required"`
		Limit      int    `json:"limit,required" query:"limit,required" default:"20"`
		Offset     int    `json:"offset,required" query:"offset,required" default:"0"`
	}

	SearchAnnouncementOutput struct {
		Result []*entity.Announcement `json:"result,required"`
	}

	SearchAnnouncementSuccessResponse struct {
		Code int                       `json:"code,required" default:"200"`
		Data *SearchAnnouncementOutput `json:"data,required"`
	}
)

type (
	CreateAnnouncementInput struct {
		CategoryID   uint   `json:"category_id,required" default:"1"`
		KoreanTitle  string `json:"korean_title,required" default:"테스트" `
		EnglishTitle string `json:"english_title,required" default:"test"`
		ChineseTitle string `json:"chinese_title,required" default:"测验"`
		Korean       string `json:"korean,required" default:"공지사항 테스트입니다."`
		English      string `json:"english,required" default:"This is an announcement test."`
		Chinese      string `json:"chinese,required" default:"是公告事项测试"`
		Visible      bool   `json:"visible,required" default:"true" `
		Pinned       bool   `json:"pinned,required" default:"true" `
	}

	CreateAnnouncementOutput struct {
		Msg string `json:"msg,required"`
	}

	CreateAnnouncementSuccessResponse struct {
		Code int                      `json:"code,required" default:"200"`
		Data CreateAnnouncementOutput `json:"data,required"`
	}
)

type (
	UpdateAnnouncementInput struct {
		ID           uint   `json:"id,required" default:"1"`
		CategoryID   uint   `json:"category_id,required"`
		KoreanTitle  string `json:"korean_title,required" default:"테스트" `
		EnglishTitle string `json:"english_title,required" default:"test"`
		ChineseTitle string `json:"chinese_title,required" default:"测验"`
		Korean       string `json:"korean,required"`
		English      string `json:"english,required"`
		Chinese      string `json:"chinese,required"`
		Visible      bool   `json:"visible,required"  default:"true"`
		Pinned       bool   `json:"pinned,required" default:"true"`
	}

	UpdateAnnouncementOutput struct {
		Msg string `json:"msg,required" default:"Successfully updated"`
	}

	UpdateAnnouncementSuccessResponse struct {
		Code int                      `json:"code,required" default:"200"`
		Data UpdateAnnouncementOutput `json:"data,required"`
	}
)

type (
	DeleteAnnouncementInput struct {
		AnnouncementID uint `json:"announcement_id,required" query:"announcement_id,required" default:"1"`
	}

	DeleteAnnouncementOutput struct {
		Msg string `json:"msg,required" default:"Successfully deleted"`
	}

	DeleteAnnouncementSuccessResponse struct {
		Code int                      `json:"code,required" default:"200"`
		Data DeleteAnnouncementOutput `json:"data,required"`
	}
)

type (
	CreateCategoryInput struct {
		KoreanType  string `json:"korean_type,required" default:"남현우"`
		EnglishType string `json:"english_type,required" default:"Hyunwoo Nam"`
		ChineseType string `json:"chinese_type,required" default:"南贤宇"`
	}

	CreateCategoryOutput struct {
		Msg string `json:"msg,required" default:"Successfully created"`
	}

	CreateCategorySuccessResponse struct {
		Code int                  `json:"code,required" default:"200"`
		Data CreateCategoryOutput `json:"data,required"`
	}
)

type (
	FindAllCategoryOutput struct {
		Result []*entity.AnnouncementCategory `json:"result,required"`
	}

	FindAllCategorySuccessResponse struct {
		Code int                   `json:"code,required" default:"200"`
		Data FindAllCategoryOutput `json:"data,required"`
	}
)

type (
	UpdateCategoryInput struct {
		ID          uint   `json:"id,required" default:"1"`
		KoreanType  string `json:"korean_type,required" default:"의자"`
		EnglishType string `json:"english_type,required" default:"chair"`
		ChineseType string `json:"chinese_type,required" default:"椅子"`
	}

	UpdateCategoryOutput struct {
		Msg string `json:"msg,required" default:"Successfully updated"`
	}

	UpdateCategorySuccessResponse struct {
		Code int                  `json:"code,required" default:"200"`
		Data UpdateCategoryOutput `json:"data,required"`
	}
)

type (
	DeleteCategoryInput struct {
		ID uint `json:"id,required" default:"1"`
	}

	DeleteCategoryOutput struct {
		Msg string `json:"msg,required" default:"Successfully deleted"`
	}

	DeleteCategorySuccessResponse struct {
		Code int                  `json:"code,required" default:"200"`
		Data DeleteCategoryOutput `json:"data,required"`
	}
)
