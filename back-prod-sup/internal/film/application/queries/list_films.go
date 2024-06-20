package queries

import "github.com/felipemalacarne/back-prod-sup/internal/film/domain"

type ListFilmsQuery struct{}

type ListFilmsHandler struct {
	repository domain.FilmRepository
}

func NewListFilmsHandler(repository domain.FilmRepository) *ListFilmsHandler {
    return &ListFilmsHandler{repository}
}

func (h ListFilmsHandler) Handle(query ListFilmsQuery) ([]domain.Film, error) {
    return h.repository.FindAll()
}
