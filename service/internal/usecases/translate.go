package usecases

import (
	"context"
	"errors"
	"log"
	"translate-remember-repeat/internal/domain"
)

type apiSender interface {
	SendRequest(ctx context.Context, text []string) (*domain.ResponseOfDeepL, error)
}

type cliWriter interface {
	WriteEnteringMessage()
	WriteCasualMessage(ctx context.Context) (string, error)
	LogResponse(sentence string, resp *domain.ResponseOfDeepL)
}

type repo interface {
	InsertTranslationsPSQL(ctx context.Context, sentence string, resp *domain.ResponseOfDeepL) error
	SelectFromTranslationsPSQL(ctx context.Context, sentence string) (*domain.ResponseOfDeepL, error)
}

type translateUsecase struct {
	apiSender
	cliWriter
	repo
}

func NewTranslateUsecase(
	apiSender apiSender,
	cliWriter cliWriter,
	repo repo,
) *translateUsecase {
	return &translateUsecase{
		apiSender: apiSender,
		cliWriter: cliWriter,
		repo:      repo,
	}
}

func (t *translateUsecase) Exec(ctx context.Context) {

	t.WriteEnteringMessage()

	for {
		sentence, err := t.WriteCasualMessage(ctx)
		if err != nil {
			if errors.Is(err, domain.ReadingCLI) {
				panic(err)
			}
			log.Println(err)
			continue
		}

		resp, err := t.SelectFromTranslationsPSQL(ctx, sentence)
		if err == nil {
			log.Println("Used psql")
			t.LogResponse(sentence, resp)
			continue
		}
		if !errors.Is(err, domain.NoRowsFound) {
			log.Println(err)
		}

		resp, err = t.SendRequest(ctx, []string{sentence})
		if err != nil {
			log.Println(err)
			continue
		}

		t.LogResponse(sentence, resp)
		if err = t.InsertTranslationsPSQL(ctx, sentence, resp); err != nil {
			log.Panicln(err)
		} else {
			log.Println("Inserted")
		}
	}

}
