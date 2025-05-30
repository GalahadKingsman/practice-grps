package database

import (
	"database/sql"
	"fmt"
	"github.com/GalahadKingsman/messenger_users/internal/config"
	"github.com/GalahadKingsman/messenger_users/internal/models"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() error {
	cfg := config.GetConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("ошибка подключения: %w", err)
	}
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("ошибка ping: %w", err)
	}
	return createTables()
}

func createTables() error {
	// Таблицы (если не существуют)
	_, err := DB.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            login TEXT NOT NULL,
            first_name TEXT NOT NULL,
            last_name TEXT NOT NULL,
            email varchar(255) NOT NULL CHECK (email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
            phone TEXT NOT NULL
        );
    `)
	return err
}

// Методы таблицы Users

func CreateUser(user models.User) (int, error) {
	var id int
	err := DB.QueryRow(
		"INSERT INTO users (login, first_name, last_name, email, phone) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		user.Login, user.FirstName, user.LastName, user.Email, user.Phone,
	).Scan(&id)
	return id, err
}

func DeleteUser(id int) error {
	_, err := DB.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func SearchUsersByFirstNameOrLastName(firstName, lastName string) ([]models.User, error) {
	rows, err := DB.Query(
		"SELECT id, login, first_name, last_name, email, phone FROM users WHERE first_name LIKE $1 AND last_name LIKE $2",
		"%"+firstName+"%", "%"+lastName+"%",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Login, &user.FirstName, &user.LastName, &user.Email, &user.Phone)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
