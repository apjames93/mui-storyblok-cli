package config

import "fmt"

// Config is the publically available API
type Config struct {
	storyBlokSpaceID   string
	storyBlokOathToken string
	muiThemeURL        string
	accessToken        string
}

// New returns a pointer to a new Config
func New(
	storyBlokSpaceID,
	storyBlokOathToken,
	muiThemeURL,
	accessToken string,
) *Config {
	return &Config{
		storyBlokSpaceID:   storyBlokSpaceID,
		storyBlokOathToken: storyBlokOathToken,
		muiThemeURL:        muiThemeURL,
		accessToken:        accessToken,
	}
}

// GenerateEnvString creates the string needed for writing to the env file
func (c Config) GenerateEnvString() string {
	return fmt.Sprintf(
		"STORYBLOK_OAUTH_TOKEN=%s\nSTORYBLOK_SPACE_ID=%s\nMUI_THEME_URL=%s\nACCESS_TOKEN=%s",
		c.storyBlokOathToken,
		c.storyBlokSpaceID,
		c.muiThemeURL,
		c.accessToken,
	)
}
