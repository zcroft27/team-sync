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
	makeDBQuery(context.Background(), pgx.Conn{})
}

func makeDBQuery(ctx context.Context, db pgx.Conn) ([]User, error) {
	rows, err := db.Query(ctx, "SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
