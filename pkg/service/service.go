package service

import (
	"advert"
	"advert/pkg/repository"
)

type AdvertList interface {
	Create(advert advert.Advert) (int, error)
	GetById(id int, fields string) (advert.AdvertDTO, error)
	GetAll(page int, typeSort, subjectSort string) ([]advert.AdvertDTO, error)
}

type Service struct {
	AdvertList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		NewAdvertService(repos.AdvertList),
	}
}
