package main

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	makeDBQuery(context.Background(), &pgx.Conn{})
}

func makeDBQuery(ctx context.Context, db *pgx.Conn) ([]User, error) {
	rows, err := db.Query(ctx, "SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	return pgx.CollectRows(rows, pgx.RowToStructByName[User])
}
