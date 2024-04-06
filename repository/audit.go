package repository

import (
	"context"
)

func AuditOperation(ctx context.Context, id string, table string, operationType string) error {
	return implementation.AuditOperation(ctx, id, table, operationType)
}
