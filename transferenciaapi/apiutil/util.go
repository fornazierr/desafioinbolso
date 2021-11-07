package apiutil

import (
	"fmt"
	"os"
	"strings"
	"transferenciaapi/models"
)

func GetConfig() models.Config {
	c := models.Config{
		DB_NAME:    "",
		DB_PORT:    "",
		DB_USER:    "",
		DB_HOST:    "",
		DB_PASS:    "",
		URL_PORT:   "",
		URL_HOST:   "",
		CONTA_API:  "",
		BOLETO_API: "",
	}

	c.DB_NAME = os.Getenv("DB_NAME")
	c.DB_PASS = os.Getenv("DB_PASS")
	c.DB_PORT = os.Getenv("DB_PORT")
	c.DB_HOST = os.Getenv("DB_HOST")
	c.DB_USER = os.Getenv("DB_USER")
	c.URL_PORT = os.Getenv("URL_PORT")
	c.URL_PORT = os.Getenv("URL_PORT")
	c.CONTA_API = os.Getenv("CONTA_API")
	c.BOLETO_API = os.Getenv("BOLETO_API")

	args := os.Args[1:]

	for i := 0; i < len(args); i++ {
		arg := strings.Split(args[i], "=")
		switch strings.ToUpper(arg[0]) {
		case "DB_NAME":
			c.DB_NAME = arg[1]
		case "DB_PASS":
			c.DB_PASS = arg[1]
		case "DB_PORT":
			c.DB_PORT = arg[1]
		case "DB_HOST":
			c.DB_HOST = arg[1]
		case "DB_USER":
			c.DB_USER = arg[1]
		case "URL_PORT":
			c.URL_PORT = arg[1]
		case "URL_HOST":
			c.URL_HOST = arg[1]
		case "CONTA_API":
			c.CONTA_API = arg[1]
		case "BOLETO_API":
			c.BOLETO_API = arg[1]
		}
	}

	if c.DB_NAME == "" {
		c.DB_NAME = "inbolso"
	}
	if c.DB_PASS == "" {
		c.DB_PASS = "12345678"
	}
	if c.DB_PORT == "" {
		c.DB_PORT = "15123"
	}
	if c.DB_HOST == "" {
		c.DB_HOST = "localhost"
	}
	if c.DB_USER == "" {
		c.DB_USER = "postgres"
	}
	if c.URL_PORT == "" {
		c.URL_PORT = "15002"
	}
	if c.URL_HOST == "" {
		c.URL_HOST = "127.0.0.1"
	}
	if c.CONTA_API == "" {
		c.CONTA_API = "http://127.0.0.1:15001"
	}
	if c.BOLETO_API == "" {
		c.BOLETO_API = "http://127.0.0.1:15003"
	}

	return c
}

//Monta uma url de conexão como a seguinte "postgres://username:password@localhost:5432/database_name"
func GetUrlConnection() string {
	conf := GetConfig()
	base := "postgres://%s:%s@%s:%s/%s"
	return fmt.Sprintf(base, conf.DB_USER, conf.DB_PASS, conf.DB_HOST, conf.DB_PORT, conf.DB_NAME)
}
