package logger

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
)

// 记录临时日志，用于日志模块初始化之前
func LogTemp(format string, a ...interface{}) (n int, err error) {
	fileInfo := ":"
	_, file, line, ok := runtime.Caller(1)
	if ok {
		tmp := strings.Split(file, "/")
		fileInfo = fmt.Sprintf("%s:%d", tmp[len(tmp)-1], line)
	}

	// "2006-01-02 15:04:05" golang 格式化时间需写死这个。据说是golang诞生的时间
	prefix := fmt.Sprintf("[%s] [%s] ", time.Now().Format("2006-01-02 15:04:05"), fileInfo)

	return fmt.Printf(prefix+format+"\n", a...)
}

type commonLog struct {
	log *logs.BeeLogger
}

// 日志模块实例，外部调用
var Log = &commonLog{}

func init() {
	Log.log = logs.NewLogger(1)
	Log.log.EnableFuncCallDepth(true)
	Log.log.SetLogFuncCallDepth(3)
}

// 初始化日志实例
func (d *commonLog) Initialize(level int, config string) error {
	err := d.SetLogger(logs.AdapterFile, config)
	if err != nil {
		LogTemp("SetLogger AdapterFile error: %s", err)
		return err
	}

	err = d.SetLogger(logs.AdapterConsole)
	if err != nil {
		LogTemp("SetLogger AdapterConsole error: %s", err)
		return err
	}

	d.SetLevel(level)

	return nil
}

//  configs like:
//  {
//  "filename":"logs/beego.log",
//  "maxLines":10000,
//  "maxsize":1024,
//  "daily":true,
//  "maxDays":15,
//  "rotate":true,
//  "perm":"0600"
//  }
func (d *commonLog) SetLogger(adapterName string, configs ...string) error {
	config := append(configs, "{}")[0]

	return d.log.SetLogger(adapterName, config)
}

func (d *commonLog) SetLevel(l int) {
	d.log.SetLevel(l)
}

func (d *commonLog) Emergency(format string, v ...interface{}) {
	d.log.Emergency(format, v...)
}

func (d *commonLog) Alert(format string, v ...interface{}) {
	d.log.Alert(format, v...)
}

func (d *commonLog) Critical(format string, v ...interface{}) {
	d.log.Critical(format, v...)
}

func (d *commonLog) Error(format string, v ...interface{}) {
	d.log.Error(format, v...)
}

func (d *commonLog) Warn(format string, v ...interface{}) {
	d.log.Warn(format, v...)
}

func (d *commonLog) Info(format string, v ...interface{}) {
	d.log.Info(format, v...)
}

func (d *commonLog) Debug(format string, v ...interface{}) {
	d.log.Debug(format, v...)
}
