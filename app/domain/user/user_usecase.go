package userdomain

import (
	"context"
	"github.com/codespace-id/codespace-x/app/dto"
)

type Usecase interface {
	Create(ctx context.Context, dto dto.RegisterRequest) error
}
