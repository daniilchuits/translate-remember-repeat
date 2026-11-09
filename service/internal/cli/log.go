package cli

import (
	"context"
	"log"
)

func (cli *cliStruct) Log(ctx context.Context, msg string) error {

	if err := ctxCheck(ctx); err != nil {
		return err
	}
	log.Println(msg)
	return nil
}
