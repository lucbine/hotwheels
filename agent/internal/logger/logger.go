/*
@Time : 2020/7/13 3:38 下午
@Author : lucbine
@File : logger.go
*/
package logger

import (
	"hotwheels/agent/internal/config"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerObj struct {
	*zap.Logger
}

var Logger = &LoggerObj{}

type LoggerConfig struct {
	AddStacktrace         bool               `mapstructure:"add_stacktrace"`
	AddStacktraceMinLevel zapcore.Level      `mapstructure:"add_stacktrace_min_level"`
	AddCaller             bool               `mapstructure:"add_caller"`
	Development           bool               `mapstructure:"development"`
	AddCallerSkip         bool               `mapstructure:"add_caller_skip"`
	AddCallerSkipVal      int                `mapstructure:"add_caller_skip_val"`
	WriterSyncConfigs     []WriterSyncConfig `mapstructure:"loggers"`
}

type WriterSyncConfig struct {
	OutputPath string        `mapstructure:"output_path"`
	MaxSize    int           `mapstructure:"max_size"`
	MaxBackups int           `mapstructure:"max_backups"`
	MaxAge     int           `mapstructure:"max_age"`
	LocalTime  bool          `mapstructure:"localtime"`
	Compress   bool          `mapstructure:"compress"`
	TimeFormat string        `mapstructure:"time_format"`
	MinLevel   zapcore.Level `mapstructure:"min_level"`
	MaxLevel   zapcore.Level `mapstructure:"max_level"`
}

//日志配置初始化
func InitLog() error {
	var loggerConfig LoggerConfig

	if err := config.Unmarshal("logger", &loggerConfig); err != nil {
		return err
	}

	zapCores := make([]zapcore.Core, 0)
	for _, writerSyncerConfig := range loggerConfig.WriterSyncConfigs {
		newWriter := zapcore.AddSync(&lumberjack.Logger{
			Filename:   writerSyncerConfig.OutputPath,
			MaxSize:    writerSyncerConfig.MaxSize,
			MaxBackups: writerSyncerConfig.MaxBackups,
			MaxAge:     writerSyncerConfig.MaxAge,
			LocalTime:  writerSyncerConfig.LocalTime,
			Compress:   writerSyncerConfig.Compress,
		})
		cfg := zap.NewProductionEncoderConfig()
		cfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.000000"))
		}
		newCore := func(minLevel, maxLevel zapcore.Level) zapcore.Core {
			return zapcore.NewCore(
				zapcore.NewJSONEncoder(cfg),
				newWriter,
				zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
					return lvl >= minLevel && lvl <= maxLevel
				}),
			)
		}(writerSyncerConfig.MinLevel, writerSyncerConfig.MaxLevel)
		zapCores = append(zapCores, newCore)
	}

	core := zapcore.NewTee(zapCores...)
	opts := make([]zap.Option, 0)
	if loggerConfig.AddStacktrace {
		opts = append(opts, zap.AddStacktrace(loggerConfig.AddStacktraceMinLevel))
	}
	if loggerConfig.AddCaller {
		opts = append(opts, zap.AddCaller())
	}
	if loggerConfig.Development {
		opts = append(opts, zap.Development())
	}
	if loggerConfig.AddCallerSkip {
		opts = append(opts, zap.AddCallerSkip(loggerConfig.AddCallerSkipVal))
	}
	Logger.Logger = zap.New(core, opts...)
	return nil
}

// runtime.Caller ORZ
func (c *LoggerObj) Info(msg string, fields ...zap.Field) {
	c.Logger.Info(msg, fields...)
}

func (c *LoggerObj) Error(msg string, fields ...zap.Field) {
	c.Logger.Error(msg, fields...)
}

func (c *LoggerObj) Warn(msg string, fields ...zap.Field) {
	c.Logger.Warn(msg, fields...)
}

func (c *LoggerObj) Debug(msg string, fields ...zap.Field) {
	c.Logger.Debug(msg, fields...)
}
