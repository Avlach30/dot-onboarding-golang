package domain

import "context"

type ProjectHistoryRepository interface {
	Get(ctx context.Context, projectUUID string, page, perPage int) (res []ProjectHistoryEntity, err error)
}
