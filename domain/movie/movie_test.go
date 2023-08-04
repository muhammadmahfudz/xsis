package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMovie(t *testing.T) {
	movie := NewMovie()
	movie.SetID(1)
	movie.SetTitle("Pengabdi Setan 2 Comunion")
	movie.SetDescription("dalah sebuah film horor Indonesia tahun 2022 yang disutradarai dan ditulis oleh Joko Anwar sebagai sekuel dari film tahun 2017, Pengabdi Setan.")
	movie.SetRating(7.2)
	movie.SetImage("")
	build := movie.Build()

	assert.Equal(t, 1, build.GetID())
	assert.Equal(t, "Pengabdi Setan 2 Comunion", build.GetTitle())
	assert.Equal(t, "dalah sebuah film horor Indonesia tahun 2022 yang disutradarai dan ditulis oleh Joko Anwar sebagai sekuel dari film tahun 2017, Pengabdi Setan.", build.GetDescription())
	assert.Equal(t, 7.2, build.GetRating())
	assert.Equal(t, "", build.GetImage())
}
