package log

import "go.uber.org/zap"

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger() *Logger {
	// TODO: switch to zap.NewDevelopment() when in development
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return &Logger{logger.Sugar()}
}
