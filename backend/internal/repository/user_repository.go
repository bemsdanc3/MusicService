package repository

import (
	"database/sql"
	"musicService/internal/entities"
)

type UserRepository interface {
	GetAllUsers() ([]entities.User, error)
	GetUserByID(id int) (entities.User, error)
	CreateUser(user entities.User) (int64, error)
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

func (r *userRepository) GetUserByID(id int) (entities.User, error) {
	var user entities.User
	err := r.db.QueryRow("SELECT Login, email FROM users WHERE ID = ?", id).Scan(&user.Login, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) CreateUser(user entities.User) (int64, error) {
	result, err := r.db.Exec("INSERT INTO users (Login, email, pass) VALUES (?, ?, ?)", user.Login, user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
