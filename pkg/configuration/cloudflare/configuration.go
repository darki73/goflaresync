package cloudflare

// Configuration is the Cloudflare configuration.
type Configuration struct {
	// Email is the email of the Cloudflare account.
	Email string `json:"email" yaml:"email" xml:"email" toml:"email" mapstructure:"email"`
	// Token is the token of the Cloudflare account.
	Token string `json:"token" yaml:"token" xml:"token" toml:"token" mapstructure:"token"`
}

// InitializeWithDefaults initializes the configuration with defaults.
func InitializeWithDefaults() *Configuration {
	return &Configuration{
		Email: "",
		Token: "",
	}
}

// GetEmail returns the email of the Cloudflare account.
func (configuration *Configuration) GetEmail() string {
	return configuration.Email
}

// GetToken returns the token of the Cloudflare account.
func (configuration *Configuration) GetToken() string {
	return configuration.Token
}
