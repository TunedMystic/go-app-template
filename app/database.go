package main

type Database struct{}

func (db *Database) AllNotes() ([]*Note, error) {
	return []*Note{
		{
			Title: "Distributed Cache and Pub/Sub with Redis",
		},
		{
			Title: "Quant Trading - Simulate Market Trends",
		},
		{
			Title: "Building a Crypto Exchange in Go",
		},
	}, nil
}

func (db *Database) ActiveUsers() ([]*User, error) {
	return []*User{
		{
			Email: "alice@server1.com",
		},
		{
			Email: "bob@rootnode.com",
		},
	}, nil
}

func (db *Database) GetUser(email string) (*User, error) {
	return &User{
		Email: "bob@rootnode.com",
	}, nil
}
