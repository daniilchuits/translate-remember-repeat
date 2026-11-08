package usecases

import (
	"context"
	"errors"
	"log"
	"strings"
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

type redis interface {
	Save(ctx context.Context, sentence string, resp *domain.ResponseOfDeepL) error
	Read(ctx context.Context, sentence string) (*domain.ResponseOfDeepL, error)
}

type tester interface {
	Test(ctx context.Context) error
}

type translateUsecase struct {
	apiSender
	cliWriter
	repo
	redis
	tester
}

func NewTranslateUsecase(
	apiSender apiSender,
	cliWriter cliWriter,
	repo repo,
	redis redis,
	tester tester,
) *translateUsecase {
	return &translateUsecase{
		apiSender: apiSender,
		cliWriter: cliWriter,
		repo:      repo,
		redis:     redis,
		tester:    tester,
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

		if sentence == domain.Test.String() {
			if err = t.Test(ctx); err != nil {
				log.Panic(err)
			}
			continue
		}

		sentence = strings.ToLower(sentence)

		resp, err := t.Read(ctx, sentence)
		if err == nil {
			t.LogResponse(sentence, resp)
			continue
		}

		resp, err = t.SelectFromTranslationsPSQL(ctx, sentence)
		if err == nil {
			t.LogResponse(sentence, resp)
			if err = t.Save(ctx, sentence, resp); err != nil {
				log.Println(err)
			}
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
		}
	}
}
