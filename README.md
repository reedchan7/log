# log

`log` 是一个基于 [zap](https://github.com/uber-go/zap) 封装的一个简单的日志库，开箱即用。


## 示例

```go
Info("Hello, World!")
Warn("Hello, Warning!")
Error("nil pointer dereference", ErrorField(errors.New("NPE")))
Error("some error", ErrorStr(errors.New("what's wrong")))
Infof("Today is %s, so happy!", "Friday")
Debug("This log should not be displayed")
SetDebug()
Debug("This log should be displayed")
SetLevel(ErrorLevel)
Info("This log should not be displayed")
Warn("This log should not be displayed")

logger, _ := New(&Config{Encoding: "console"})
ReplaceLogger(logger)
SetLevel(InfoLevel)
Infof("What's your name? My name is %s", "Chan")
Error("illegal argument error", String("arg", "name"))
```

输出：

```text
{"level":"info","time":"2022-11-11 13:26:43.890","caller":"log/logger_test.go:9","msg":"Hello, World!"}
{"level":"warn","time":"2022-11-11 13:26:43.890","caller":"log/logger_test.go:10","msg":"Hello, Warning!"}
{"level":"error","time":"2022-11-11 13:26:43.890","caller":"log/logger_test.go:11","msg":"nil pointer dereference","error":"NPE","stacktrace":"e.coding.net/coding-artifacts/log.TestLogger\n\t/Users/chenxinyu/Workspace/go/coding/log/logger_test.go:11\ntesting.tRunner\n\t/usr/local/go/src/testing/testing.go:1446"}
{"level":"error","time":"2022-11-11 13:26:43.890","caller":"log/logger_test.go:12","msg":"some error","error":"what's wrong","stacktrace":"e.coding.net/coding-artifacts/log.TestLogger\n\t/Users/chenxinyu/Workspace/go/coding/log/logger_test.go:12\ntesting.tRunner\n\t/usr/local/go/src/testing/testing.go:1446"}
{"level":"info","time":"2022-11-11 13:26:43.890","caller":"log/logger_test.go:13","msg":"Today is Friday, so happy!"}
{"level":"debug","time":"2022-11-11 13:26:43.890","caller":"log/logger_test.go:16","msg":"This log should be displayed"}
2022-11-11 13:26:43.891	INFO	log/logger_test.go:24	What's your name? My name is Chan
2022-11-11 13:38:35.483	ERROR	log/logger_test.go:25	illegal argument error	{"arg": "name"}
e.coding.net/coding-artifacts/log.TestLogger
	/Users/chenxinyu/Workspace/go/coding/log/logger_test.go:25
testing.tRunner
	/usr/local/go/src/testing/testing.go:1446
```
