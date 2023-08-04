package domain

import "time"

type Movie struct {
	id          int
	title       string
	description string
	rating      float64
	image       string
	createdAt   time.Time
	updatedAt   time.Time
}

func NewMovie() *Movie {
	return &Movie{}
}

func (movie *Movie) Build() *Movie {
	return &Movie{
		id:          movie.id,
		title:       movie.title,
		description: movie.description,
		rating:      movie.rating,
		image:       movie.image,
		createdAt:   movie.createdAt,
		updatedAt:   movie.updatedAt,
	}
}

func (movie *Movie) SetID(id int) {
	movie.id = id
}

func (movie *Movie) SetTitle(title string) {
	movie.title = title
}

func (movie *Movie) SetDescription(desc string) {
	movie.description = desc
}

func (movie *Movie) SetRating(rating float64) {
	movie.rating = rating
}

func (movie *Movie) SetImage(image string) {
	movie.image = image
}

func (movie *Movie) GetID() int {
	return movie.id
}

func (movie *Movie) GetTitle() string {
	return movie.title
}

func (movie *Movie) GetDescription() string {
	return movie.description
}

func (movie *Movie) GetRating() float64 {
	return movie.rating
}

func (movie *Movie) GetImage() string {
	return movie.image
}
