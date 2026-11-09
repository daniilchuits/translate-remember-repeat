package cli

import (
	"context"
	"fmt"
)

func (cli *cliStruct) Exit(ctx context.Context) error {

	if err := ctxCheck(ctx); err != nil {
		return err
	}
	fmt.Print(cli.exit)
	return nil
}
