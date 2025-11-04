package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID       int64
	Username string
	Email    string
	Password string
}

type UserRepository interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetByID(ctx context.Context, id int64) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	UpdatePassword(ctx context.Context, userId int64, newPassword string) error
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user *User) (*User, error) {
	query := `INSERT INTO users (username, email, password, created_at) VALUES ($1, $2, $3, NOW()) RETURNING id`

	err := r.db.QueryRow(ctx, query, user.Username, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetByID(ctx context.Context, id int64) (*User, error) {
	query := `SELECT id, username, email, password FROM users WHERE id = $1 AND deleted_at IS NULL`

	user := &User{}
	err := r.db.QueryRow(ctx, query, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*User, error) {
	query := `SELECT id, username, email, password FROM users WHERE username = $1 AND deleted_at IS NULL`

	user := &User{}
	err := r.db.QueryRow(ctx, query, username).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := `SELECT id, username, email, password FROM users WHERE email = $1 AND deleted_at IS NULL`

	user := &User{}
	err := r.db.QueryRow(ctx, query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) UpdatePassword(ctx context.Context, userId int64, newPassword string) error {
	query := `UPDATE users SET password = $1, updated_at = NOW() WHERE id = $2`

	_, err := r.db.Exec(ctx, query, newPassword, userId)
	return err
}
