package config

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var configInis *ini.File

// 配置文件路径 config file path.
const (
	APP_INI_PATH = "config/app.ini"
	DATABASE_INI_PATH = "config/database.ini"
	REDIS_INI_PATH = "config/redis.ini"
	SERVER_INI_PATH = "config/server.ini"
)

type App struct {
	PageSize int
	JwtSecret string

	RuntimeRootPath string
	UploadRootPath string
	ImageSavePath string
	FileSavePath string

	ImageMaxSize int
	ImageAllowExts []string
	FileMaxSize int
	FileAllowExts []int

	LogSavePath string
	LogSaveName string
	LogFileExt string
	LogTimeFormat string
}

var AppConfig = &App{}

type server struct {
	RunMode string
	HttpPort string
	ReadTimeout time.Duration
	WriteTimeout time.Duration

	StaticPath string
}

var ServerConfig = &server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string


	MaxIdleConns int
	MaxOpenConns int
}

var DatabaseConfig = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisConfig = &Redis{}

func AppInit()  {
	var err error
	// 获取配置目录下所有文件
	configInis, err = ini.LooseLoad(APP_INI_PATH, DATABASE_INI_PATH, REDIS_INI_PATH, SERVER_INI_PATH)
	if err != nil {
		log.Fatalf("load configs, fail to parse : %v", err)
	}
	mapToSection("app", AppConfig)
	mapToSection("server", ServerConfig)
	mapToSection("database", DatabaseConfig)
	mapToSection("redis", RedisConfig)

	AppConfig.ImageMaxSize = AppConfig.ImageMaxSize * 1024 * 1024
	AppConfig.FileMaxSize = AppConfig.FileMaxSize * 1024 * 1024
	AppConfig.ImageSavePath = AppConfig.UploadRootPath + AppConfig.ImageSavePath
	AppConfig.FileSavePath = AppConfig.FileSavePath + AppConfig.FileSavePath
	ServerConfig.ReadTimeout = ServerConfig.ReadTimeout * time.Second
	ServerConfig.WriteTimeout = ServerConfig.WriteTimeout * time.Second
	RedisConfig.IdleTimeout = RedisConfig.IdleTimeout * time.Second
}

/**
映射配置.
 */
func mapToSection(section string, v interface{})  {
	err := configInis.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("config ini mapto %s err : %v", section, err)
	}
}
