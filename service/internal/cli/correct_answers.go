package cli

import (
	"context"
	"fmt"
)

func (cli *cliStruct) CorrectAnswers(ctx context.Context, must, have int) error {

	if err := ctxCheck(ctx); err != nil {
		return err
	}

	percent := (must * 100) / have

	if percent >= 66 {
		fmt.Printf("Great! The percent of correct sentences is %d\n", percent)
	} else {
		fmt.Printf("LOSE :( . You need to train more. The percent of correct sentences is %d\n", percent)
	}
	return nil
}
