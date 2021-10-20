package main

import (
	"log"

	"github.com/sirupsen/logrus"
	"github.com/tzY15368/lazarus/config"
	"github.com/tzY15368/lazarus/master"
	"github.com/tzY15368/lazarus/worker"
)

func main() {

	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
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
