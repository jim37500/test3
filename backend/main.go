package main

import (
	"log"
	"os"

	"test3/configuration"
	"test3/database"
	"test3/router"

	"github.com/kardianos/service"
)

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	// 讀取設定檔
	configuration.ReadConfiguration()

	_ = logger.Info("Opening test3 DB...")

	// 連線資料庫
	database.Open()

	_ = logger.Info("Starting test3...")

	// 監聽網頁服務
	router.Run()

	_ = logger.Info("Stoping test3...")
}

func (p *program) Stop(s service.Service) error {
	go func() {
		// 關閉網頁服務
		_ = router.HttpApplication.Shutdown()
	}()
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "test3",
		DisplayName: "test3",
		Description: "test3",
	}

	myService, _ := service.New(&program{}, svcConfig)

	logger, _ = myService.Logger(nil)

	if len(os.Args) > 1 {
		err := service.Control(myService, os.Args[1])
		if err != nil {
			log.Println(err)
		}
		return
	}

	err := myService.Run()
	log.Println(err)
}
