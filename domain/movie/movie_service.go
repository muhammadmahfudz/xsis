package domain

type MovieService interface {
	MovieLists() (interface{}, error)
	MovieById(movieID int) (interface{}, error)
	MovieAdd(payload MovieList) error
	MovieUpdate(payload MovieList) error
	MovieDelete(movieID int) error
}
