package sqlite

import (
	"database/sql"

	"github.com/druc/posty/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, password)
	VALUES (?, ?, ?)`

	_, err = m.DB.Exec(stmt, name, email, passwordHash)
	if err != nil {
		return err
	}

	return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	var id int
	var passwordHash []byte

	stmt := `SELECT id, password FROM users WHERE email = ?`
	row := m.DB.QueryRow(stmt, email)
	err := row.Scan(&id, &passwordHash)
	if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword(passwordHash, []byte(password))
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *UserModel) Find(id int) (models.User, error) {
	var u models.User

	stmt := `SELECT id, name, email, password FROM users WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	if err != nil {
		return models.User{}, err
	}

	return u, nil
}
