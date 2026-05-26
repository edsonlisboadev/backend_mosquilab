package repository

import (
	"database/sql"
	"mosquilab/internal/models"
)

type UserRepo struct{ db *sql.DB }

func NewUserRepo(db *sql.DB) *UserRepo { return &UserRepo{db: db} }

func (r *UserRepo) FindByEmail(email string) (*models.User, error) {
	var u models.User
	err := r.db.QueryRow(`SELECT id, email, password_hash, created_at FROM users WHERE email=$1`, email).
		Scan(&u.ID, &u.Email, &u.PasswordHash, &u.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &u, err
}

// Seed creates a default admin user if none exists (run on startup).
func (r *UserRepo) SeedAdmin(email, hash string) error {
	_, err := r.db.Exec(`
		INSERT INTO users (email, password_hash) VALUES ($1, $2)
		ON CONFLICT (email) DO NOTHING`, email, hash)
	return err
}
