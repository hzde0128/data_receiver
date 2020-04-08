package initializer

import (
	. "github.com/hzde0128/data_receiver/common/logger"
	"github.com/hzde0128/data_receiver/common/signal"
	"github.com/hzde0128/data_receiver/initializer/runconfig"
	"github.com/hzde0128/data_receiver/models"

	"fmt"
	"strings"

	"github.com/astaxie/beego"
)

// 初始化工作
func Initlize() error {
	var err error

	if err = parseParam(); err != nil {
		LogTemp("parseParam error: %s", err)
		return err
	}

	if err = runconfig.Config.Parse(); err != nil {
		LogTemp("runconfig.Config.Parse error: %s", err)
		return err
	}

	if err = initLog(); err != nil {
		LogTemp("initLog error: %s", err)
		return err
	}

	if err = initRealTimeModel(); err != nil {
		LogTemp("initRealTimeModel error: %s", err)
		return err
	}

	// 捕获信号
	signal.Trap(cleanup)

	return nil
}

func parseParam() error {
	if *flConfigPath != "" {
		return beego.LoadAppConfig("ini", *flConfigPath)
	}

	return nil
}

func initLog() error {
	logConfig := fmt.Sprintf(`{
		"filename": %q,
		"perm":     "0775"
	}`, runconfig.Config.Log.FilePath)

	return Log.Initialize(runconfig.Config.Log.Level, logConfig)
}

func initRealTimeModel() error {
	return models.Initialize(strings.Split(runconfig.Config.Kafka.Zookeepers, ","))
}

// *******************************************************
// 清理工作
func cleanup() {
	Log.Info("==========cleanup")
}
