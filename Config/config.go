package Config

import (
	"github.com/joho/godotenv"
	"os"
)

type Conf struct {
	DbUri string
	Port  string
}

var Config *Conf

func Init() error {
	err := godotenv.Load(".env")
	if err == nil {
		var c Conf
		c.DbUri = os.Getenv("DB_URI")
		c.Port = os.Getenv("PORT")
		Config = &c
		return nil
	} else {
		return err
	}
}
