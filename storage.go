package main

import (
	"database/sql"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)


type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
	// Get(key string) (string, error)
	// Set(key string, value string) error
}

type PostgresStorage struct {
	db *sql.DB
}

// GetAccountByID implements Storage.
func (s *PostgresStorage) GetAccountByID(int) (*Account, error) {
	panic("unimplemented")
}

func NewPostgresStorage() (*PostgresStorage, error) {
	connStr := "user=postgres dbname=postgres password=newpassword host=localhost port=5433 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) Init() error {
	return s.CreateAccountTable() 
}

func (s *PostgresStorage) CreateAccountTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS accounts (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NOT NULL,
		number SERIAL NOT NULL,
		balance SERIAL NOT NULL,
		CREATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP	
	)`
	_, err := s.db.Exec(query)	
	return err			
	
}
func (s *PostgresStorage) CreateAccount(*Account) error {
	return nil
}

func (s *PostgresStorage) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStorage) DeleteAccount(id int) error {
	return nil
}

func (s *PostgresStorage) getAccountByID(id int) (*Account, error) {
	return nil, nil
}
