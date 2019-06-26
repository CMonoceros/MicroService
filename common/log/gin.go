package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

func GetGinDefaultWriter() io.Writer {
	return io.MultiWriter(
		os.Stdout,
		fileHandler.fws[_requestInfoIdx],
	)
}

func GetGinDefaultErrorWriter() io.Writer {
	return io.MultiWriter(
		os.Stderr,
		fileHandler.fws[_requestErrorIdx],
	)
}

func GetGinFormatter() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		var statusColor, methodColor, resetColor string
		if param.IsOutputColor() {
			statusColor = param.StatusCodeColor()
			methodColor = param.MethodColor()
			resetColor = param.ResetColor()
		}

		if param.Latency > time.Minute {
			// Truncate in a golang < 1.8 safe way
			param.Latency = param.Latency - param.Latency%time.Second
		}
		return fmt.Sprintf("[%v] [GIN] [%s%3d%s] [%13v] [%15s] [%s%7s%s] %s\n%s",
			param.TimeStamp.Format("2006/01/02 15:04:05.999"),
			statusColor, param.StatusCode, resetColor,
			param.Latency,
			param.ClientIP,
			methodColor, param.Method, resetColor,
			param.Path,
			param.ErrorMessage,
		)
	})
}
