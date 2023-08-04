package repository

import (
	"database/sql"
	"fmt"
	domain "xsis/movie/domain/movie"
)

type MovieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (mysql *MovieRepository) Get() ([]domain.MovieList, error) {

	var lists []domain.MovieList

	rows, err := mysql.db.Query("SELECT * FROM movie")

	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var list domain.MovieList
		if err := rows.Scan(&list.ID, &list.Title, &list.Description, &list.Rating, &list.Image, &list.CreatedAt, &list.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error: %v", err)
		}
		lists = append(lists, list)
	}

	return lists, nil
}

func (mysql *MovieRepository) GetById(id int) (domain.MovieList, error) {

	var data domain.MovieList
	row := mysql.db.QueryRow("SELECT * FROM movie WHERE id = ?", id)
	if err := row.Scan(&data.ID, &data.Title, &data.Description, &data.Rating, &data.Image, &data.CreatedAt, &data.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return domain.MovieList{}, fmt.Errorf("moviebyid %d: no such movie", id)
		} else {
			return domain.MovieList{}, fmt.Errorf("moviebyid %d: %v", id, err)
		}
	}

	return data, nil
}

func (mysql *MovieRepository) Add(payload domain.MovieList) error {
	_, err := mysql.db.Exec("INSERT INTO movie (title,description,rating,image) VALUES (?,?,?,?)",
		payload.Title,
		payload.Description,
		payload.Rating,
		payload.Image,
	)

	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

func (mysql *MovieRepository) Update(payload domain.MovieList) error {
	_, err := mysql.db.Exec("UPDATE movie SET title = ?, description = ?, rating = ?, image = ?",
		payload.Title,
		payload.Description,
		payload.Rating,
		payload.Image,
	)

	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

func (mysql *MovieRepository) Delete(id int) error {

	_, err := mysql.db.Exec("DELETE FROM movie WHERE id = ?", id)

	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	return nil
}
