package repository

import (
	"advert"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AdvertList interface {
	Create(advert advert.Advert) (int, error)
	GetById(id int, fields string) (advert.AdvertDTO, error)
	GetAll(page int, typeSort, subjectSort string) ([]advert.AdvertDTO, error)
}

type Repository struct {
	AdvertList
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		NewAdvertPostgres(db),
	}
}
