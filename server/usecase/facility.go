package usecase

import (
	"context"

	"github.com/homma509/9rece/server/domain/model"
	"github.com/homma509/9rece/server/domain/repository"
)

// FacilityUsecase 施設ユースケースのインターフェース
type FacilityUsecase interface {
	Store(context.Context, model.Facilities) error
}

type facilityUsecase struct {
	facilityRepository repository.FacilityRepository
}

// NewFacilityUsecase 施設ユースケースを生成します
func NewFacilityUsecase(r repository.FacilityRepository) FacilityUsecase {
	return &facilityUsecase{
		facilityRepository: r,
	}
}

// Store 施設を登録します
func (u *facilityUsecase) Store(ctx context.Context, facilities model.Facilities) error {
	return u.facilityRepository.Save(ctx, facilities)
}
