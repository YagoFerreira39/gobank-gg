package main

import (
	"fmt"
	"log"
	"flag"
)

func seedAccount(store *PostgresStore, firstName, lastName, password string) *Account {
	account, err := NewAccount(firstName, lastName, password)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(account); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new account => ", account.Number)

	return account
}

func seedAccounts(s *PostgresStore) {
	seedAccount(s, "Nikola", "Jokic", "bigboatmvp")
}

func main() {
	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	
	// seed stuff
	if *seed {
		fmt.Println("seeding the database")
		seedAccounts(store)
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}