package main

import (
	. "github.com/hzde0128/data_receiver/common/logger"
	"github.com/hzde0128/data_receiver/initializer"
	_ "github.com/hzde0128/data_receiver/routers"

	"flag"
	"fmt"

	"github.com/astaxie/beego"
)

const VERSION = "v1.0.0"

func main() {
	flag.Parse()

	if *initializer.FlVersion {
		showVersion()
		return
	}

	LogTemp("DRGW begin to start")
	if err := initializer.Initlize(); err != nil {
		return
	}

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	Log.Info("***** DRGW start, listen on port %d *****", beego.BConfig.Listen.HTTPPort)
	beego.Run()
}

func showVersion() {
	fmt.Printf("DRGW version: %s\n", VERSION)
}
