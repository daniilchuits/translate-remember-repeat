package cli

import (
	"context"
	"fmt"
)

func (cli *cliStruct) TestAnnouncement(ctx context.Context) error {
	if err := ctxCheck(ctx); err != nil {
		return err
	}
	fmt.Print(cli.testAnnouncement)
	return nil
}
