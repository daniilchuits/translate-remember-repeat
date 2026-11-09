package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"
	"translate-remember-repeat/internal/cli"
	"translate-remember-repeat/internal/db"
	"translate-remember-repeat/internal/ending"
	deeplapi "translate-remember-repeat/internal/infrastructure/deepl_api"
	"translate-remember-repeat/internal/infrastructure/redis"
	"translate-remember-repeat/internal/infrastructure/repo"
	"translate-remember-repeat/internal/test"
	"translate-remember-repeat/internal/usecases"
)

func fatalErr(err error, msg string) {
	if err != nil {
		log.Panicf("Error %s: %s", msg, err.Error())
	}
}

func main() {

	globalCtx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	err := db.InitTranslationsTable(globalCtx)
	fatalErr(err, "initializing table")

	deeplStruct := deeplapi.New()
	cliStruct := cli.New()

	repoManager, err := repo.OpenDB()
	fatalErr(err, "creating repoManager")
	defer repoManager.Close()

	redisClient, err := redis.New()
	fatalErr(err, "connecting redis")

	sentLister := test.NewGetKeysUsecase(
		redisClient,
		repoManager,
	)

	postgresChecker := test.NewPostgresChecker(
		repoManager,
		cliStruct,
	)
	redisChecker := test.NewRedisChecker(
		redisClient,
		cliStruct,
	)

	testUsecase := usecases.NewTestUsecase(
		cliStruct,
		sentLister,
		postgresChecker,
		redisChecker,
	)

	translateUsecase := usecases.NewTranslateUsecase(
		deeplStruct,
		cliStruct,
		repoManager,
		redisClient,
		testUsecase,
		cliStruct,
	)

	go translateUsecase.Exec(globalCtx)

	<-globalCtx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	endingStruct := ending.New(
		cliStruct,
	)

	go endingStruct.Exec(ctx)

	<-ctx.Done()
}
