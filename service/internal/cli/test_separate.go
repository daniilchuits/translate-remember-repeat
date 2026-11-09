package cli

import (
	"context"
	"fmt"
)

func (cli *cliStruct) SeparateTest(ctx context.Context) error {

	if err := ctxCheck(ctx); err != nil {
		return err
	}
	fmt.Print(cli.separatorsTest)
	return nil
}
