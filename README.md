# 使用默认日志记录器
```go
package main

import (
	log "github.com/fisherzrj/zaplog"
)

func main() {
	// 输出到日志文件
	log.Config.EnableFileOut("./zaplog.log")

	// 开启控制台输出
	log.Config.EnableConsoleOut()

	// 应用当前配置
	log.ApplyConfig()

    log.Info("新日志")
    log.Infof("格式化的日志: %s", "zaplog")
}
```

# 创建新日志记录器
```go
package main

import (
	"github.com/fisherzrj/zaplog/logger"
)

func main() {
    log := logger.New()

    // 设置新logger名称，多个logger时可设置用来区分
    log.Config.SetLoggerName("newzaplog")

	// 输出到日志文件
	log.Config.EnableFileOut("./zaplog.log")

	// 开启控制台输出
	log.Config.EnableConsoleOut()

	// 应用当前配置
	log.ApplyConfig()

    log.Info("新日志")
    log.Infof("格式化的日志: %s", "zaplog")
}
```