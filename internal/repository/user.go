package repository

import (
	"context"
	"errors"

	"github.com/alexbednarczyk/crud-gopher/internal/api/v0alpha/models"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	ErrNotFound = errors.New("not found")
)

// UserRepository defines the interface for user data access.
type UserRepository interface {
	Create(ctx context.Context, user *models.UserResponse) error
	GetByGUID(ctx context.Context, guid string) (*models.UserResponse, error)
	GetAll(ctx context.Context) ([]models.UserResponse, error)
	Update(ctx context.Context, user *models.UserResponse) error
	Delete(ctx context.Context, guid string) error
}

// PostgresUserRepository implements UserRepository for PostgreSQL.
type PostgresUserRepository struct {
	db *pgxpool.Pool
}

// NewPostgresUserRepository creates a new PostgresUserRepository.
func NewPostgresUserRepository(db *pgxpool.Pool) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(ctx context.Context, user *models.UserResponse) error {
	query := `
		INSERT INTO users (guid, display_name, first_name, last_name, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.Exec(ctx, query, user.GUID, user.DisplayName, user.FirstName, user.LastName, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *PostgresUserRepository) GetByGUID(ctx context.Context, guid string) (*models.UserResponse, error) {
	query := `SELECT guid, display_name, first_name, last_name, created_at, updated_at FROM users WHERE guid=$1`
	row := r.db.QueryRow(ctx, query, guid)

	var user models.UserResponse
	err := row.Scan(&user.GUID, &user.DisplayName, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) GetAll(ctx context.Context) ([]models.UserResponse, error) {
	query := `SELECT guid, display_name, first_name, last_name, created_at, updated_at FROM users`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.UserResponse
	for rows.Next() {
		var user models.UserResponse
		if err := rows.Scan(&user.GUID, &user.DisplayName, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, rows.Err()
}

func (r *PostgresUserRepository) Update(ctx context.Context, user *models.UserResponse) error {
	query := `
		UPDATE users
		SET display_name=$2, first_name=$3, last_name=$4, created_at=$5, updated_at=$6
		WHERE guid=$1
	`
	commandTag, err := r.db.Exec(ctx, query, user.GUID, user.DisplayName, user.FirstName, user.LastName, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *PostgresUserRepository) Delete(ctx context.Context, guid string) error {
	query := `DELETE FROM users WHERE guid=$1`
	commandTag, err := r.db.Exec(ctx, query, guid)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return ErrNotFound
	}
	return nil
}
