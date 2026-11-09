package cli

import (
	"context"
	"fmt"
)

func (cli *cliStruct) Exit(ctx context.Context) error {

	err := ctxCheck(ctx)
	if err != nil {
		return fmt.Errorf("Error exiting: %w", err)
	}
	fmt.Print(cli.exitText)
	return nil
}
