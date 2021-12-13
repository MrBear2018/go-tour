package global

import (
	"github.com/go-tour/blog-service/utils/config"
	"github.com/go-tour/blog-service/utils/logger"
)

// 全局变量，保存全局配置
var (
	ServerSetting   *config.ServerSettingS
	AppSetting      *config.AppSettingS
	DatabaseSetting *config.DatabaseSettingS
	Logger          *logger.Logger
)
