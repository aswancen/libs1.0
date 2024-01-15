package xhttp

import (
	"fmt"
	"time"

	"github.com/aswancen/libs1.0/pkg/xlogger"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	*resty.Client
}

func New(debug bool, attempt int32, timeout time.Duration, waitTime time.Duration) *Client {
	return &Client{
		Client: resty.New().
			EnableTrace().
			SetDebug(debug).
			SetRetryCount(int(attempt)).
			SetLogger(&Logger{logger: xlogger.NewHelper()}).
			SetTimeout(timeout).
			SetRetryWaitTime(waitTime).
			SetHeader("Content-Type", "application/json").
			SetHeader("User-Agent", "Device/2.0 (Linux; Android 10; Redmi K30 Pro)"),
	}
}

type Logger struct {
	logger *log.Helper
}

func (this *Logger) Errorf(_ string, v ...interface{}) {
	this.logger.Error(fmt.Sprint(v[0]))
}
func (this *Logger) Warnf(_ string, v ...interface{}) {
	this.logger.Warn(fmt.Sprint(v[0]))
}
func (this *Logger) Debugf(_ string, v ...interface{}) {
	this.logger.Debug(fmt.Sprint(v[0]))
}
