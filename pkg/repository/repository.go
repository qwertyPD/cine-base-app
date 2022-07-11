package repository

type Authorization interface {
}

type MovieList interface {
}

type MovieItem interface {
}

type Repository struct {
	Authorization
	MovieList
	MovieItem
}

func NewRepository() *Repository {
	return &Repository{}
}
