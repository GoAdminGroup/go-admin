// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package logger

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/mgutz/ansi"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	defaultEncoderCfg = EncoderCfg{
		TimeKey:       "ts",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		Level:         "capitalColor",
		Time:          "ISO8601",
		Duration:      "seconds",
		Caller:        "short",
		Encoding:      "console",
	}
	defaultRotateCfg = RotateCfg{
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}

	logger = &Logger{
		rotate:  defaultRotateCfg,
		encoder: defaultEncoderCfg,
		Level:   zapcore.InfoLevel,
	}

	infoLevelEnabler = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel
	})

	errorLevelEnabler = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	accessLevelEnabler = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel
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

	Level zapcore.Level
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
		zapcore.NewCore(l.getEncoder(l.encoder.LevelKey), l.getLogWriter(l.infoLogPath), infoLevelEnabler),
		zapcore.NewCore(l.getEncoder(l.encoder.LevelKey), l.getLogWriter(l.errorLogPath), errorLevelEnabler),
		zapcore.NewCore(l.getEncoder(""), l.getLogWriter(l.accessLogPath), accessLevelEnabler),
	), zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(errorLevelEnabler))
	l.sugaredLogger = zapLogger.Sugar()
	l.logger = zapLogger
}

func (l *Logger) getEncoder(levelKey string) zapcore.Encoder {

	var (
		timeEncoder     = new(zapcore.TimeEncoder)
		durationEncoder = new(zapcore.DurationEncoder)
		callerEncoder   = new(zapcore.CallerEncoder)
		nameEncoder     = new(zapcore.NameEncoder)
		levelEncoder    = new(zapcore.LevelEncoder)
	)

	_ = timeEncoder.UnmarshalText([]byte(l.encoder.Time))
	_ = durationEncoder.UnmarshalText([]byte(l.encoder.Duration))
	_ = callerEncoder.UnmarshalText([]byte(l.encoder.Caller))
	_ = nameEncoder.UnmarshalText([]byte("full"))
	_ = levelEncoder.UnmarshalText([]byte(l.encoder.Level))

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        l.encoder.TimeKey,
		LevelKey:       levelKey,
		NameKey:        l.encoder.NameKey,
		CallerKey:      l.encoder.CallerKey,
		MessageKey:     l.encoder.MessageKey,
		StacktraceKey:  l.encoder.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    *levelEncoder,
		EncodeTime:     *timeEncoder,
		EncodeDuration: *durationEncoder,
		EncodeCaller:   *callerEncoder,
		EncodeName:     *nameEncoder,
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
		if l.debug {
			return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
		}
		return zapcore.AddSync(lumberJackLogger)
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

	Level int8

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
	logger.Level = filterZapAtomicLevelByViper(cfg.Level)
	logger.Init()
}

func SetRotate(cfg RotateCfg) {
	logger.rotate = cfg
	logger.Init()
}

// OpenSQLLog set the sqlLogOpen true.
func OpenSQLLog() {
	logger.sqlLogOpen = true
}

// Debug print the debug message.
func Debug(info ...interface{}) {
	if !logger.infoLogOff {
		if logger.Level <= zapcore.DebugLevel {
			logger.sugaredLogger.Info(info...)
		}
	}
}

// Debugf print the debug message.
func Debugf(template string, args ...interface{}) {
	if !logger.infoLogOff && logger.Level <= zapcore.DebugLevel {
		logger.sugaredLogger.Infof(template, args...)
	}
}

// Info print the info message.
func Info(info ...interface{}) {
	if !logger.infoLogOff && logger.Level <= zapcore.InfoLevel {
		logger.sugaredLogger.Info(info...)
	}
}

// Info print the info message.
func Infof(template string, args ...interface{}) {
	if !logger.infoLogOff && logger.Level <= zapcore.InfoLevel {
		logger.sugaredLogger.Infof(template, args...)
	}
}

// Warn print the warning message.
func Warn(info ...interface{}) {
	if !logger.infoLogOff && logger.Level <= zapcore.WarnLevel {
		logger.sugaredLogger.Warn(info...)
	}
}

// Warnf print the warning message.
func Warnf(template string, args ...interface{}) {
	if !logger.infoLogOff && logger.Level <= zapcore.WarnLevel {
		logger.sugaredLogger.Warnf(template, args...)
	}
}

// Error print the error message.
func Error(err ...interface{}) {
	if !logger.errorLogOff && logger.Level <= zapcore.ErrorLevel {
		logger.sugaredLogger.Error(err...)
	}
}

// Errorf print the error message.
func Errorf(template string, args ...interface{}) {
	if !logger.errorLogOff && logger.Level <= zapcore.ErrorLevel {
		logger.sugaredLogger.Errorf(template, args...)
	}
}

// Fatal print the fatal message.
func Fatal(info ...interface{}) {
	if !logger.errorLogOff && logger.Level <= zapcore.ErrorLevel {
		logger.sugaredLogger.Fatal(info...)
	}
}

// Fatalf print the fatal message.
func Fatalf(template string, args ...interface{}) {
	if !logger.errorLogOff && logger.Level <= zapcore.ErrorLevel {
		logger.sugaredLogger.Fatalf(template, args...)
	}
}

// Fatal print the panic message.
func Panic(info ...interface{}) {
	logger.sugaredLogger.Panic(info...)
}

// Panicf print the panic message.
func Panicf(template string, args ...interface{}) {
	logger.sugaredLogger.Panicf(template, args...)
}

// Access print the access message.
func Access(ctx *context.Context) {
	if !logger.accessLogOff && logger.Level <= zapcore.InfoLevel {
		temp := "[GoAdmin] %s %s %s"
		if logger.accessAssetsLogOff {
			if filepath.Ext(ctx.Path()) == "" {
				logger.sugaredLogger.Warnf(temp,
					ansi.Color(" "+strconv.Itoa(ctx.Response.StatusCode)+" ", "white:blue"),
					ansi.Color(" "+string(ctx.Method())+"   ", "white:blue+h"),
					ctx.Path())
			}
		} else {
			logger.sugaredLogger.Warnf(temp,
				ansi.Color(" "+strconv.Itoa(ctx.Response.StatusCode)+" ", "white:blue"),
				ansi.Color(" "+string(ctx.Method())+"   ", "white:blue+h"),
				ctx.Path())
		}
	}
}

// LogSQL print the sql info message.
func LogSQL(statement string, args []interface{}) {
	if !logger.infoLogOff && logger.sqlLogOpen && statement != "" {
		if logger.Level <= zapcore.InfoLevel {
			logger.sugaredLogger.With("statement", statement, "args", args).Info("[GoAdmin]")
		}
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

func filterZapAtomicLevelByViper(level int8) zapcore.Level {
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
