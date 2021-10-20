package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/tzY15368/proxypanel/config"
	"github.com/tzY15368/proxypanel/master"
	"github.com/tzY15368/proxypanel/worker"
)

func main() {

	cfg, err := config.InitConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if cfg.Master.Enabled {
		master.StartMaster(cfg.Master)
		logrus.Infof("started master %v\n", cfg.Master)
	}
	if cfg.Worker.Enabled {
		worker.StartWorker(cfg.Worker)
		logrus.Infof("started worker %v\n", cfg.Worker)
	}

	select {}
}