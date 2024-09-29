package requests

import "database/sql"

type User struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

func CreateUser(db *sql.DB, user User) error {
	query := `INSERT INTO users (login, email, pass) VALUES (?, ?, ?)`

	_, err := db.Exec(query, user.Login, user.Email, user.Pass)

	if err != nil {
		return err
	}

	return nil
}
