package authdomain

import "context"

type OtpRepository interface {
	Create(ctx context.Context, payload OtpEntity) error
	FindByIdentifier(ctx context.Context, identifier string) (OtpEntity, error)
	UpdateByIdentifier(ctx context.Context, identifier string, payload OtpEntity) error
	Upsert(ctx context.Context, payload OtpEntity) error
}
