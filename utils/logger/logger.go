package logger

import (
	"fmt"
	"futong_server_agent_go/utils/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

// 对zap进行封装
// github：https://github.com/uber-go/zap
var Logger *zap.Logger
var Sugar *zap.SugaredLogger

// 日志编码配置
func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// 设置时间格式
func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// 初始化
func init() {
	initLogDir()
	fmt.Println("init error.log......")
	// 初始化日志
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:  global.Global.Path + "/logs/error.log", // 日志文件路径
		MaxSize:    10,               // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 3,                // 日志文件最多保存多少个备份
		MaxAge:     7,                // days
		Compress:   true,             // 压缩
	})

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(NewEncoderConfig()),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), w),
		zap.InfoLevel, // 日志级别
	)

	Logger = zap.New(core, zap.AddCaller())
	Sugar = Logger.Sugar()
}

func initLogDir() {
	if _, err := os.Stat(global.Global.Path + "/logs"); err != nil {
		if err := os.Mkdir(global.Global.Path + "/logs", os.ModePerm); err != nil {
			fmt.Println("init logs dir err: ", err)
		}
	}
}
