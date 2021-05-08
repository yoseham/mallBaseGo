package services

import (
	"app/repositories"
	"fmt"
)
type MovieService interface {
	ShowMovieName() string
}

type MovieServiceManager struct {
	repo repositories.MovieRepository
}

func NewMovieServiceManager(repo repositories.MovieRepository) MovieService {
	return &MovieServiceManager{repo: repo}
}

func (m *MovieServiceManager) ShowMovieName() string{
	fmt.Println(m.repo.GetMovieName())
	return m.repo.GetMovieName()
}
