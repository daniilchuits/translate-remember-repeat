package usecases

import (
	"context"
	"fmt"
)

type cliTestWriter interface {
	TestAnnouncement()
}

type sentLister interface {
	GetSentences(ctx context.Context) ([]string, error)
}

type testUsecase struct {
	cliTestWriter
	sentLister
}

func NewTestUsecase(
	cliTestWriter cliTestWriter,
	sentLister sentLister,
) *testUsecase {
	return &testUsecase{
		cliTestWriter: cliTestWriter,
		sentLister:    sentLister,
	}
}

func (t *testUsecase) Test(ctx context.Context) error {

	t.TestAnnouncement()

	sentences, err := t.GetSentences(ctx)
	if err != nil {
		return fmt.Errorf("Test execution error: %w", err)
	}
	for i := range sentences {
		fmt.Println(sentences[i])
	}
	return nil
}
