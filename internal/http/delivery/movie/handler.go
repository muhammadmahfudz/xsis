package movie_handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	domain "xsis/movie/domain/movie"

	"github.com/gorilla/mux"
)

type HandlerMovie struct {
	usecase domain.MovieService
}

func NewHandlerMovie(usecase domain.MovieService) *HandlerMovie {
	return &HandlerMovie{usecase: usecase}
}

func (ha *HandlerMovie) HandleMovieList(w http.ResponseWriter, r *http.Request) {
	resp, err := ha.usecase.MovieLists()
	if err != nil {
		fmt.Println("====================")
		fmt.Printf("error: %v", err)
		fmt.Println("====================")
		BuldErrorResponse(w, false, http.StatusInternalServerError, "internal server error")
		return
	}
	BuildResponse(w, true, http.StatusOK, resp)
}

func (ha *HandlerMovie) HandleMovieById(w http.ResponseWriter, r *http.Request) {
	va := mux.Vars(r)
	id, _ := strconv.Atoi(va["id"])
	resp, _ := ha.usecase.MovieById(id)
	BuildResponse(w, true, http.StatusOK, resp)
}

func (ha *HandlerMovie) HandleMovieAdd(w http.ResponseWriter, r *http.Request) {
	var p domain.MovieList
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&p); err != nil {
		BuldErrorResponse(w, false, http.StatusBadRequest, "bad request")
		return
	}

	err := ha.usecase.MovieAdd(p)

	if err != nil {
		fmt.Println("====================")
		fmt.Printf("error: %v", err)
		fmt.Println("====================")
		BuldErrorResponse(w, false, http.StatusInternalServerError, "internal server error")
		return
	}

	BuildResponse(w, true, http.StatusOK, map[string]string{"message": "successfully"})
}

func (ha *HandlerMovie) HandleMovieUpdate(w http.ResponseWriter, r *http.Request) {
	var p domain.MovieList

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		BuldErrorResponse(w, false, http.StatusBadRequest, "bad request")
		return
	}

	err := ha.usecase.MovieUpdate(p)

	if err != nil {
		BuldErrorResponse(w, false, http.StatusInternalServerError, "internal server error")
		return
	}

	BuildResponse(w, true, http.StatusOK, map[string]string{"message": "successfully"})
}

func (ha *HandlerMovie) HandleMovieDelete(w http.ResponseWriter, r *http.Request) {
	var p domain.MovieList

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		BuldErrorResponse(w, false, http.StatusBadRequest, "bad request")
		return
	}

	err := ha.usecase.MovieDelete(1)

	if err != nil {
		BuldErrorResponse(w, false, http.StatusInternalServerError, "bad request")
		return
	}

	BuildResponse(w, true, http.StatusOK, map[string]string{"message": "successfully"})
}
