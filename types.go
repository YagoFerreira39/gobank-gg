package main

import (
	"math/rand"
	"time"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Number int64 `json:"number"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Number int64  `json:"number"`
	Token  string `json:"token"`
}

type TransferRequest struct {
	ToAccount int `json:"toAccount"`
	Amount int `json:"amount"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastname"`
	Password string `json:"password"`
}

type Account struct {
	ID int `json: "id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastname"`
	Number int64 `json:"number"`
	EncryptedPassword string `json:"-"`
	Balance int64 `json:"balance"`
	CreatedAt time.Time `json:'createdAt"`
}

func (account *Account) ValidatePassword( password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(account.EncryptedPassword), []byte(password)) == nil
}

func NewAccount (firstName, lastName string, password string) (*Account, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Account {
		FirstName: firstName,
		LastName: lastName,
		EncryptedPassword: string(encryptedPassword),
		Number: int64(rand.Intn(100000)),
		CreatedAt: time.Now().UTC(),
	}, nil
}