package cli

import (
	"context"
	"fmt"
)

func (cli *cliStruct) WrongLen(ctx context.Context, must, have int) error {
	if err := ctxCheck(ctx); err != nil {
		return err
	}
	fmt.Printf("%s Must be %d, has %d\n", cli.lenSentences, must, have)
	return nil
}
