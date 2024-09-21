package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

type PostgressStorage struct {
	db *sql.DB
}

func NewPostgressStorage() (*PostgressStorage, error) {
	connStr := "user=postgres password=041199 dbname=postgres port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgressStorage{
		db: db,
	}, nil
}

func (s *PostgressStorage) CreateAccount(account *Account) error {
	return nil
}

func (s *PostgressStorage) DeleteAccount(id int) error {
	return nil
}

func (s *PostgressStorage) UpdateAccount(account *Account) error {
	return nil
}

func (s *PostgressStorage) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}