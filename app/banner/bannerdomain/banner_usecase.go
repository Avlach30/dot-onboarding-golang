package bannerdomain

import (
	"context"
	"github.com/codespace-id/codespace-x/app/banner/bannerdto"
)

type Usecase interface {
	Get(ctx context.Context, page, perPage int) (res []bannerdto.BannerResponse, err error)
}
