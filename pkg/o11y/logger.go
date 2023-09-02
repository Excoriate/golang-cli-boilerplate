package o11y

import (
	"os"

	"github.com/Excoriate/golang-cli-boilerplate/pkg/env"
	"github.com/pterm/pterm"
)

var defaultLogger = pterm.DefaultLogger.
	WithFormatter(pterm.LogFormatterJSON)

type LoggerOptions struct {
	RegisterCallerFunction bool
	WriteToStdout          bool
}

type LoggerArg struct {
	Key   string
	Value interface{}
}

type LoggerInterface interface {
	Trace(msg string, args ...LoggerArg)
	Debug(msg string, args ...LoggerArg)
	Info(msg string, args ...LoggerArg)
	Warn(msg string, args ...LoggerArg)
	Error(msg string, args ...LoggerArg)
	Fatal(msg string, args ...LoggerArg)
	Args(args ...LoggerArg) []pterm.LoggerArgument
}

type LogImpl struct {
	impl *pterm.Logger
}

func (l *LogImpl) Fatal(msg string, args ...LoggerArg) {
	if len(args) == 0 {
		l.impl.Fatal(msg)
		return
	}
	l.impl.Fatal(msg, l.Args(args...))
}

func (l *LogImpl) Error(msg string, args ...LoggerArg) {
	if len(args) == 0 {
		l.impl.Error(msg)
		return
	}
	l.impl.Error(msg, l.Args(args...))
}

func (l *LogImpl) Trace(msg string, args ...LoggerArg) {
	if len(args) == 0 {
		l.impl.Trace(msg)
		return
	}
	l.impl.Trace(msg, l.Args(args...))
}

func (l *LogImpl) Debug(msg string, args ...LoggerArg) {
	if len(args) == 0 {
		l.impl.Debug(msg)
		return
	}
	l.impl.Debug(msg, l.Args(args...))
}

func (l *LogImpl) Info(msg string, args ...LoggerArg) {
	if len(args) == 0 {
		l.impl.Info(msg)
		return
	}
	l.impl.Info(msg, l.Args(args...))
}

func (l *LogImpl) Warn(msg string, args ...LoggerArg) {
	if len(args) == 0 {
		l.impl.Warn(msg)
		return
	}
	l.impl.Warn(msg, l.Args(args...))
}

func (l *LogImpl) Args(args ...LoggerArg) []pterm.LoggerArgument {
	var loggerArgs []pterm.LoggerArgument

	for _, arg := range args {
		ptermArg := pterm.LoggerArgument{
			Key:   arg.Key,
			Value: arg.Value,
		}

		loggerArgs = append(loggerArgs, ptermArg)
	}
	return loggerArgs
}

func NewLogger(options LoggerOptions) LoggerInterface {
	l := &LogImpl{
		impl: defaultLogger,
	}

	if options.RegisterCallerFunction {
		l.impl.WithCaller()
	}

	levelSetFromEnv := env.GetStringOrDefault("LOG_LEVEL", "")

	if options.WriteToStdout {
		l.impl.WithWriter(os.Stdout)
	} else {
		l.impl.WithWriter(os.Stderr)
	}

	if levelSetFromEnv == "" {
		l.impl.WithLevel(pterm.LogLevelInfo)
		return l
	}

	switch levelSetFromEnv {
	case "info":
		l.impl.WithLevel(pterm.LogLevelInfo)
	case "warn":
		l.impl.WithLevel(pterm.LogLevelWarn)
	case "trace":
		l.impl.WithLevel(pterm.LogLevelTrace)
	case "debug":
		l.impl.WithLevel(pterm.LogLevelDebug)
	case "error":
		l.impl.WithLevel(pterm.LogLevelError)
	case "fatal":
		l.impl.WithLevel(pterm.LogLevelFatal)
	}

	return l
}
