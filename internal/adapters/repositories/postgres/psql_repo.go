package repositories

import (
	"context"
	"fmt"

	"github.com/jmechavez/my-hexagonal-app/internal/core/domain"
	"github.com/jmechavez/my-hexagonal-app/internal/core/ports"
	database "github.com/jmechavez/my-hexagonal-app/pkg/database/sqlc"
)

type postgresRepository struct {
	queries *database.Queries
}

// NewPostgresRepository creates a new PostgreSQL repository
func NewPostgresRepository(queries *database.Queries) ports.UserRepository {
	return &postgresRepository{
		queries: queries,
	}
}

func (r *postgresRepository) Create(ctx context.Context, user *domain.User) error {
	result, err := r.queries.CreateUser(ctx, database.CreateUserParams{
		Name:  user.Name,
		Email: user.Email,
	})
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	user.ID = result.ID
	user.CreatedAt = result.CreatedAt
	user.UpdatedAt = result.UpdatedAt
	return nil
}

func (r *postgresRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	result, err := r.queries.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID %d: %w", id, err)
	}

	user := &domain.User{
		ID:        result.ID,
		Name:      result.Name,
		Email:     result.Email,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	return user, nil
}

func (r *postgresRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	result, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email %s: %w", email, err)
	}

	user := &domain.User{
		ID:        result.ID,
		Name:      result.Name,
		Email:     result.Email,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}

	return user, nil
}

func (r *postgresRepository) List(ctx context.Context) ([]domain.User, error) {
	results, err := r.queries.ListUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	users := make([]domain.User, len(results))
	for i, result := range results {
		users[i] = domain.User{
			ID:        result.ID,
			Name:      result.Name,
			Email:     result.Email,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		}
	}

	return users, nil
}

func (r *postgresRepository) Update(ctx context.Context, user *domain.User) error {
	result, err := r.queries.UpdateUser(ctx, database.UpdateUserParams{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
	if err != nil {
		return fmt.Errorf("failed to update user ID %d: %w", user.ID, err)
	}

	user.Name = result.Name
	user.Email = result.Email
	user.UpdatedAt = result.UpdatedAt
	return nil
}

func (r *postgresRepository) Delete(ctx context.Context, id int64) error {
	err := r.queries.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete user ID %d: %w", id, err)
	}
	return nil
}
