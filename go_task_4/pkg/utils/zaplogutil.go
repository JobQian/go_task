package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger
var atomicLevel zap.AtomicLevel

func InitLogger() {

	//可以动态修改日志级别 而不用重新new logger实例
	levelStr := viper.GetString("log.level")
	if err := atomicLevel.UnmarshalText([]byte(levelStr)); err != nil {
		atomicLevel.SetLevel(zap.InfoLevel)
	}

	// 日志格式选择
	var encoder zapcore.Encoder
	switch viper.GetString("log.format") {
	case "console":
		encoder = zapcore.NewConsoleEncoder(encoderConfig())
	default: // json
		encoder = zapcore.NewJSONEncoder(encoderConfig())
	}

	//按小时记录日志
	now := time.Now()
	fileName := fmt.Sprintf("%04d-%02d-%02d-%02d.log",
		now.Year(), now.Month(), now.Day(), now.Hour())
	logDir := viper.GetString("log.logfilepath")
	fullPath := filepath.Join(logDir, fileName)

	// 输出目标
	output := viper.GetString("log.output")
	var ws zapcore.WriteSyncer
	switch output {
	case "stdout":
		ws = zapcore.AddSync(os.Stdout)
	case "stderr":
		ws = zapcore.AddSync(os.Stderr)
	case "file":
		file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic("无法创建日志文件，程序终止")
		}
		ws = zapcore.AddSync(file)
	case "stdout&file":
		// 日志输出到控制台
		consoleSyncer := zapcore.AddSync(os.Stdout)
		// 日志输出到文件
		file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic("无法创建日志文件，程序终止")
		}
		fileSyncer := zapcore.AddSync(file)
		// 合并两个 syncer
		ws = zapcore.NewMultiWriteSyncer(consoleSyncer, fileSyncer)
	default:
		ws = zapcore.AddSync(os.Stdout)
	}

	core := zapcore.NewCore(encoder, ws, atomicLevel)

	opts := []zap.Option{}
	if viper.GetBool("log.options.add_caller") {
		opts = append(opts, zap.AddCaller())
	}
	if viper.GetBool("log.options.development") {
		opts = append(opts, zap.Development())
	}

	Logger = zap.New(core, opts...)
}

// 编码器配置
func encoderConfig() zapcore.EncoderConfig {
	cfg := zap.NewProductionEncoderConfig()
	cfg.TimeKey = viper.GetString("log.encoding.time_key")
	cfg.LevelKey = viper.GetString("log.encoding.level_key")
	cfg.CallerKey = viper.GetString("log.encoding.caller_key")
	cfg.MessageKey = viper.GetString("log.encoding.message_key")
	cfg.StacktraceKey = viper.GetString("log.encoding.stacktrace_key")

	// 时间格式
	switch viper.GetString("log.encoding.time_format") {
	case "epoch":
		cfg.EncodeTime = zapcore.EpochTimeEncoder
	case "iso8601":
		cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	case "custom":
		// 自定义格式
		cfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format(viper.GetString("log.encoding.time_format")))
		}
	}

	cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder // INFO/ERROR 大写显示｜LowercaseLevelEncoder → 小写 info, debug｜CapitalColorLevelEncoder → 带颜色（console 模式）
	cfg.EncodeCaller = zapcore.ShortCallerEncoder      // 短文件名:行号｜FullCallerEncoder → 输出完整路径
	return cfg
}
