package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger 日志信息
type Logger struct {
	Config *Config
	logger *zap.SugaredLogger
}

// New 创建新 Logger
func New() *Logger {
	logger := &Logger{
		Config: newConfig(),
	}

	logger.ApplyConfig()
	return logger
}

// ApplyConfig 应用当前Config配置
func (l *Logger) ApplyConfig() {
	conf := l.Config
	cores := []zapcore.Core{}

	var encoder zapcore.Encoder
	if conf.jsonFormat {
		// NewJSONEncoder 输出一个 json 编码器
		encoder = zapcore.NewJSONEncoder(getEncoderConfig())
	} else {
		// NewConsoleEncoder 输出一个非 json 的人类可读的编码器
		encoder = zapcore.NewConsoleEncoder(getEncoderConfig())
	}

	// 控制台输出
	if conf.consoleOut {
		level := conf.logLevel
		if conf.consoleLevel != "" {
			level = conf.consoleLevel
		}

		writeSyncer := zapcore.Lock(os.Stdout)
		core := zapcore.NewCore(encoder, writeSyncer, getLevel(level))
		cores = append(cores, core)
	}

	// 输出到文件
	if conf.fileOut.enable {
		writeSyncer := getFileWriteSyncer(*conf.fileOut)
		core := zapcore.NewCore(encoder, writeSyncer, getLevel(conf.logLevel))
		cores = append(cores, core)
	}

	// 创建新 logger
	logger := zap.New(
		// NewTee 合并多个 cores
		zapcore.NewTee(cores...),
		// AddCallerSkip 指定在调用栈中跳过的调用深度, 因为封装的一层logger，不跳过一层一直显示的是当前日志组件中的行号
		zap.AddCallerSkip(conf.callerSkip),
		// AddStacktrace 打印指定级别以上日志的堆栈信息
		zap.AddStacktrace(getLevel(conf.stacktraceLevel)),
		// AddCaller 将调用函数信息记录到日志中
		zap.AddCaller(),
	)

	// 设置后日志中会打印 logger 名称
	if conf.loggerName != "" {
		logger = logger.Named(conf.loggerName)
	}

	// 将 buffer 中的日志写到文件中
	defer logger.Sync()

	l.logger = logger.Sugar()
}

// Debug
func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.logger.Debugw(msg, keysAndValues...)
}

func (l *Logger) Debugln(args ...interface{}) {
	l.logger.Debugln(args...)
}

// Info
func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l *Logger) Infow(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}

func (l *Logger) Infoln(args ...interface{}) {
	l.logger.Infoln(args...)
}

// Warn
func (l *Logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	l.logger.Warnw(msg, keysAndValues...)
}

func (l *Logger) Warnln(args ...interface{}) {
	l.logger.Warnln(args...)
}

// Error
func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.logger.Errorw(msg, keysAndValues...)
}

func (l *Logger) Errorln(args ...interface{}) {
	l.logger.Errorln(args...)
}

// DPanic
func (l *Logger) DPanic(args ...interface{}) {
	l.logger.DPanic(args...)
}

func (l *Logger) DPanicf(template string, args ...interface{}) {
	l.logger.DPanicf(template, args...)
}

func (l *Logger) DPanicw(msg string, keysAndValues ...interface{}) {
	l.logger.DPanicw(msg, keysAndValues...)
}

func (l *Logger) DPanicln(args ...interface{}) {
	l.logger.DPanicln(args...)
}

// Panic
func (l *Logger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *Logger) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}

func (l *Logger) Panicw(msg string, keysAndValues ...interface{}) {
	l.logger.Panicw(msg, keysAndValues...)
}

func (l *Logger) Panicln(args ...interface{}) {
	l.logger.Panicln(args...)
}

// Fatal
func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

func (l *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.logger.Fatalw(msg, keysAndValues...)
}

func (l *Logger) Fatalln(args ...interface{}) {
	l.logger.Fatalln(args...)
}
