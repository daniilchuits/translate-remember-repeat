package usecases

import (
	"context"
	"errors"
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

type logger interface {
	Log(ctx context.Context, msg string) error
	Panic(msg string)
}

type translateUsecase struct {
	apiSender
	cliWriter
	repo
	redis
	tester
	logger
}

func NewTranslateUsecase(
	apiSender apiSender,
	cliWriter cliWriter,
	repo repo,
	redis redis,
	tester tester,
	logger logger,
) *translateUsecase {
	return &translateUsecase{
		apiSender: apiSender,
		cliWriter: cliWriter,
		repo:      repo,
		redis:     redis,
		tester:    tester,
		logger:    logger,
	}
}

func (t *translateUsecase) Exec(ctx context.Context) {

	t.WriteEnteringMessage()

	for {
		sentence, err := t.WriteCasualMessage(ctx)
		if err != nil {

			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				return
			}

			if errors.Is(err, domain.ReadingCLI) {
				t.Panic(err.Error())
			}
			if err = t.Log(ctx, err.Error()); err != nil {
				t.Panic(err.Error())
			}
			continue
		}

		if sentence == domain.Test.String() {
			if err = t.Test(ctx); err != nil {
				if err = t.Log(ctx, err.Error()); err != nil {
					t.Panic(err.Error())
				}
			}
			continue
		}

		sentence = strings.ToLower(sentence)

		resp, err := t.Read(ctx, sentence)
		if err == nil {
			t.LogResponse(sentence, resp)
			continue
		}
		if !errors.Is(err, domain.NoNote) {
			if err = t.Log(ctx, err.Error()); err != nil {
				t.Panic(err.Error())
			}
		}

		resp, err = t.SelectFromTranslationsPSQL(ctx, sentence)
		if err == nil {
			t.LogResponse(sentence, resp)
			if err = t.Save(ctx, sentence, resp); err != nil {
				if err = t.Log(ctx, err.Error()); err != nil {
					t.Panic(err.Error())
				}
			}
			continue
		}
		if !errors.Is(err, domain.NoRowsFound) {
			if err = t.Log(ctx, err.Error()); err != nil {
				t.Panic(err.Error())
			}
		}

		resp, err = t.SendRequest(ctx, []string{sentence})
		if err != nil {
			if err = t.Log(ctx, err.Error()); err != nil {
				t.Panic(err.Error())
			}
			continue
		}

		t.LogResponse(sentence, resp)
		if err = t.InsertTranslationsPSQL(ctx, sentence, resp); err != nil {
			if err = t.Log(ctx, err.Error()); err != nil {
				t.Panic(err.Error())
			}
		}
	}
}
