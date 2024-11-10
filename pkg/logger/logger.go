package logger

import (
	"context"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

type Logger struct {
	*zap.SugaredLogger
	level zapcore.Level
}

func NewLogger(level string) *Logger {
	// 解析日志级别
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}

	// 创建基础配置
	config := zap.Config{
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Development:       false,
		Encoding:         "json",
		EncoderConfig:    getEncoderConfig(),
		OutputPaths:      []string{"stdout", "logs/app.log"},
		ErrorOutputPaths: []string{"stderr", "logs/error.log"},
	}

	// 创建logger
	logger, err := config.Build(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	if err != nil {
		panic(err)
	}

	return &Logger{
		SugaredLogger: logger.Sugar(),
		level:         zapLevel,
	}
}

func getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// GrpcInterceptor 创建gRPC拦截器用于日志记录
func (l *Logger) GrpcInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		
		// 记录请求
		l.Infow("Received gRPC request",
			"method", info.FullMethod,
			"request", req,
		)

		// 处理请求
		resp, err := handler(ctx, req)

		// 记录响应
		l.Infow("Completed gRPC request",
			"method", info.FullMethod,
			"duration", time.Since(start),
			"error", err,
		)

		return resp, err
	}
}

// WithContext 添加上下文信息到日志
func (l *Logger) WithContext(ctx context.Context) *Logger {
	// 从上下文中提取相关信息
	// 例如：请求ID、用户ID等
	return l
}

// Sync 确保所有日志都被写入
func (l *Logger) Sync() error {
	return l.SugaredLogger.Sync()
} 
