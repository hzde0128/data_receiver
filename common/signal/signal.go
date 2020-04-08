package signal

import (
	. "github.com/hzde0128/data_receiver/common/logger"

	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
)

func Trap(cleanup func()) {
	c := make(chan os.Signal, 1)

	signals := []os.Signal{os.Interrupt, syscall.SIGTERM}
	signal.Notify(c, signals...)

	go func() {
		interruptCount := uint32(0) // 记录接收到信号的次数
		for sig := range c {
			go func(sig os.Signal) {
				Log.Info("Received signal: '%v'", sig)
				switch sig {
				case os.Interrupt, syscall.SIGTERM:
					// 接收停止信号小于3次的，第一次清理工作然后退出进程，后面的信号不处理
					// 接收信号达到3次，直接退出，用于用户强制退出的场景
					if atomic.LoadUint32(&interruptCount) < 3 {
						atomic.AddUint32(&interruptCount, 1)
						if atomic.LoadUint32(&interruptCount) == 1 {
							cleanup()
							os.Exit(0)
						} else {
							return
						}
					} else {
						Log.Info("Force stop, interrupting cleanup")
						os.Exit(128 + int(sig.(syscall.Signal)))
					}
				}
			}(sig)
		}
	}()
}
