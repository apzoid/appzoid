package config

import (
	"fmt"
	"os"
	"path/filepath"
)

type AppConfiguration struct {
	App_Name        string
	App_Upload_Path string
	App_Upload_Size int
	App_Env         string
	App_Key         string
	App_Url         string
	App_Port        string
}

var AppConfig *AppConfiguration //nolint:gochecknoglobals

func LoadAppConfig() {
	loadDefaultConfig()
	ViperConfig.Unmarshal(&AppConfig)
	if AppConfig.App_Url == "" {
		AppConfig.App_Url = fmt.Sprintf("http://localhost:%s", AppConfig.App_Port)
	}
	AppConfig.App_Upload_Path = filepath.Join(".", AppConfig.App_Upload_Path)
	AppConfig.App_Upload_Size = AppConfig.App_Upload_Size * 1024 * 1024
	if _, err := os.Stat(AppConfig.App_Upload_Path); os.IsNotExist(err) {
		os.MkdirAll(AppConfig.App_Upload_Path, os.ModePerm)
	}
}

func loadDefaultConfig() {
	ViperConfig.SetDefault("APP_NAME", "Appzoid")
	ViperConfig.SetDefault("APP_ENV", "dev")
	ViperConfig.SetDefault("APP_UPLOAD_PATH", "uploads")
	ViperConfig.SetDefault("APP_UPLOAD_SIZE", 4)
	ViperConfig.SetDefault("APP_KEY", "1894cde6c936a294a478cff0a9227fd276d86df6573b51af5dc59c9064edf426")
	ViperConfig.SetDefault("APP_PORT", "8080")
}
