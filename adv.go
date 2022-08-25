package advert

import "time"

type Advert struct {
	Id          int      `json:"id"`
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Photos      []string `json:"photos" binding:"required" db:"photos"`
	Price       int      `json:"price" binding:"required"`
	CreateDate time.Time `json:"create_date"`
}

type AdvertDTO struct{
	Id          int      `json:"id"`
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Photos      []string `json:"photos" binding:"required" db:"photos"`
	Price       int      `json:"price" binding:"required"`
	CreateDate time.Time `json:"create_date"`
}