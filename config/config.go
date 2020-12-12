package config

import (
	"github.com/go-ini/ini"
	"gowith/logging"
	"time"
)

var(
	Cfg *ini.File

	RunMode string

	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration

	JwtSecret string

	DBType string
	DBName string
	User string
	Password string
	Host string
)

func init() {
	var err error
	Cfg, err = ini.Load("config/config.ini")
	if err != nil{
		logging.Info("Failed to parse `config.ini`, ", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
	LoadDatabase()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil{
		logging.Fatal("Failed to get section 'app': ", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!WE@#%%^&&*")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil{
		logging.Fatal("failed to get section `server`: ", err)
	}
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadDatabase() {
	sec, err := Cfg.GetSection("database")
	if err != nil{
		logging.Fatal(2, " Failed to get section 'database': %v", err)
	}

	DBType = sec.Key("TYPE").String()
	DBName = sec.Key("NAME").String()
	User = sec.Key("USER").String()
	Password = sec.Key("PASSWORD").String()
	Host = sec.Key("HOST").String()
}