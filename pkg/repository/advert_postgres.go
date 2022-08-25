package repository

import (
	"advert"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AdvertPostgres struct {
	db *pgxpool.Pool
}

func NewAdvertPostgres(db *pgxpool.Pool) *AdvertPostgres {
	return &AdvertPostgres{db: db}
}

func (r *AdvertPostgres) Create(advert advert.Advert) (int, error) {
	var id int
	query := "INSERT INTO adverts (title, description, photos, price, create_date) values ($1, $2, $3::text[], $4, $5) RETURNING id"
	err := r.db.QueryRow(context.Background(), query, advert.Title, advert.Description, advert.Photos, advert.Price, advert.CreateDate).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (r *AdvertPostgres) GetById(id int, fields string) (advert.AdvertDTO, error) {
	var advert advert.AdvertDTO
	if fields == "true"{
		query := "SELECT description, photos from adverts where id = $1"
		if err := r.db.QueryRow(context.Background(), query, id).Scan(&advert.Description, &advert.Photos); err != nil {
			return advert, err
		}
		return advert, nil
	}
	advert.Photos = append(advert.Photos, "")
	query := "SELECT title, price, photos[1] from adverts where id = $1"
	if err := r.db.QueryRow(context.Background(), query, id).Scan(&advert.Title, &advert.Price, &advert.Photos[0]); err != nil {
		return advert, err
	}
	return advert, nil
}

func (r *AdvertPostgres) GetAll(page int, typeSort, subjectSort string) ([]advert.AdvertDTO, error) {
	var adverts []advert.AdvertDTO
	query := fmt.Sprintf(`SELECT id, title, photos[1], price from adverts
					      order by %s %s
						  offset %d limit 10`, subjectSort, typeSort, (page - 1) * 10)
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		return adverts, err
	}
	for rows.Next(){
		advert := advert.AdvertDTO{}
		advert.Photos = append(advert.Photos, "")
		err := rows.Scan(&advert.Id, &advert.Title, &advert.Photos[0], &advert.Price)
		if err != nil{
			log.Println(err)
		}
		adverts = append(adverts, advert)
	}

	if rows.Err() != nil {
		log.Printf("rows err %s\n", err)
		return adverts, rows.Err()
	}
	return adverts, nil
}

/*
SELECT title, photos[1], price, create_date from adverts
order by price asc
offset 2 limit 10
*/