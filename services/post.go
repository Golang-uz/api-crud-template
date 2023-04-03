package services

import (
	"context"

	"github.com/realtemirov/api-crud-template/models"
	"github.com/realtemirov/api-crud-template/storage"
	"github.com/rs/zerolog"
)

type PostService struct {
	log  zerolog.Logger
	repo storage.PostI
}

func NewPostService(repo storage.PostI, log zerolog.Logger) *PostService {
	return &PostService{
		log:  log,
		repo: repo,
	}
}

// type PostI interface {
// 	CreatePost(ctx context.Context, post *models.Post) (*models.Post, error)
// 	GetPostByID(ctx context.Context, id int) (*models.Post, error)
// 	GetPostByUserID(ctx context.Context, userID int) (*models.Post, error)
// 	GetPosts(ctx context.Context, meta *models.Meta) (*models.GetAllPostsResponse, error)
// 	UpdatePost(ctx context.Context, post *models.Post) (*models.Post, error)
// 	DeletePost(ctx context.Context, id int) (*models.Post, error)
// }

func (p *PostService) CreatePost(ctx context.Context, post *models.Post) (*models.Post, error) {

	// TODO

	return p.repo.CreatePost(ctx, post)
}
func (p *PostService) GetPostByID(ctx context.Context, id int) (*models.Post, error) {

	// TODO

	return p.repo.GetPostByID(ctx, id)
}
func (p *PostService) GetPostByUserID(ctx context.Context, userID int, meta *models.Meta) (*models.GetAllPostsResponse, error) {

	// TODO

	return p.repo.GetPostByUserID(ctx, userID, meta)
}
func (p *PostService) GetPosts(ctx context.Context, meta *models.Meta) (*models.GetAllPostsResponse, error) {

	// TODO

	return p.repo.GetPosts(ctx, meta)
}
func (p *PostService) UpdatePost(ctx context.Context, post *models.Post) (*models.Post, error) {

	// TODO

	return p.repo.UpdatePost(ctx, post)
}
func (p *PostService) DeletePost(ctx context.Context, id int) (*models.Post, error) {

	// TODO

	return p.repo.DeletePost(ctx, id)
}
