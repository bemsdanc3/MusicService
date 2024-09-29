package requests

import (
	"database/sql"
	"errors"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID    uint   `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

func CreateUser(db *sql.DB, user User) error {
	// Хэширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Pass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO users (login, email, pass) VALUES (?, ?, ?)`
	_, err = db.Exec(query, user.Login, user.Email, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func GetAllUsers(db *sql.DB) ([]User, error) {
	query := `SELECT id, login, email from users`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Login, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUserById(db *sql.DB, id uint) (User, error) {
	var user User
	query := `Select id, login, email FROM users WHERE id = ?`
	err := db.QueryRow(query, id).Scan(&user.ID, &user.Login, &user.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("User not found")
		}
		return user, err
	}

	return user, nil
}

func UpdateUser(db *sql.DB, id uint, user User) error {
	// Проверяем, указан ли новый пароль
	if user.Pass == "" {
		return errors.New("password must be provided")
	}

	// Хешируем новый пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Pass), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Формируем SQL-запрос для обновления логина, email и пароля
	query := `UPDATE users SET login = ?, email = ?, pass = ? WHERE id = ?`
	result, err := db.Exec(query, user.Login, user.Email, hashedPassword, id)
	if err != nil {
		// Проверяем, была ли ошибка уникальности
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return errors.New("email already exists")
		}
		return err
	}

	// Проверка на количество затронутых строк
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("user not found")
	}

	return nil
}
