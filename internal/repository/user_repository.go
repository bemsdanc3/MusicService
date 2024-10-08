package repository

import (
	"database/sql"
	"musicService/internal/entities"
)

type UserRepository interface {
	GetAllUsers() ([]entities.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAllUsers() ([]entities.User, error) {
	users := []entities.User{}
	rows, err := r.db.Query("SELECT ID, Login, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entities.User
		if err = rows.Scan(&user.ID, &user.Login, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
