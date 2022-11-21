package infra

import (
	"errors"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Conf configuration
type Conf struct {
	HTTP HTTPConf
	DB   DBConf

	Timezone string
	Mode     string
	Instance string
}

// HTTPConf .
type HTTPConf struct {
	Port      string
	BasicAuth string
	AuthUser  AuthUser
}

// AuthUser .
type AuthUser struct {
	Name     string
	Password string
	Token    string
}

// DBConf .
type DBConf struct {
	Provider string
	Host     string
	Name     string
	User     string
	Pass     string
}

// LoadConf loads conf from environment variables or dotenv
func LoadConf(path string) (*Conf, error) {
	// try loading dotenv
	_ = godotenv.Load(path + ".env")
	conf := Conf{
		HTTP: HTTPConf{
			Port:      os.Getenv("HTTP_PORT"),
			BasicAuth: os.Getenv("HTTP_BASIC_AUTH"),
		},
		DB: DBConf{
			Provider: os.Getenv("DB_PROVIDER"),
			Host:     os.Getenv("DB_HOST"),
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Pass:     os.Getenv("DB_PASS"),
		},

		Instance: os.Getenv("INSTANCE"),
		Mode:     os.Getenv("MODE"),
		Timezone: os.Getenv("TIMEZONE"),
	}
	if len(conf.HTTP.BasicAuth) > 0 {
		fields := strings.Split(conf.HTTP.BasicAuth, ":")
		conf.HTTP.AuthUser.Name = fields[0]
		conf.HTTP.AuthUser.Password = fields[1]
		conf.HTTP.AuthUser.Token = fields[2]
	} else {
		return nil, errors.New("No basic auth users")
	}
	return &conf, nil
}

// AuthAccounts .
func (c Conf) AuthAccounts() gin.Accounts {
	return map[string]string{
		c.HTTP.AuthUser.Name: c.HTTP.AuthUser.Password,
	}
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
