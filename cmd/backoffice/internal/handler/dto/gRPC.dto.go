package dto

type (
	ProxyHealthCheckInput struct {
		Domain string `json:"domain" required:"true"`
	}

	ProxyHealthCheckOutput struct {
		Message string `json:"message" required:"true"`
	}

	ProxyWhoAmIInput struct {
		Domain string `json:"domain" required:"true"`
	}

	ProxyWhoAmIOutput struct {
		Message string `json:"message" required:"true"`
	}
)