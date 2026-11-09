package cli

import (
	"context"
	"fmt"
)

func (cli *cliStruct) WrongAnswer(ctx context.Context, must string) error {

	if err := ctxCheck(ctx); err != nil {
		return err
	}
	fmt.Printf("%s: %s\n", cli.wrongAnswear, must)
	return nil
}
