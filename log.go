/*
Copyright 2023 HuanShiTianXin and The ZapLog Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package log

import "github.com/fisherzrj/zaplog/logger"

var (
	log    *logger.Logger
	Config *logger.Config
)

func init() {
	log = logger.New()
	Config = log.Config

	log.Config.SetCallerSkip(2)
	log.ApplyConfig()
}

// ApplyConfig 应用当前Config配置
func ApplyConfig() {
	log.ApplyConfig()
}

// PrintSliceOutSlice 打印输出到切片的日志
func PrintSliceOutSlice() {
	log.PrintSliceOutSlice()
}

// Debug
func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	log.Debugf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	log.Debugw(msg, keysAndValues...)
}

func Debugln(args ...interface{}) {
	log.Debugln(args...)
}

// Info
func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(template string, args ...interface{}) {
	log.Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	log.Infow(msg, keysAndValues...)
}

func Infoln(args ...interface{}) {
	log.Infoln(args...)
}

// Warn
func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	log.Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	log.Warnw(msg, keysAndValues...)
}

func Warnln(args ...interface{}) {
	log.Warnln(args...)
}

// Error
func Error(args ...interface{}) {
	log.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	log.Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	log.Errorw(msg, keysAndValues...)
}

func Errorln(args ...interface{}) {
	log.Errorln(args...)
}

// DPanic
func DPanic(args ...interface{}) {
	log.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	log.DPanicf(template, args...)
}

func DPanicw(msg string, keysAndValues ...interface{}) {
	log.DPanicw(msg, keysAndValues...)
}

func DPanicln(args ...interface{}) {
	log.DPanicln(args...)
}

// Panic
func Panic(args ...interface{}) {
	log.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	log.Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	log.Panicw(msg, keysAndValues...)
}

func Panicln(args ...interface{}) {
	log.Panicln(args...)
}

// Fatal
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	log.Fatalf(template, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	log.Fatalw(msg, keysAndValues...)
}

func Fatalln(args ...interface{}) {
	log.Fatalln(args...)
}
