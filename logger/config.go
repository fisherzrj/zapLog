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

package logger

import (
	"fmt"
	"io"
)

// Config log配置文件
type Config struct {
	logLevel        string    // 日志记录级别
	stacktraceLevel string    // 记录堆栈的级别
	loggerName      string    // logger名称
	callerSkip      int       // CallerSkip次数
	jsonFormat      bool      // 输出json格式
	consoleOut      bool      // 是否输出到console
	consoleLevel    string    // 控制台日志记录级别
	fileOut         *fileOut  // 日志文件输出
	sliceOut        bool      // 是否输出到切片
	sliceOutSlice   *[]string // 输出的切片
	output          io.Writer // 将日志写入指定输出
}

// fileOut 日志文件输出配置
type fileOut struct {
	enable     bool   // 是否将日志输出到文件
	filename   string // 日志文件的位置
	maxSize    int    // 在进行切割之前，日志文件的最大大小（以MB为单位）
	maxBackups int    // 保留旧文件的最大个数
	maxAge     int    // 保留旧文件的最大天数
	compress   bool   // 是否压缩/归档旧文件
}

// newConfig 创建新 Config
func newConfig() *Config {
	return &Config{
		logLevel:        "info",
		stacktraceLevel: "panic",
		loggerName:      "",
		callerSkip:      1,
		jsonFormat:      false,
		consoleOut:      false,
		consoleLevel:    "",
		sliceOut:        false,
		sliceOutSlice:   &[]string{},
		output:          nil,
		fileOut: &fileOut{
			enable:     false,
			filename:   "",
			maxSize:    10,
			maxBackups: 5,
			maxAge:     30,
			compress:   false,
		},
	}
}

// SetLogLevel 设置日志记录级别
func (c *Config) SetLogLevel(level string) {
	c.logLevel = level
}

// SetStacktraceLevel 设置堆栈跟踪的日志级别
func (c *Config) SetStacktraceLevel(level string) {
	c.stacktraceLevel = level
}

// SetLoggerName 设置logger名称
func (c *Config) SetLoggerName(loggerName string) {
	c.loggerName = loggerName
}

// SetCallerSkip 设置callerSkip次数
func (c *Config) SetCallerSkip(callerSkip int) {
	c.callerSkip = callerSkip
}

// EnableJSONFormat 开启JSON格式化输出
func (c *Config) EnableJSONFormat() {
	c.jsonFormat = true
}

// DisableJSONFormat 关闭JSON格式化输出
func (c *Config) DisableJSONFormat() {
	c.jsonFormat = false
}

// EnableConsoleOut 开启Console输出
func (c *Config) EnableConsoleOut() {
	c.consoleOut = true
}

// DisableConsoleOut 关闭Console输出
func (c *Config) DisableConsoleOut() {
	c.consoleOut = false
}

// SetConsoleLevel 设置控制台日志级别
func (c *Config) SetConsoleLevel(level string) {
	c.consoleLevel = level
}

// SetConsoleLevel 恢复默认控制台日志级别，和 logLevel 保持一致
func (c *Config) ClearConsoleLevel() {
	c.consoleLevel = ""
}

// EnableSliceOut 开启输出到切片
func (c *Config) EnableSliceOut() {
	c.sliceOut = true
}

// DisableSliceOuts 关闭输出到切片
func (c *Config) DisableSliceOut() {
	c.sliceOut = false
}

// SetSliceOutSlice 设置输出日志的切片
func (c *Config) SetSliceOutSlice(s *[]string) {
	c.sliceOutSlice = s
}

// GetSliceOutSlice 获取输出日志的切片
func (c *Config) GetSliceOutSlice() *[]string {
	return c.sliceOutSlice
}

// ClearSliceOutSlice 清除切片中的日志
func (c *Config) ClearSliceOutSlice() {
	*c.sliceOutSlice = []string{}
}

// PrintSliceOutSlice 获取输出日志的切片
func (c *Config) PrintSliceOutSlice() {
	for _, s := range *c.sliceOutSlice {
		fmt.Print(s)
	}
	c.ClearSliceOutSlice()
}

// EnableFileOut 开启日志文件输出
func (c *Config) EnableFileOut(filename string) {
	c.fileOut.enable = true
	c.fileOut.filename = filename
}

// DisableFileOut 关闭日志文件输出
func (c *Config) DisableFileOut(filename string) {
	c.fileOut.enable = false
}

// SetFileOutMaxSize 设置在进行切割之前，日志文件的最大大小（以MB为单位）
func (c *Config) SetFileOutMaxSize(maxSize int) {
	c.fileOut.maxSize = maxSize
}

// SetFileOutMaxBackups 设置保留旧文件的最大个数
func (c *Config) SetFileOutMaxBackups(maxBackups int) {
	c.fileOut.maxBackups = maxBackups
}

// SetFileOutMaxAge 设置保留旧文件的最大天数
func (c *Config) SetFileOutMaxAge(maxAge int) {
	c.fileOut.maxAge = maxAge
}

// EnableFileOutCompress 开启压缩/归档旧文件
func (c *Config) EnableFileOutCompress() {
	c.fileOut.compress = true
}

// DisableFileOutCompress 关闭压缩/归档旧文件
func (c *Config) DisableFileOutCompress() {
	c.fileOut.compress = false
}

// SetOutput 设置日志的输出目标
func (c *Config) SetOutput(w io.Writer) {
	c.output = w
}

// ClearOutput 清除设置的日志输出目标
func (c *Config) ClearOutput() {
	c.output = nil
}
