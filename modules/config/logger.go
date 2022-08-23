package config

import "github.com/GoAdminGroup/go-admin/modules/logger"

type Logger struct {
	Encoder EncoderCfg `json:"encoder,omitempty" yaml:"encoder,omitempty" ini:"encoder,omitempty"`
	Rotate  RotateCfg  `json:"rotate,omitempty" yaml:"rotate,omitempty" ini:"rotate,omitempty"`
	Level   int8       `json:"level,omitempty" yaml:"level,omitempty" ini:"level,omitempty"`
}

type EncoderCfg struct {
	TimeKey       string `json:"time_key,omitempty" yaml:"time_key,omitempty" ini:"time_key,omitempty"`
	LevelKey      string `json:"level_key,omitempty" yaml:"level_key,omitempty" ini:"level_key,omitempty"`
	NameKey       string `json:"name_key,omitempty" yaml:"name_key,omitempty" ini:"name_key,omitempty"`
	CallerKey     string `json:"caller_key,omitempty" yaml:"caller_key,omitempty" ini:"caller_key,omitempty"`
	MessageKey    string `json:"message_key,omitempty" yaml:"message_key,omitempty" ini:"message_key,omitempty"`
	StacktraceKey string `json:"stacktrace_key,omitempty" yaml:"stacktrace_key,omitempty" ini:"stacktrace_key,omitempty"`
	Level         string `json:"level,omitempty" yaml:"level,omitempty" ini:"level,omitempty"`
	Time          string `json:"time,omitempty" yaml:"time,omitempty" ini:"time,omitempty"`
	Duration      string `json:"duration,omitempty" yaml:"duration,omitempty" ini:"duration,omitempty"`
	Caller        string `json:"caller,omitempty" yaml:"caller,omitempty" ini:"caller,omitempty"`
	Encoding      string `json:"encoding,omitempty" yaml:"encoding,omitempty" ini:"encoding,omitempty"`
}

type RotateCfg struct {
	MaxSize    int  `json:"max_size,omitempty" yaml:"max_size,omitempty" ini:"max_size,omitempty"`
	MaxBackups int  `json:"max_backups,omitempty" yaml:"max_backups,omitempty" ini:"max_backups,omitempty"`
	MaxAge     int  `json:"max_age,omitempty" yaml:"max_age,omitempty" ini:"max_age,omitempty"`
	Compress   bool `json:"compress,omitempty" yaml:"compress,omitempty" ini:"compress,omitempty"`
}

func initLogger(cfg *Config) {
	logger.InitWithConfig(logger.Config{
		InfoLogOff:         cfg.InfoLogOff,
		ErrorLogOff:        cfg.ErrorLogOff,
		AccessLogOff:       cfg.AccessLogOff,
		SqlLogOpen:         cfg.SqlLog,
		InfoLogPath:        cfg.InfoLogPath,
		ErrorLogPath:       cfg.ErrorLogPath,
		AccessLogPath:      cfg.AccessLogPath,
		AccessAssetsLogOff: cfg.AccessAssetsLogOff,
		Rotate: logger.RotateCfg{
			MaxSize:    cfg.Logger.Rotate.MaxSize,
			MaxBackups: cfg.Logger.Rotate.MaxBackups,
			MaxAge:     cfg.Logger.Rotate.MaxAge,
			Compress:   cfg.Logger.Rotate.Compress,
		},
		Encode: logger.EncoderCfg{
			TimeKey:       cfg.Logger.Encoder.TimeKey,
			LevelKey:      cfg.Logger.Encoder.LevelKey,
			NameKey:       cfg.Logger.Encoder.NameKey,
			CallerKey:     cfg.Logger.Encoder.CallerKey,
			MessageKey:    cfg.Logger.Encoder.MessageKey,
			StacktraceKey: cfg.Logger.Encoder.StacktraceKey,
			Level:         cfg.Logger.Encoder.Level,
			Time:          cfg.Logger.Encoder.Time,
			Duration:      cfg.Logger.Encoder.Duration,
			Caller:        cfg.Logger.Encoder.Caller,
			Encoding:      cfg.Logger.Encoder.Encoding,
		},
		Debug: cfg.Debug,
		Level: cfg.Logger.Level,
	})
}
