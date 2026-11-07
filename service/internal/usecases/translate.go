package usecases

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"translate-remember-repeat/internal/domain"
)

type apiSender interface {
	SendRequest(ctx context.Context, text []string) (string, error)
}

type cliWriter interface {
	WriteEnteringMessage()
	WriteCasualMessage(ctx context.Context) (string, error)
}

type translateUsecase struct {
	apiSender
	cliWriter
}

func NewTranslateUsecase(
	apiSender apiSender,
	cliWriter cliWriter,
) *translateUsecase {
	return &translateUsecase{
		apiSender: apiSender,
		cliWriter: cliWriter,
	}
}

func (t *translateUsecase) Exec(ctx context.Context) {

	fmt.Println("CnnStr:", os.Getenv("CNN_STR"))
	t.WriteEnteringMessage()

	for {
		str, err := t.WriteCasualMessage(ctx)
		if err != nil {
			if errors.Is(err, domain.ReadingCLI) {
				panic(err)
			}
			log.Println(err)
			continue
		}

		translation, err := t.SendRequest(ctx, []string{str})
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println("Translation:", translation)
	}

}
