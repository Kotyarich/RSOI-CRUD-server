package user

import (
	"context"
	"rsoi-2021-lab1-ci-cd-Kotyarich/models"
)

type UseCase interface {
	Create(ctx context.Context, user *models.User) (int, error)
	GetProfile(ctx context.Context, id int) (*models.User, error)
	ChangeProfile(ctx context.Context, user *models.User, id int) (*models.User, error)
	DeletePerson(ctx context.Context, id int) error
	GetAllPersons(ctx context.Context) ([]*models.User, error)
}
