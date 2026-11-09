package cli

import (
	"context"
	"fmt"
)

func (cli *cliStruct) CorrectAnswer(ctx context.Context) error {

	if err := ctxCheck(ctx); err != nil {
		return err
	}
	fmt.Print(cli.correct)
	return nil
}
