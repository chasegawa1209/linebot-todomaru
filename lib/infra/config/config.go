package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Config 設定内容
type Config struct {
	Debug                   *Debug
	LineMessagingAPISetting *LineMessagingAPISetting
	ScheduleSetting         *ScheduleSetting
}

// Debug デバッグ設定
type Debug struct {
	IsDebug bool
}

// LineMessagingAPISetting LINEMessagingAPIの設定
type LineMessagingAPISetting struct {
	AccessToken string
	PushURL     string
	ReplyURL    string
	AdminID     string
	GroupID     string
}

// ScheduleSetting botが動くスケジュールの設定
type ScheduleSetting struct {
	Towel string
}

type CalendarID struct {
	ChisakiID string
	KensukeID string
}

// Load 設定ファイル読み込み
func Load(path string) (*Config, error) {
	viper.SetConfigType("yml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("config file error: %s", err)
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("config file error: %s", err)
	}

	return config, nil
}
