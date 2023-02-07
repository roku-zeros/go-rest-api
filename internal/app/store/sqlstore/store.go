package sqlstore

import (
	"database/sql"
	"github.com/roku-zeros/go-rest-api/internal/app/store"
)

type Store struct {
	db             *sql.DB
	UserRepository *UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}

	s.UserRepository = &UserRepository{
		Store: s,
	}

	return s.UserRepository
}
