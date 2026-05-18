package models

type Product struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Price float64 `json:"price"`
	ImageUrl string `json:"image_url"`
	Description string `json:"description"`
	Rating float64 `json:"rating"`
	ReviewCount uint `json:"review_count"`
}