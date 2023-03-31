package Middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"goFrame/app/Tools"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		beginTime := time.Now().UnixNano()
		endTime := time.Now().UnixNano()
		duration := endTime - beginTime

		s := "%s %s \"%s %s\" " +
			"%s %d %d %dµs " +
			"\"%s\""

		layout := "2006-01-02 15:04:05"
		timeNow := time.Now().Format(layout)

		Tools.Log.AccessLog.Infof(s,
			Tools.GetRealIp(c),
			timeNow,
			c.Request.Method,
			c.Request.RequestURI,
			c.Request.Proto,
			bodyWriter.Status(),
			bodyWriter.body.Len(),
			duration/1000,
			c.Request.Header.Get("User-Agent"),
		)

		c.Next()

		defer Tools.Log.AccessLog.Sync()
	}
}
