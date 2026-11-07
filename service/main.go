package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"translate-remember-repeat/internal/cli"
	deeplapi "translate-remember-repeat/internal/infrastructure/deepl_api"
	"translate-remember-repeat/internal/usecases"
)

func fatalErr(err error, msg string) {
	if err != nil {
		log.Panicf("Error %s: %s", msg, err.Error())
	}
}

func main() {

	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	deeplStruct := deeplapi.New(os.Getenv("DEEPL_AUTH_KEY"))
	cliStruct := cli.New()

	translateUsecase := usecases.NewTranslateUsecase(
		deeplStruct,
		cliStruct,
	)

	go translateUsecase.Exec(ctx)

	<-ctx.Done()
}

// cnnStr := os.Getenv("REDIS_CNN_STR")
// client := redis.NewClient(&redis.Options{
// 	Addr: cnnStr,
// })
