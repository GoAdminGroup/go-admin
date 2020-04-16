// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package logger

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/mgutz/ansi"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var (
	defaultEncoderCfg = EncoderCfg{
		TimeKey:       "ts",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		Level:         "lowercaseLevelEncoder",
		Time:          "epochTimeEncoder",
		Duration:      "secondsDurationEncoder",
		Caller:        "shortCallerEncoder",
		Encoding:      "console",
	}

	logger = &Logger{
		rotate: RotateCfg{
			MaxSize:    10,
			MaxBackups: 5,
			MaxAge:     30,
			Compress:   false,
		},
		encoder: defaultEncoderCfg,
	}

	infoLevelEnabler = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel
	})

	errorLevelEnabler = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.ErrorLevel
	})

	accessLevelEnabler = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel
	})
)

func init() {
	logger.Init()
}

type Logger struct {
	logger        *zap.Logger
	sugaredLogger *zap.SugaredLogger

	infoLogOff   bool
	errorLogOff  bool
	accessLogOff bool

	accessAssetsLogOff bool

	debug bool

	sqlLogOpen bool

	infoLogPath   string
	errorLogPath  string
	accessLogPath string

	rotate  RotateCfg
	encoder EncoderCfg
}

type EncoderCfg struct {
	TimeKey       string
	LevelKey      string
	NameKey       string
	CallerKey     string
	MessageKey    string
	StacktraceKey string
	Level         string
	Time          string
	Duration      string
	Caller        string
	Encoding      string
}

type RotateCfg struct {
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

func (l *Logger) Init() {
	zapLogger := zap.New(zapcore.NewTee(
		zapcore.NewCore(l.getEncoder(), l.getLogWriter(l.infoLogPath), infoLevelEnabler),
		zapcore.NewCore(l.getEncoder(), l.getLogWriter(l.errorLogPath), errorLevelEnabler),
		zapcore.NewCore(l.getEncoder(), l.getLogWriter(l.accessLogPath), accessLevelEnabler),
	), zap.AddCaller())
	l.sugaredLogger = zapLogger.Sugar()
	l.logger = zapLogger
}

func (l *Logger) getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        l.encoder.TimeKey,
		LevelKey:       l.encoder.LevelKey,
		NameKey:        l.encoder.NameKey,
		CallerKey:      l.encoder.CallerKey,
		MessageKey:     l.encoder.MessageKey,
		StacktraceKey:  l.encoder.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    filterZapEncodeLevel(l.encoder.Level),
		EncodeTime:     filterZapTimeEncoder(l.encoder.Time),
		EncodeDuration: filterZapDurationEncoder(l.encoder.Duration),
		EncodeCaller:   filterZapCallerEncoder(l.encoder.Caller),
	}

	return filterZapEncoder(l.encoder.Encoding, encoderConfig)
}

func (l *Logger) getLogWriter(path string) zapcore.WriteSyncer {
	if path != "" {
		lumberJackLogger := &lumberjack.Logger{
			Filename:   path,
			MaxSize:    l.rotate.MaxSize,
			MaxBackups: l.rotate.MaxBackups,
			MaxAge:     l.rotate.MaxAge,
			Compress:   l.rotate.Compress,
		}
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(os.Stdout)
}

func (l *Logger) SetRotate(cfg RotateCfg) {
	if cfg.MaxSize != 0 && cfg.MaxAge != 0 && cfg.MaxBackups != 0 {
		l.rotate = cfg
	}
}

func (l *Logger) SetEncoder(cfg EncoderCfg) {
	cfg.TimeKey = utils.SetDefault(cfg.TimeKey, "", defaultEncoderCfg.TimeKey)
	cfg.LevelKey = utils.SetDefault(cfg.LevelKey, "", defaultEncoderCfg.LevelKey)
	cfg.NameKey = utils.SetDefault(cfg.NameKey, "", defaultEncoderCfg.NameKey)
	cfg.CallerKey = utils.SetDefault(cfg.CallerKey, "", defaultEncoderCfg.CallerKey)
	cfg.MessageKey = utils.SetDefault(cfg.MessageKey, "", defaultEncoderCfg.MessageKey)
	cfg.StacktraceKey = utils.SetDefault(cfg.StacktraceKey, "", defaultEncoderCfg.StacktraceKey)
	cfg.Level = utils.SetDefault(cfg.Level, "", defaultEncoderCfg.Level)
	cfg.Time = utils.SetDefault(cfg.Time, "", defaultEncoderCfg.Time)
	cfg.Duration = utils.SetDefault(cfg.Duration, "", defaultEncoderCfg.Duration)
	cfg.Caller = utils.SetDefault(cfg.Caller, "", defaultEncoderCfg.Caller)
	cfg.Encoding = utils.SetDefault(cfg.Encoding, "", defaultEncoderCfg.Encoding)
	l.encoder = cfg
}

type Config struct {
	InfoLogOff   bool
	ErrorLogOff  bool
	AccessLogOff bool

	SqlLogOpen bool

	InfoLogPath   string
	ErrorLogPath  string
	AccessLogPath string

	AccessAssetsLogOff bool

	Rotate RotateCfg
	Encode EncoderCfg

	Debug bool
}

func InitWithConfig(cfg Config) {
	logger.infoLogPath = cfg.InfoLogPath
	logger.infoLogOff = cfg.InfoLogOff
	logger.errorLogPath = cfg.ErrorLogPath
	logger.errorLogOff = cfg.ErrorLogOff
	logger.accessLogPath = cfg.AccessLogPath
	logger.accessLogOff = cfg.AccessLogOff
	logger.sqlLogOpen = cfg.SqlLogOpen
	logger.accessAssetsLogOff = cfg.AccessAssetsLogOff
	logger.debug = cfg.Debug
	logger.SetRotate(cfg.Rotate)
	logger.SetEncoder(cfg.Encode)
	logger.Init()
}

func SetRotate(cfg RotateCfg) {
	logger.rotate = cfg
	logger.Init()
}

// SetInfoLogger set the info logger.
func SetInfoLogger(path string, isInfoLogOff bool) {
	logger.infoLogPath = path
	logger.infoLogOff = isInfoLogOff
	logger.Init()
}

// SetErrorLogger set the error logger.
func SetErrorLogger(path string, isErrorLogOff bool) {
	logger.errorLogPath = path
	logger.errorLogOff = isErrorLogOff
	logger.Init()
}

// SetAccessLogger set the access logger.
func SetAccessLogger(path string, isAccessLogOff bool) {
	logger.accessLogPath = path
	logger.accessLogOff = isAccessLogOff
	logger.Init()
}

// OpenSQLLog set the sqlLogOpen true.
func OpenSQLLog() {
	logger.sqlLogOpen = true
}

// Error print the error message.
func Error(err ...interface{}) {
	if !logger.errorLogOff {
		logger.sugaredLogger.Error(err...)
	}
}

// Info print the info message.
func Info(info ...interface{}) {
	if !logger.infoLogOff {
		logger.sugaredLogger.Info(info...)
	}
}

// Warn print the warning message.
func Warn(info ...interface{}) {
	if !logger.infoLogOff {
		logger.sugaredLogger.Warn(info...)
	}
}

// Access print the access message.
func Access(ctx *context.Context) {
	if !logger.accessLogOff {
		if logger.accessAssetsLogOff {
			if filepath.Ext(ctx.Path()) == "" {
				logger.sugaredLogger.Debug("[GoAdmin]",
					ansi.Color(" "+strconv.Itoa(ctx.Response.StatusCode)+" ", "white:blue"),
					ansi.Color(" "+string(ctx.Method()[:])+"   ", "white:blue+h"),
					ctx.Path())
			}
		} else {
			logger.sugaredLogger.Debug("[GoAdmin]",
				ansi.Color(" "+strconv.Itoa(ctx.Response.StatusCode)+" ", "white:blue"),
				ansi.Color(" "+string(ctx.Method()[:])+"   ", "white:blue+h"),
				ctx.Path())
		}
	}
}

// LogSQL print the sql info message.
func LogSQL(statement string, args []interface{}) {
	if logger.sqlLogOpen && statement != "" {
		logger.sugaredLogger.Info("[GoAdmin]", "statement", statement, "args", args)
	}
}

// default ISO8601TimeEncoder
func filterZapTimeEncoder(timeEncoder string) func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	switch timeEncoder {
	default:
		return zapcore.ISO8601TimeEncoder
	case "ISO8601TimeEncoder":
		return zapcore.ISO8601TimeEncoder
	case "epochMillisTimeEncoder":
		return zapcore.EpochMillisTimeEncoder
	case "epochNanosTimeEncoder":
		return zapcore.EpochNanosTimeEncoder
	case "epochTimeEncoders":
		return zapcore.EpochTimeEncoder
	}
}

// default SecondsDurationEncoder
func filterZapDurationEncoder(encodeDuration string) func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	switch encodeDuration {
	default:
		return zapcore.SecondsDurationEncoder
	case "secondsDurationEncoder":
		return zapcore.SecondsDurationEncoder
	case "nanosDurationEncoder":
		return zapcore.NanosDurationEncoder
	case "stringDurationEncoder":
		return zapcore.StringDurationEncoder
	}
}

// default FullCallerEncoder
func filterZapCallerEncoder(encodeCaller string) func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	switch encodeCaller {
	default:
		return zapcore.FullCallerEncoder
	case "fullCallerEncoder":
		return zapcore.FullCallerEncoder
	case "shortCallerEncoder":
		return zapcore.ShortCallerEncoder
	}

}

// default CapitalLevelEncoder
func filterZapEncodeLevel(encodeLevel string) func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch encodeLevel {
	default:
		return zapcore.CapitalLevelEncoder
	case "capitalLevelEncoder":
		return zapcore.CapitalLevelEncoder
	case "capitalColorLevelEncoder":
		return zapcore.CapitalColorLevelEncoder
	case "lowercaseLevelEncoder":
		return zapcore.LowercaseLevelEncoder
	case "lowercaseColorLevelEncoder":
		return zapcore.LowercaseColorLevelEncoder
	}
}

func filterZapEncoder(encoding string, encoderConfig zapcore.EncoderConfig) zapcore.Encoder {
	var encoder zapcore.Encoder
	switch encoding {
	default:
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	case "json":
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	case "console":
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
	return encoder
}

func filterZapAtomicLevelByViper(level int) zapcore.Level {
	var atomViper zapcore.Level
	switch level {
	default:
		atomViper = zap.InfoLevel
	case -1:
		atomViper = zap.DebugLevel
	case 0:
		atomViper = zap.InfoLevel
	case 1:
		atomViper = zap.WarnLevel
	case 2:
		atomViper = zap.ErrorLevel
	}
	return atomViper
}
