package infra

import (
	"os"

	"github.com/joho/godotenv"
)

// Conf configuration
type Conf struct {
	HTTP  HTTPConf
	DB    DBConf
	Auth0 Auth0Conf

	Timezone string
	Mode     string
	Instance string
}

// HTTPConf .
type HTTPConf struct {
	Port  string
	Token string
}

// DBConf .
type DBConf struct {
	Provider string
	Host     string
	Name     string
	User     string
	Pass     string
}

// Auth0Conf .
type Auth0Conf struct {
	Domain       string
	ClientID     string
	ClientSecret string
	CallbackURL  string
	UserID       string
}

// LoadConf loads conf from environment variables or dotenv
func LoadConf(path string) (*Conf, error) {
	// try loading dotenv
	_ = godotenv.Load(path + ".env")
	conf := Conf{
		HTTP: HTTPConf{
			Port:  os.Getenv("HTTP_PORT"),
			Token: os.Getenv("HTTP_TOKEN"),
		},
		DB: DBConf{
			Provider: os.Getenv("DB_PROVIDER"),
			Host:     os.Getenv("DB_HOST"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Pass:     os.Getenv("DB_PASS"),
		},
		Auth0: Auth0Conf{
			Domain:       os.Getenv("AUTH0_DOMAIN"),
			ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
			ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
			CallbackURL:  os.Getenv("AUTH0_CALLBACK_URL"),
			UserID:       os.Getenv("AUTH0_USER_ID"),
		},

		Instance: os.Getenv("INSTANCE"),
		Mode:     os.Getenv("MODE"),
		Timezone: os.Getenv("TIMEZONE"),
	}
	return &conf, nil
}

// IsDev .
func (c Conf) IsDev() bool {
	return c.Mode == "dev"
}

// IsTesting .
func (c Conf) IsTesting() bool {
	return c.Mode == "testing"
}

// IsProduction .
func (c Conf) IsProduction() bool {
	return c.Mode == "production"
}
