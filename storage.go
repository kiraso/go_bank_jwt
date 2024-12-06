package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)


type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	GetAccounts() ([]*Account, error)
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
	// Get(key string) (string, error)
	// Set(key string, value string) error
}

type PostgresStorage struct {
	db *sql.DB
}

// GetAccountByID implements Storage.
// func (s *PostgresStorage) GetAccountByID(int) (*Account, error) {
// 	panic("unimplemented")
// }

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
func (s *PostgresStorage) CreateAccount(account *Account) error {
	query  := `
	INSERT INTO accounts
	(first_name, last_name, number, balance,created_at) 
	VALUES ($1, $2, $3, $4,$5) 
	`
	resp, err := s.db.Query(
		query,
		account.FirstName,
		account.LastName,
		account.Number,
		account.Balance,
		account.CreateAt,
	)
	if err != nil {
		return err
	}
	fmt.Println("%v\n",resp)

	return nil
}

func (s *PostgresStorage) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStorage) DeleteAccount(id int) error {
	_,err :=  s.db.Query("DELETE FROM accounts WHERE id = $1", id)
	return err
}

func (s *PostgresStorage) GetAccountByID(id int) (*Account, error) {
	rows, err := s.db.Query("SELECT * FROM accounts WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoAccount(rows)
	}
	// account, err := scanIntoAccount(rows)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, fmt.Errorf("account %d not found",id)
}



func (s *PostgresStorage) GetAccounts() ([]*Account, error) {
	rows,err :=  s.db.Query("SELECT * FROM accounts")
	if err != nil {
		return nil, err
	}
	accounts  := []*Account{}
	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts,account)
	}
	return accounts, nil
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(
		&account.ID, 
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreateAt); 
	return account, err
}