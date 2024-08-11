package usecase

import (
	"context"

	bannerdomain "github.com/codespace-id/codespace-x/app/domain/banner"
)

type bannerUsecase struct {
	bannerRepo bannerdomain.Repository
}

func NewBannerUsecase(bannerRepo bannerdomain.Repository) bannerdomain.Usecase {
	return &bannerUsecase{
		bannerRepo: bannerRepo,
	}
}

// Get implements bannerdomain.Usecase.
func (b *bannerUsecase) Get(ctx context.Context, page int, perPage int) (res []bannerdomain.Entity, err error) {
	panic("unimplemented")
}
