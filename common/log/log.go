package log

import (
	"context"
	"github.com/bilibili/kratos/pkg/log"
)

func Init(conf *log.Config) {
	log.Init(conf)
}

// Info logs a message at the info log level.
func Info(format string, args ...interface{}) {
	log.Info(format, args...)
}

// Warn logs a message at the warning log level.
func Warn(format string, args ...interface{}) {
	log.Warn(format, args...)
}

// Error logs a message at the error log level.
func Error(format string, args ...interface{}) {
	log.Error(format, args...)
}

// Infoc logs a message at the info log level.
func Infoc(ctx context.Context, format string, args ...interface{}) {
	log.Infoc(ctx, format, args...)
}

// Errorc logs a message at the error log level.
func Errorc(ctx context.Context, format string, args ...interface{}) {
	log.Errorc(ctx, format, args...)
}

// Warnc logs a message at the warning log level.
func Warnc(ctx context.Context, format string, args ...interface{}) {
	log.Warnc(ctx, format, args...)
}

// Infov logs a message at the info log level.
func Infov(ctx context.Context, args ...log.D) {
	log.Infov(ctx, args...)
}

// Warnv logs a message at the warning log level.
func Warnv(ctx context.Context, args ...log.D) {
	log.Warnv(ctx, args...)
}

// Errorv logs a message at the error log level.
func Errorv(ctx context.Context, args ...log.D) {
	log.Errorv(ctx, args...)
}

// SetFormat only effective on stdout and file handler
// %T time format at "15:04:05.999" on stdout handler, "15:04:05 MST" on file handler
// %t time format at "15:04:05" on stdout handler, "15:04" on file on file handler
// %D data format at "2006/01/02"
// %d data format at "01/02"
// %L log level e.g. INFO WARN ERROR
// %M log message and additional fields: key=value this is log message
// NOTE below pattern not support on file handler
// %f function name and line number e.g. model.Get:121
// %i instance id
// %e deploy env e.g. dev uat fat prod
// %z zone
// %S full file name and line number: /a/b/c/d.go:23
// %s final file name element and line number: d.go:23
func SetFormat(format string) {
	log.SetFormat(format)
}

// Infow logs a message with some additional context. The variadic key-value pairs are treated as they are in With.
func Infow(ctx context.Context, args ...interface{}) {
	log.Infow(ctx, args...)
}

// Warnw logs a message with some additional context. The variadic key-value pairs are treated as they are in With.
func Warnw(ctx context.Context, args ...interface{}) {
	log.Warnw(ctx, args...)
}

// Errorw logs a message with some additional context. The variadic key-value pairs are treated as they are in With.
func Errorw(ctx context.Context, args ...interface{}) {
	log.Errorw(ctx, args...)
}

// Close close resource.
func Close() (err error) {
	return log.Close()
}
