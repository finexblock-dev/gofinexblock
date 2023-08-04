package dto

type (
	FindAllAnnouncementInput struct {
		Limit  int `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset int `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}
)

type (
	FindAnnouncementByIDInput struct {
		ID uint `json:"id" query:"id" binding:"required"`
	}
)

type (
	SearchAnnouncementInput struct {
		Word       string `json:"word" query:"word"`
		Title      string `json:"title" query:"title"`
		Visible    bool   `json:"visible" query:"visible"`
		Pinned     bool   `json:"pinned" query:"pinned"`
		CategoryID uint   `json:"categoryId" query:"categoryId"`
		Limit      int    `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset     int    `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}
)

type (
	CreateAnnouncementInput struct {
		CategoryID   uint   `json:"categoryId" default:"1"`
		KoreanTitle  string `json:"koreanTitle" default:"테스트" `
		EnglishTitle string `json:"englishTitle" default:"test"`
		ChineseTitle string `json:"chineseTitle" default:"测验"`
		Korean       string `json:"korean" default:"공지사항 테스트입니다."`
		English      string `json:"english" default:"This is an announcement test."`
		Chinese      string `json:"chinese" default:"是公告事项测试"`
		Visible      bool   `json:"visible" default:"true" `
		Pinned       bool   `json:"pinned" default:"true" `
	}
)

type (
	UpdateAnnouncementInput struct {
		ID           uint   `json:"ID" default:"1"`
		CategoryID   uint   `json:"categoryId"`
		KoreanTitle  string `json:"koreanTitle" default:"테스트" `
		EnglishTitle string `json:"englishTitle" default:"test"`
		ChineseTitle string `json:"chineseTitle" default:"测验"`
		Korean       string `json:"korean" default:"공지사항 테스트입니다."`
		English      string `json:"english" default:"This is an announcement test."`
		Chinese      string `json:"chinese" default:"是公告事项测试"`
		Visible      bool   `json:"visible"  default:"true"`
		Pinned       bool   `json:"pinned" default:"true"`
	}
)

type (
	DeleteAnnouncementInput struct {
		AnnouncementID uint `json:"announcementId" query:"announcementId" default:"1"`
	}
)

type (
	CreateCategoryInput struct {
		KoreanType  string `json:"koreanType" default:"남현우"`
		EnglishType string `json:"englishType" default:"Hyunwoo Nam"`
		ChineseType string `json:"chineseType" default:"南贤宇"`
	}
)

type (
	FindAllCategoryInput struct {
		Limit  int `json:"limit" query:"limit" binding:"required" default:"20" validate:"min=1,max=100"`
		Offset int `json:"offset" query:"offset" binding:"required" default:"0" validate:"min=0"`
	}
)

type (
	UpdateCategoryInput struct {
		ID          uint   `json:"id" default:"1"`
		KoreanType  string `json:"koreanType" default:"의자"`
		EnglishType string `json:"englishType" default:"chair"`
		ChineseType string `json:"chineseType" default:"椅子"`
	}
)

type (
	DeleteCategoryInput struct {
		ID uint `json:"id" default:"1"`
	}
)
