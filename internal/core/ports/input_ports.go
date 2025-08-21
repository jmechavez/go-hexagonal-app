package ports

import (
	"context"

	"github.com/jmechavez/my-hexagonal-app/internal/core/domain"
)

type UserService interface {
	CreateUser(ctx context.Context, req domain.UserCreateRequest) (*domain.User, error)
	GetUserByID(ctx context.Context, id int64) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	ListUsers(ctx context.Context) ([]domain.User, error)
	UpdateUser(ctx context.Context, req domain.UserUpdateRequest) (*domain.User, error)
	DeleteUser(ctx context.Context, id int64) error
}
