package global

import (
    "github.com/hd2yao/blog/pkg/logger"
    "github.com/hd2yao/blog/pkg/setting"
)

var (
    ServerSetting   *setting.ServerSettingS
    AppSetting      *setting.AppSettingS
    DatabaseSetting *setting.DatabaseSettingS
    Logger          *logger.Logger
)
