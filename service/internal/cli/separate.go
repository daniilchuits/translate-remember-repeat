package cli

import (
	"context"
	"fmt"
)

func (cli *cliStruct) Separate(ctx context.Context) error {

	if err := ctxCheck(ctx); err != nil {
		return err
	}
	fmt.Print(cli.separators)
	return nil
}
