package service

import (
	"advert"
	"advert/pkg/repository"
)

type AdvertService struct {
	repo repository.AdvertList
}

func NewAdvertService(repo repository.AdvertList) *AdvertService {
	return &AdvertService{repo: repo}
}

func (s *AdvertService) Create(advert advert.Advert) (int, error) {
	return s.repo.Create(advert)
}

func (s *AdvertService) GetById(id int, fields string) (advert.AdvertDTO, error) {
	return s.repo.GetById(id, fields)
}

func (s *AdvertService) GetAll(page int, typeSort, subjectSort string) ([]advert.AdvertDTO, error) {
	return s.repo.GetAll(page, typeSort, subjectSort)
}
