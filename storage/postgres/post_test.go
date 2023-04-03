package postgres_test

import (
	"math/rand"
	"testing"

	"github.com/google/uuid"
	"github.com/realtemirov/api-crud-template/models"
	"github.com/stretchr/testify/require"
)

func createRandomPost(t *testing.T) *models.Post {

	result, err := str.Post().CreatePost(ctx, &models.Post{
		Title:  uuid.NewString(),
		Body:   uuid.NewString(),
		UserID: int64(rand.Intn(100)),
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)

	return result
}

func deleteRandomPost(t *testing.T, post *models.Post) {
	delete, err := str.Post().DeletePost(ctx, post.ID)
	require.NoError(t, err)
	require.NotEmpty(t, delete)

	require.Equal(t, post.ID, delete.ID)
	require.Equal(t, post.UserID, delete.UserID)
	require.Equal(t, post.Title, delete.Title)
	require.Equal(t, post.Body, delete.Body)
	require.Equal(t, post.CreatedAt, delete.CreatedAt)

	err = str.Post().RemoveFromDB(ctx, post.ID)
	require.NoError(t, err)
}

func Test_CreatePost(t *testing.T) {
	defer deleteRandomPost(t, createRandomPost(t))
}

func Test_DeletePost(t *testing.T) {
	deleteRandomPost(t, createRandomPost(t))
}

func Test_GetPostByID(t *testing.T) {
	post := createRandomPost(t)
	defer deleteRandomPost(t, post)

	u, err := str.Post().GetPostByID(ctx, post.ID)
	require.NoError(t, err)
	require.NotEmpty(t, u)
	require.NotEqual(t, u.CreatedAt, 0)
	require.Equal(t, post.UpdatedAt, 0)
	require.Equal(t, post.DeletedAt, 0)

	require.Equal(t, post.ID, u.ID)
	require.Equal(t, post.UserID, u.UserID)
	require.Equal(t, post.Title, u.Title)
	require.Equal(t, post.Body, u.Body)
	require.Equal(t, post.CreatedAt, u.CreatedAt)
	require.Equal(t, post.UpdatedAt, u.UpdatedAt)
	require.Equal(t, post.DeletedAt, u.DeletedAt)
}

func Test_GetPostByUserID(t *testing.T) {
	post := createRandomPost(t)
	defer deleteRandomPost(t, post)

	u, err := str.Post().GetPostByUserID(ctx, int(post.UserID), &models.Meta{
		PerPage:     1,
		CurrentPage: 1,
	})
	require.NoError(t, err)
	require.NotEmpty(t, u)

	require.Equal(t, post.ID, u.Data[0].ID)
	require.Equal(t, post.UserID, u.Data[0].UserID)
	require.Equal(t, post.Title, u.Data[0].Title)
	require.Equal(t, post.Body, u.Data[0].Body)
	require.Equal(t, post.CreatedAt, u.Data[0].CreatedAt)

	require.Equal(t, len(u.Data), 1)
}
func Test_GetPosts(t *testing.T) {

	var Posts []*models.Post
	for i := 0; i < 10; i++ {
		Post := createRandomPost(t)
		defer deleteRandomPost(t, Post)
		Posts = append(Posts, Post)
	}

	u, err := str.Post().GetPosts(ctx, &models.Meta{
		PerPage:     10,
		CurrentPage: 1,
	})

	require.NoError(t, err)
	require.NotEmpty(t, u)
	require.Equal(t, len(u.Data), len(Posts))

}
func Test_UpdatePost(t *testing.T) {
	post := createRandomPost(t)
	defer deleteRandomPost(t, post)

	post.Title = uuid.NewString()
	post.Body = uuid.NewString()

	u, err := str.Post().UpdatePost(ctx, post)
	require.NoError(t, err)
	require.NotEmpty(t, u)
	require.NotEqual(t, u.UpdatedAt, 0, "this")

	require.Equal(t, post.ID, u.ID)
	require.Equal(t, post.UserID, u.UserID)
	require.Equal(t, post.Title, u.Title)
	require.Equal(t, post.Body, u.Body)
	require.Equal(t, post.CreatedAt, u.CreatedAt)
	require.NotEqual(t, post.UpdatedAt, u.UpdatedAt)
	require.Equal(t, post.DeletedAt, u.DeletedAt)
}
