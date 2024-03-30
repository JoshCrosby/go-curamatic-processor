package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func CreateTableQuery(p *pgxpool.Pool) {
	_, err := p.Exec(context.Background(), "CREATE TABLE users (id SERIAL PRIMARY KEY,name VARCHAR(255) NOT NULL,email VARCHAR(255) UNIQUE NOT NULL);")
	if err != nil {
		log.Fatal("Error while creating the table")
	}
}

func InsertQuery(p *pgxpool.Pool) {
	_, err := p.Exec(context.Background(), "insert into users(name, email) values($1, $2)", "Josh", "joshdcrosby@gmail.com")
	if err != nil {
		log.Fatal("Error while inserting value into the table")
	}
}

func SelectQuery(p *pgxpool.Pool) {
	rows, err := p.Query(context.Background(), "select name, email from users")
	if err != nil {
		log.Fatal("Error while selecting value from the table")
	}
	defer rows.Close()
	for rows.Next() {
		var name, email string
		err = rows.Scan(&name, &email)
		if err != nil {
			log.Fatal("Error while scanning the rows")
		}
		log.Println(name, email)
	}
}
