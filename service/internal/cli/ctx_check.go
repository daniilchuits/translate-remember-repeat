package cli

import (
	"context"
	"fmt"
	"translate-remember-repeat/internal/domain"
)

func ctxCheck(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("%w: %w", domain.Shutdown, err)
	}
	return nil
}
