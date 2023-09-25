package store

import (
	"database/sql"
	"tasbih/pkg/database"
)

type UserStore struct {
	db *sql.DB
}

type Store interface {
	GetUserByEmail(string) (*database.User, error)
	Create(*database.User) error
}

type ErrorMap = map[string]string

func NewUserStore(db *sql.DB) *UserStore {
	db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		name TEXT,
		email TEXT UNIQUE,
		password TEXT
	)`)
	return &UserStore{db}
}

func (us *UserStore) Create(u *database.User) error {
	_, err := us.db.Exec(
		`INSERT INTO users(id,name,email,password) VALUES(?,?,?,?)`,
		u.ID,
		u.Name,
		u.Email,
		u.Password,
	)
	return err
}

func (us *UserStore) GetUserByEmail(e string) (*database.User, error) {
	var m database.User
	err := us.db.QueryRow("SELECT * FROM users WHERE email = ?", e).Scan(
		&m.ID, &m.Name, &m.Email, &m.Password,
	)
	return &m, err

}
