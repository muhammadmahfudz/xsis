package usecase

import (
	"fmt"
	domain "xsis/movie/domain/movie"
)

type MovieUsecase struct {
	repository domain.MovieDomainRepository
}

func NewMovieUsecase(repository domain.MovieDomainRepository) *MovieUsecase {
	return &MovieUsecase{repository}
}

func (movie *MovieUsecase) MovieLists() (interface{}, error) {
	data, err := movie.repository.Get()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (movie *MovieUsecase) MovieById(movieID int) (interface{}, error) {
	data, err := movie.repository.GetById(movieID)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	return data, nil
}

func (movie *MovieUsecase) MovieAdd(payload domain.MovieList) error {
	err := movie.repository.Add(payload)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

func (movie *MovieUsecase) MovieUpdate(payload domain.MovieList) error {
	err := movie.repository.Update(payload)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

func (movie *MovieUsecase) MovieDelete(movieID int) error {
	err := movie.repository.Delete(movieID)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}
