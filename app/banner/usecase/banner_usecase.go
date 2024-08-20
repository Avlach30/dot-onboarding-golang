package usecase

import (
	"context"
	"github.com/codespace-id/codespace-x/app/banner/bannerdomain"
	"github.com/codespace-id/codespace-x/app/banner/bannerdto"
	"github.com/pkg/errors"
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
func (uc *bannerUsecase) Get(ctx context.Context, page int, perPage int) (res []bannerdto.BannerResponse, err error) {

	bannerData, err := uc.bannerRepo.Get(ctx, page, perPage)
	if err != nil {
		return nil, errors.WithMessage(err, "bannerUsecase.Get")
	}

	for _, val := range bannerData {
		res = append(res, bannerdto.BannerResponse{
			Name:        val.Title,
			Description: val.Description,
			SourceURL:   val.ImageURL,
		})
	}

	return res, nil
}
