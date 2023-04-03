package postgres_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/realtemirov/api-crud-template/models"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) *models.User {

	result, err := str.User().CreateUser(ctx, &models.User{
		FirstName: uuid.NewString(),
		LastName:  uuid.NewString(),
		UserName:  uuid.NewString(),
		Email:     uuid.NewString(),
		Password:  uuid.NewString(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, result)

	return result
}

func deleteRandomUser(t *testing.T, user *models.User) {
	delete, err := str.User().DeleteUser(ctx, user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, delete)

	require.Equal(t, user.ID, delete.ID)
	require.Equal(t, user.FirstName, delete.FirstName)
	require.Equal(t, user.LastName, delete.LastName)
	require.Equal(t, user.UserName, delete.UserName)
	require.Equal(t, user.Password, delete.Password)
	require.Equal(t, user.Email, delete.Email)
	require.Equal(t, user.CreatedAt, delete.CreatedAt)

	err = str.User().RemoveFromDB(ctx, user.ID)
	require.NoError(t, err)
}

func Test_CreateUser(t *testing.T) {
	defer deleteRandomUser(t, createRandomUser(t))
}

func Test_DeleteUser(t *testing.T) {
	deleteRandomUser(t, createRandomUser(t))
}

func Test_GetUserByID(t *testing.T) {
	user := createRandomUser(t)
	defer deleteRandomUser(t, user)

	u, err := str.User().GetUserByID(ctx, user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, u)
	require.NotEqual(t, u.CreatedAt, 0)
	require.Equal(t, user.UpdatedAt, 0)
	require.Equal(t, user.DeletedAt, 0)

	require.Equal(t, user.ID, u.ID)
	require.Equal(t, user.FirstName, u.FirstName)
	require.Equal(t, user.LastName, u.LastName)
	require.Equal(t, user.UserName, u.UserName)
	require.Equal(t, user.Password, u.Password)
	require.Equal(t, user.Email, u.Email)
	require.Equal(t, user.CreatedAt, u.CreatedAt)
}

func Test_GetUserByEmail(t *testing.T) {
	user := createRandomUser(t)
	defer deleteRandomUser(t, user)

	u, err := str.User().GetUserByEmail(ctx, user.Email)
	require.NoError(t, err)
	require.NotEmpty(t, u)
	require.NotEqual(t, u.CreatedAt, 0)
	require.Equal(t, user.UpdatedAt, 0)
	require.Equal(t, user.DeletedAt, 0)

	require.Equal(t, user.ID, u.ID)
	require.Equal(t, user.FirstName, u.FirstName)
	require.Equal(t, user.LastName, u.LastName)
	require.Equal(t, user.UserName, u.UserName)
	require.Equal(t, user.Password, u.Password)
	require.Equal(t, user.Email, u.Email)
	require.Equal(t, user.CreatedAt, u.CreatedAt)

}
func Test_GetUserByUserName(t *testing.T) {
	user := createRandomUser(t)
	defer deleteRandomUser(t, user)

	u, err := str.User().GetUserByUserName(ctx, user.UserName)
	require.NoError(t, err)
	require.NotEmpty(t, u)
	require.NotEqual(t, u.CreatedAt, 0)
	require.Equal(t, user.UpdatedAt, 0)
	require.Equal(t, user.DeletedAt, 0)

	require.Equal(t, user.ID, u.ID)
	require.Equal(t, user.FirstName, u.FirstName)
	require.Equal(t, user.LastName, u.LastName)
	require.Equal(t, user.UserName, u.UserName)
	require.Equal(t, user.Password, u.Password)
	require.Equal(t, user.Email, u.Email)
	require.Equal(t, user.CreatedAt, u.CreatedAt)
}
func Test_GetUsers(t *testing.T) {

	var users []*models.User
	for i := 0; i < 10; i++ {
		user := createRandomUser(t)
		defer deleteRandomUser(t, user)
		users = append(users, user)
	}

	u, err := str.User().GetUsers(ctx, &models.Meta{
		PerPage:     10,
		CurrentPage: 1,
	})

	require.NoError(t, err)
	require.NotEmpty(t, u)
	require.Equal(t, len(u.Data), len(users))

}
func Test_UpdateUser(t *testing.T) {
	user := createRandomUser(t)
	defer deleteRandomUser(t, user)

	user.FirstName = uuid.NewString()
	user.LastName = uuid.NewString()
	user.UserName = uuid.NewString()
	user.Password = uuid.NewString()
	user.Email = uuid.NewString()

	u, err := str.User().UpdateUser(ctx, user)
	require.NoError(t, err)
	require.NotEmpty(t, u)
	require.NotEqual(t, u.UpdatedAt, 0)
	require.Equal(t, user.ID, u.ID)
	require.Equal(t, user.FirstName, u.FirstName)
	require.Equal(t, user.LastName, u.LastName)
	require.Equal(t, user.UserName, u.UserName)
	require.Equal(t, user.Password, u.Password)
	require.Equal(t, user.Email, u.Email)
	require.Equal(t, user.CreatedAt, u.CreatedAt)
	require.Equal(t, user.DeletedAt, u.DeletedAt)

}
