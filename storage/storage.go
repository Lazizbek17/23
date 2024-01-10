package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage struct {
	db         *sql.DB
	// User       *User
	// Author     *Author
	// Post *Post
	Like *Like
}

func NewStorage(ConnStr string) *Storage {
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		return nil
	}

	err = db.Ping()
	if err != nil {
		return nil
	}

	return &Storage{
		db:         db,
		Like: 	NewLike(db),
	}
}

func (s *Storage) Close() {
	s.db.Close()
}