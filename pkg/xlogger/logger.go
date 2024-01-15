package xlogger

import (
	"github.com/go-kratos/kratos/v2/log"
)

func NewHelper(logLevel ...log.Level) *log.Helper {
	var level = log.LevelDebug
	if len(logLevel) != 0 {
		switch logLevel[0] {
		case 1:
			level = log.LevelDebug
		case 2:
			level = log.LevelInfo
		case 3:
			level = log.LevelWarn
		case 4:
			level = log.LevelError
		default:
			level = log.LevelDebug
		}
	}
	return log.NewHelper(
		log.NewFilter(log.With(
			log.DefaultLogger,
			"ts", log.Timestamp("2006-01-02 15:04:05.000000"),
			"caller", log.Caller(5),
		),
			// 等级过滤
			log.FilterLevel(level),
		),
	)
}

func NewCallerHelper(caller int, logLevel ...log.Level) *log.Helper {
	var level = log.LevelDebug
	if len(logLevel) != 0 {
		switch logLevel[0] {
		case 1:
			level = log.LevelDebug
		case 2:
			level = log.LevelInfo
		case 3:
			level = log.LevelWarn
		case 4:
			level = log.LevelError
		default:
			level = log.LevelDebug
		}
	}
	return log.NewHelper(
		log.NewFilter(log.With(
			log.DefaultLogger,
			"ts", log.Timestamp("2006-01-02 15:04:05.000000"),
			"caller", log.Caller(caller),
		),
			// 等级过滤
			log.FilterLevel(level),
		),
	)
}
