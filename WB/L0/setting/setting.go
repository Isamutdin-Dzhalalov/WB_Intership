package setting

import (
	"encoding/json"
	"log"
	"os"
)

type setting struct {
	ServerHost string
	ServerPort string
	PgHost     string
	PgPort     string
	PgUser     string
	PgPassword string
	PgNameDB   string
	HTML       string
}

var Config setting

func init() {
	file, err := os.Open("setting.cfg")
	if err != nil {
		log.Fatal("setting.Config -> os.Open(): ", err)
		os.Exit(1)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Fatal("setting.Config -> file.Stat(): ", err)
		os.Exit(1)
	}

	readByte := make([]byte, stat.Size())

	_, err = file.Read(readByte)
	if err != nil {
		log.Fatal("setting.Config -> file.Read(): ", err)
		os.Exit(1)
	}
	err = json.Unmarshal(readByte, &Config)
	if err != nil {
		log.Fatal("setting.Config -> json.Unmarshal(): ", err)
		os.Exit(1)
	}
}
