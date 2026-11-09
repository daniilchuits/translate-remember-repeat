package cli

import (
	"context"
	"fmt"
)

func (cli *cliStruct) CorrectAnswers(ctx context.Context, must int, have int) error {

	if err := ctxCheck(ctx); err != nil {
		return err
	}

	percent := (have * 100) / must

	if percent >= 66 {
		fmt.Printf("Great! The percent of correct sentences is %d\n", percent)
	} else {
		fmt.Printf("LOSE :( You need to train more. The percent of correct sentences is %d\n", percent)
	}
	return nil
}
