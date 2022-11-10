package users

import (
	"context"
	"dynamodb_go/models"
)

type Service interface {
	Store(context.Context, *models.User) error
	GetOne(context.Context, string) (*models.User, error)
	UpdateEmailAndPass(context.Context, string, string, string) error
	Delete(context.Context, string) error
}

type service struct {
	rDynamo RepositoryDynamo
}

func NewService(repoDynamo RepositoryDynamo) Service {
	return &service{
		rDynamo: repoDynamo,
	}
}

func (s *service) GetOne(ctx context.Context, id string) (*models.User, error) {
	return s.rDynamo.GetOne(ctx, id)
}

func (s *service) Store(ctx context.Context, u *models.User) error {
	return s.rDynamo.Store(ctx, u)
}

func (s *service) UpdateEmailAndPass(ctx context.Context, email string, pass string, id string) error{
	return s.rDynamo.UpdateEmailAndPass(ctx, email, pass, id)
}

func (s *service) Delete(ctx context.Context, id string) error {
	return s.rDynamo.Delete(ctx, id)
}
