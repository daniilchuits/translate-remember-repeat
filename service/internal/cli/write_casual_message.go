package cli

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"translate-remember-repeat/internal/domain"
)

type scanResult struct {
	message string
	err     error
}

func (c *cliStruct) WriteCasualMessage(ctx context.Context) (string, error) {
	fmt.Print(c.casualMessage)

	res := make(chan scanResult, 1)

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
		if strings.ToLower(r.message) == "test" {
			return domain.Test.String(), nil
		}
		return r.message, nil
	}
}
