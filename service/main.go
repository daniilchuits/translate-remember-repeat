package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"translate-remember-repeat/internal/cli"
	"translate-remember-repeat/internal/db"
	deeplapi "translate-remember-repeat/internal/infrastructure/deepl_api"
	"translate-remember-repeat/internal/infrastructure/repo"
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

	err := db.InitTranslationsTable(ctx)
	fatalErr(err, "initializing table")
	deeplStruct := deeplapi.New()
	cliStruct := cli.New()
	repoManager, err := repo.OpenDB()
	fatalErr(err, "creating repoManager")
	defer repoManager.Close()

	translateUsecase := usecases.NewTranslateUsecase(
		deeplStruct,
		cliStruct,
		repoManager,
	)

	go translateUsecase.Exec(ctx)

	<-ctx.Done()
}

// cnnStr := os.Getenv("REDIS_CNN_STR")
// client := redis.NewClient(&redis.Options{
// 	Addr: cnnStr,
// })
