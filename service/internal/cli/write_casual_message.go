package cli

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"translate-remember-repeat/internal/domain"
)

func (c *cliStruct) WriteCasualMessage(ctx context.Context) (string, error) {
	fmt.Print(c.casualMessage)

	type scanResult struct {
		message string
		err     error
	}
	res := make(chan scanResult)

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			res <- scanResult{
				message: scanner.Text(),
			}
			return
		}
		err := scanner.Err()
		if err == nil {
			err = io.EOF
		}
		res <- scanResult{
			err: err,
		}
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case r := <-res:
		if r.err != nil {
			return "", fmt.Errorf("%w: %w", domain.Shutdown, r.err)
		}
		return r.message, nil
	}
}
