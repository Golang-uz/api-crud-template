package storage

import (
	"context"

	"github.com/realtemirov/api-crud-template/models"
)

type StorageI interface {
	CloseDB() error
	User() UserI
	Post() PostI
}

type UserI interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByID(ctx context.Context, id int) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByUserName(ctx context.Context, userName string) (*models.User, error)
	GetUsers(ctx context.Context, meta *models.Meta) (*models.GetAllUsersResponse, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id int) (*models.User, error)
	RemoveFromDB(ctx context.Context, id int) error
}

type PostI interface {
	CreatePost(ctx context.Context, post *models.Post) (*models.Post, error)
	GetPostByID(ctx context.Context, id int) (*models.Post, error)
	GetPostByUserID(ctx context.Context, userID int, meta *models.Meta) (*models.GetAllPostsResponse, error)
	GetPosts(ctx context.Context, meta *models.Meta) (*models.GetAllPostsResponse, error)
	UpdatePost(ctx context.Context, post *models.Post) (*models.Post, error)
	DeletePost(ctx context.Context, id int) (*models.Post, error)
	RemoveFromDB(ctx context.Context, id int) error
}
