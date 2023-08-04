package app

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	cf "xsis/movie/common/config"

	dbsql "xsis/movie/common/database/mysql"
	hdmv "xsis/movie/internal/http/delivery/movie"
	rpmv "xsis/movie/internal/repository/movie"
	srvmv "xsis/movie/internal/usecase/movie"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func Run(cfg *cf.Config) {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*60, "the duration for which the server gracefully wait for existing connections to finish - e.g. 1m")
	flag.Parse()

	movieRepository := dbsql.NewMysql(cfg)

	rpstrymovie := rpmv.NewMovieRepository(movieRepository)
	srvmovie := srvmv.NewMovieUsecase(rpstrymovie)
	hdlrmovie := hdmv.NewHandlerMovie(srvmovie)

	r := mux.NewRouter()

	r.HandleFunc("/movie/get/lists", hdlrmovie.HandleMovieList).Methods("GET")
	r.HandleFunc("/movie/get/list/{id}", hdlrmovie.HandleMovieById).Methods("GET")
	r.HandleFunc("/movie/add", hdlrmovie.HandleMovieAdd).Methods("POST")
	r.HandleFunc("/movie/update", hdlrmovie.HandleMovieUpdate).Methods("PUT")
	r.HandleFunc("/movie/delete", hdlrmovie.HandleMovieDelete).Methods("DELETE")

	http.Handle("/", r)

	srv := &http.Server{
		Addr:         cfg.Host.Port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
