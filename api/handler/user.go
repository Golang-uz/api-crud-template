package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/realtemirov/api-crud-template/models"
	"github.com/realtemirov/api-crud-template/services"
	"github.com/rs/zerolog"
)

type userHandler struct {
	services services.UserService
	log      zerolog.Logger
}

func NewUserHandler(services services.UserService, log zerolog.Logger) *userHandler {
	return &userHandler{
		services: services,
		log:      log,
	}
}

// Create User godoc
// @ID create_user
// @Router /user/ [POST]
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.UserRequest true "user"
// @Success 201 {object} models.Response{data=object} "user created"
// @Response 400 {object} models.Response{data=object} "Bad Request"
func (u *userHandler) CreateUser(c *gin.Context) {

	// get model from request
	var user *models.User

	// checking error with binding json
	if err := c.ShouldBindJSON(&user); err != nil {

		// if error
		response(c, http.StatusBadRequest, err, nil, ErrorBadRequest)
		return
	}

	// create user
	user, err := u.services.CreateUser(c, user)

	// checking error
	if err != nil {
		// if error
		response(c, http.StatusBadRequest, err, nil, ErrorBadRequest)
		return
	}

	// return user
	response(c, http.StatusCreated, nil, user, Success)
}

// GetUserByID User godoc
// @ID get_user_by_id
// @Router /user/id/{id} [GET]
// @Summary GetUserByID User
// @Description GetUserByID User
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} models.Response{data=object} "user"
// @Response 400 {object} models.Response{data=object} "Bad Request"
// @Response 404 {object} models.Response{data=object} "Not Found"
func (u *userHandler) GetUserByID(c *gin.Context) {

	// get id from url
	id := c.Param("id")

	// checking id
	if id == "" {
		// if id is empty
		response(c, http.StatusBadRequest, errors.New(InvalidID), nil, ErrorBadRequest)
		return
	}

	// convert id to int
	int_id, err := strconv.Atoi(id)

	// checking error
	if err != nil {
		// id not valid
		response(c, http.StatusNotFound, err, nil, ErrorBadRequest)
		return
	}

	// get user by id from db
	user, err := u.services.GetUserByID(c, int_id)

	// checking error
	if err != nil {
		// if user not found
		if err == sql.ErrNoRows {
			err = fmt.Errorf(ErrorNotFound)
		}
		response(c, http.StatusNotFound, err, nil, ErrorNotFound)
		return
	}

	// return user
	response(c, http.StatusOK, nil, user, Success)
}

// GetUserByEmail User godoc
// @ID get_user_by_email
// @Router /user/email/{email} [GET]
// @Summary GetUserByEmail User
// @Description GetUserByEmail User
// @Tags User
// @Accept json
// @Produce json
// @Param email path string true "email"
// @Success 200 {object} models.Response{data=object} "user"
// @Response 400 {object} models.Response{data=object} "Bad Request"
// @Response 404 {object} models.Response{data=object} "Not Found"
func (u *userHandler) GetUserByEmail(c *gin.Context) {

	// get email from url
	email := c.Param("email")

	// checking email
	if email == "" {
		// if id is empty
		response(c, http.StatusBadRequest, errors.New(InvalidEMail), nil, ErrorBadRequest)
		return
	}

	// get user by email from db
	user, err := u.services.GetUserByEmail(c, email)

	// checking error
	if err != nil {

		// if user not found
		if err == sql.ErrNoRows {
			err = fmt.Errorf(ErrorNotFound)
		}
		response(c, http.StatusNotFound, err, nil, ErrorNotFound)
		return
	}

	// return user
	response(c, http.StatusOK, nil, user, Success)
}

// GetUserByUserName User godoc
// @ID get_user_by_username
// @Router /user/username/{username} [GET]
// @Summary GetUserByUserName User
// @Description GetUserByUserName User
// @Tags User
// @Accept json
// @Produce json
// @Param username path string true "username"
// @Success 200 {object} models.Response{data=object} "user"
// @Response 400 {object} models.Response{data=object} "Bad Request"
// @Response 404 {object} models.Response{data=object} "Not Found"
func (u *userHandler) GetUserByUserName(c *gin.Context) {

	// get username from url
	username := c.Param("username")

	// checking username
	if username == "" {
		// if id is empty
		response(c, http.StatusBadRequest, errors.New(InvalidEMail), nil, ErrorBadRequest)
		return
	}

	// get user by user name from db
	user, err := u.services.GetUserByUserName(c, username)

	// checking error
	if err != nil {

		// if user not found
		if err == sql.ErrNoRows {
			err = fmt.Errorf(ErrorNotFound)
		}
		response(c, http.StatusNotFound, err, nil, ErrorNotFound)
		return
	}

	// return user
	response(c, http.StatusOK, nil, user, Success)
}

// GetUsers User godoc
// @ID get_users
// @Router /user/ [GET]
// @Summary GetUsers User
// @Description GetUsers User
// @Tags User
// @Accept json
// @Produce json
// @Param page query int false "page"
// @Param per_page query int false "per_page"
// @Success 200 {object} models.Response{data=object} "Users"
// @Response 400 {object} models.Response{data=object} "Bad Request"
// @Response 404 {object} models.Response{data=object} "Not Found"
func (u *userHandler) GetUsers(c *gin.Context) {

	var (
		page    string
		perPage string
	)

	// get page from url
	page, ok := c.GetQuery("page")

	// checking page
	if !ok {
		// if page is empty set default page
		page = "1"
	}

	// get per page from url
	perPage, ok = c.GetQuery("per_page")

	// checking per page
	if !ok {
		// if per page is empty set default per page
		perPage = "10"
	}

	// convert page to int
	int_page, err := strconv.Atoi(page)

	// checking error
	if err != nil {
		// if error
		int_page = 1
	}

	// convert per page to int
	int_perPage, err := strconv.Atoi(perPage)

	// checking error
	if err != nil {
		// if error
		int_perPage = 10
	}

	// create meta
	meta := models.Meta{
		PerPage:     int_perPage,
		CurrentPage: int_page,
	}

	// get users from db
	users, err := u.services.GetUsers(c, &meta)

	// checking error
	if err != nil {
		// if error
		response(c, http.StatusBadRequest, err, nil, ErrorBadRequest)
		return
	}

	// return users
	response(c, http.StatusOK, nil, users, Success)
}

// UpdateUser User godoc
// @ID update_user
// @Router /user/{id} [PUT]
// @Summary UpdateUser User
// @Description UpdateUser User
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param user body models.UserRequest true "user"
// @Success 200 {object} models.Response{data=object} "User"
// @Response 400 {object} models.Response{data=object} "Bad Request"
// @Response 404 {object} models.Response{data=object} "Not Found"
func (u *userHandler) UpdateUser(c *gin.Context) {

	// get id from url
	id := c.Param("id")

	// checking id
	if id == "" {
		// if id is empty
		response(c, http.StatusBadRequest, errors.New(InvalidID), nil, ErrorBadRequest)
		return
	}

	// convert id to int
	int_id, err := strconv.Atoi(id)

	// checking error
	if err != nil {
		// if user not found
		response(c, http.StatusBadRequest, err, nil, ErrorBadRequest)
		return
	}

	// get model from request
	var user *models.User

	// checking error with binding json
	if err := c.ShouldBindJSON(&user); err != nil {

		// if error
		response(c, http.StatusBadRequest, err, nil, ErrorBadRequest)
		return
	}

	user.ID = int_id

	// update user
	user, err = u.services.UpdateUser(c, user)

	// checking error
	if err != nil {

		// if user not found
		if err == sql.ErrNoRows {
			err = fmt.Errorf(ErrorNotFound)
		}
		response(c, http.StatusNotFound, err, nil, ErrorNotFound)
		return
	}

	// return user
	response(c, http.StatusOK, nil, user, Success)
}

// DeleteUser User godoc
// @ID delete_user
// @Router /user/{id} [DELETE]
// @Summary DeleteUser User
// @Description DeleteUser User
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} models.Response{data=object} "User"
// @Response 400 {object} models.Response{data=object} "Bad Request"
// @Response 404 {object} models.Response{data=object} "Not Found"
func (u *userHandler) DeleteUser(c *gin.Context) {

	// get id from url
	id := c.Param("id")

	// checking id
	if id == "" {
		// if id is empty
		response(c, http.StatusBadRequest, errors.New(InvalidID), nil, ErrorBadRequest)
		return
	}

	// convert id to int
	int_id, err := strconv.Atoi(id)

	// checking error
	if err != nil {
		// if user not found
		response(c, http.StatusBadRequest, err, nil, ErrorBadRequest)
		return
	}

	// delete user
	user, err := u.services.DeleteUser(c, int_id)

	// checking error
	if err != nil {

		// if user not found
		if err == sql.ErrNoRows {
			err = fmt.Errorf(ErrorNotFound)
		}
		response(c, http.StatusNotFound, err, nil, ErrorNotFound)
		return
	}

	// return user
	response(c, http.StatusOK, nil, user, Success)

}
