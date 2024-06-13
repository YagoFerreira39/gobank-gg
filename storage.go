package main

import (
	"fmt"
	"database/sql"

	_ "github.com/lib/pq"
)
type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountById(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank09390912 sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore {
		db:  db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (
		id serial primary key,
		first_name VARCHAR(50) NOT NULL,
		last_name VARCHAR(50) NOT NULL,
		number SERIAL,
		balance serial,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(account *Account) error {
	query := `insert into account 
	(first_name, last_name, number, balance, created_at)
	values
	($1, $2, $3, $4, $5)`
	
	response, err  := s.db.Query(
		query, 
		account.FirstName, 
		account.LastName, 
		account.Number, 
		account.Balance, 
		account.CreatedAt)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", response)
	
	return nil
}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(accountId int) error {
	return nil
}

func (s *PostgresStore) GetAccountById(accountId int) (*Account, error) {
	rows, err := s.db.Query("select * from account where id = $1", accountId)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account %d not found", accountId)
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query(`select * from account`)
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account, err := scanIntoAccount(rows)
		
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err :=  rows.Scan(
		&account.ID, 
		&account.FirstName, 
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreatedAt)
	
	return account, err
}