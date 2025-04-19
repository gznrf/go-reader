package postgres_service

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type AuthService struct {
	db *sqlx.DB
}

func (as *AuthService) CreateUser(email, name, password_hash string) (int64, error) {
	const op = "postgres.service.auth.createUser"
	var id int64
	query := fmt.Sprintf("INSERT INTO %s (name, email, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)

	row := as.db.QueryRow(query, name, email, password_hash)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	log.Printf("User created at %s", op)
	return id, nil
}
func (as *AuthService) GetUser(email, password_hash string) (int64, error) {
	const op = "postgres.service.auth.getUser"
	var id int64
	query := fmt.Sprintf("SELECT * FROM %s WHERE email = $1 AND password_hash = $2", usersTable)
	err := as.db.Get(&id, query, email, password_hash)
	if err != nil {
		log.Printf("error - %s at %s", err, op)
		return 0, err
	}

	log.Printf("User got at %s", op)
	return id, nil
}
