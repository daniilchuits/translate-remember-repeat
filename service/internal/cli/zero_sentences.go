package cli

import (
	"context"
	"fmt"
)

func (cli *cliStruct) ZeroSentences(ctx context.Context) error {
	if err := ctxCheck(ctx); err != nil {
		return err
	}
	fmt.Print(cli.zeroSentences)
	return nil
}
