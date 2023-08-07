package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func Swagger(router fiber.Router) {
	router.Get("/api-docs/*", swagger.HandlerDefault)
	router.Get("/api-docs/*", swagger.New(swagger.Config{
		InstanceName: "Finexblock Backoffice API Documentation",
		Title:        "Finexblock Backoffice API Documentation",
		DocExpansion: "full",
		Filter:       swagger.FilterConfig{},
		TagsSorter:   "alpha",
		SyntaxHighlight: &swagger.SyntaxHighlightConfig{
			Activate: true,
			Theme:    "monokai",
		},
		RequestSnippetsEnabled: true,
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl:    "http://localhost:8000/login",
		WithCredentials:      true,
		PersistAuthorization: true,
		OAuth: &swagger.OAuthConfig{
			AppName:  "Finexblock Backoffice",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
	}))
}