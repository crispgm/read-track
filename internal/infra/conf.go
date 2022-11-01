package infra

import (
	"errors"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Conf configuration
type Conf struct {
	HTTP HTTPConf
	DB   DBConf
}

// HTTPConf .
type HTTPConf struct {
	Port      string
	BasicAuth string
	AuthUsers map[string]string
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
	}
	if len(conf.HTTP.BasicAuth) > 0 {
		conf.HTTP.AuthUsers = make(map[string]string)
		pairs := strings.Split(conf.HTTP.BasicAuth, ";")
		for _, p := range pairs {
			fields := strings.Split(p, ":")
			conf.HTTP.AuthUsers[fields[0]] = fields[1]
		}
	}
	if len(conf.HTTP.AuthUsers) == 0 {
		return nil, errors.New("No basic auth users")
	}
	return &conf, nil
}
