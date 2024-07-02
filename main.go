package main

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

type config struct {
	ENW string
}

func main() {
	log, err := initLogger()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()

	if err := run(log); err != nil {
		log.Errorf("error")
		os.Exit(1)
	}
}

func run(log *zap.SugaredLogger) error {
	log.Infow("starting")
	var cfg config
	err := envconfig.Process("APP", &cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println("Текущий стенд =", cfg.ENW)
	return nil
}

func initLogger() (*zap.SugaredLogger, error) {
	logger, _ := zap.NewProduction()
	return logger.Sugar(), nil
}

/*
добавить на гит в дз

import (
	"fmt"
)

func main() {
	test := 21
	fmt.Println(test)
}
*/

//log.Errorw log.Infow - выводить то что сейчас выводится через fmt.Println
