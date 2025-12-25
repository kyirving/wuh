package log

import (
	"basketball/internal/conf"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(config *conf.Config) (*zap.Logger, error) {
	log, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	if config.Manager.Env != "production" {
		return log, nil
	}

	// 生产环境 写日志
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	write := zapcore.AddSync(logFile)
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, write, zapcore.DebugLevel)
	log = zap.New(core, zap.AddCaller())
	return log, nil
}
