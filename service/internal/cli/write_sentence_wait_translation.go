package cli

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"translate-remember-repeat/internal/domain"
)

func (cli *cliStruct) WriteSentanceWaitTranslation(
	ctx context.Context,
	sentence string,
) (string, error) {

	res := make(chan scanResult, 1)

	fmt.Printf("%s: %s\n", cli.writeTranslation, sentence)

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			res <- scanResult{
				message: scanner.Text(),
			}
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
