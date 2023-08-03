package dto

type (
	AdminRegisterInput struct {
		Email    string `json:"email" bind:"required" default:"finexblock@gmail.com"`
		Password string `json:"password" bind:"required" default:"Metaverse123!"`
	}

	AdminRegisterOutput struct {
	}
)

type (
	LoginInput struct {
		Email    string `json:"email" bind:"required" default:"finexblock@gmail.com"`
		Password string `json:"password" bind:"required" default:"Metaverse123!"`
	}

	LoginOutput struct {
		AccessToken string `json:"accessToken" bind:"required"`
	}
)
