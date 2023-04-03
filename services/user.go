package services

import (
	"context"

	"github.com/realtemirov/api-crud-template/models"
	"github.com/realtemirov/api-crud-template/storage"
	"github.com/rs/zerolog"
)

type UserService struct {
	log  zerolog.Logger
	repo storage.UserI
}

func NewUserService(repo storage.UserI, log zerolog.Logger) UserService {
	return UserService{
		log:  log,
		repo: repo,
	}
}

// type UserI interface {
// 	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
// 	GetUserByID(ctx context.Context, id int) (*models.User, error)
// 	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
// 	GetUserByUserName(ctx context.Context, userName string) (*models.User, error)
// 	GetUsers(ctx context.Context, meta *models.Meta) (*models.GetAllUsersResponse, error)
// 	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
// 	DeleteUser(ctx context.Context, id int) (*models.User, error)
// }

func (u *UserService) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {

	// TODO

	return u.repo.CreateUser(ctx, user)
}

func (u *UserService) GetUserByID(ctx context.Context, id int) (*models.User, error) {

	// TODO

	return u.repo.GetUserByID(ctx, id)
}

func (u *UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {

	// TODO

	return u.repo.GetUserByEmail(ctx, email)
}

func (u *UserService) GetUserByUserName(ctx context.Context, userName string) (*models.User, error) {

	// TODO

	return u.repo.GetUserByUserName(ctx, userName)
}

func (u *UserService) GetUsers(ctx context.Context, meta *models.Meta) (*models.GetAllUsersResponse, error) {

	// TODO

	return u.repo.GetUsers(ctx, meta)
}

func (u *UserService) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {

	// TODO

	return u.repo.UpdateUser(ctx, user)
}

func (u *UserService) DeleteUser(ctx context.Context, id int) (*models.User, error) {

	// TODO

	return u.repo.DeleteUser(ctx, id)
}
