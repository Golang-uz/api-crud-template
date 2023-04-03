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

type postHandler struct {
	services *services.PostService
	log      zerolog.Logger
}

func NewPostHandler(services *services.PostService, log zerolog.Logger) *postHandler {
	return &postHandler{
		services: services,
		log:      log,
	}
}

// Create Post godoc
// @ID create_post
// @Router /post/ [POST]
// @Summary Create Post
// @Description Create Post
// @Tags Post
// @Accept json
// @Produce json
// @Param post body models.PostRequest true "post"
// @Success 201 {object} models.Response{data=object} "post created"
// @Response 400 {object} models.Response{data=object} "Bad Request"
func (p *postHandler) CreatePost(c *gin.Context) {

	// get model from request
	var post *models.Post

	// checking error with binding json
	if err := c.ShouldBindJSON(&post); err != nil {

		// if error
		response(c, http.StatusBadRequest, err, nil, ErrorBadRequest)
		return
	}

	// create post
	post, err := p.services.CreatePost(c, post)

	// checking error
	if err != nil {
		// if error
		response(c, http.StatusBadRequest, err, nil, ErrorBadRequest)
		return
	}

	// return post
	response(c, http.StatusCreated, nil, post, Success)
}

// GetPostByID Post godoc
// @ID get_post_by_id
// @Router /post/{id} [GET]
// @Summary GetPostByID Post
// @Description GetPostByID Post
// @Tags Post
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} models.Response{data=object} "Post"
// @Response 400 {object} models.Response{data=object} "Bad Request"
// @Response 404 {object} models.Response{data=object} "Not Found"
func (p *postHandler) GetPostByID(c *gin.Context) {

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

	// get post by id from db
	post, err := p.services.GetPostByID(c, int_id)

	// checking error
	if err != nil {
		// if post not found
		if err == sql.ErrNoRows {
			err = fmt.Errorf(ErrorNotFound)
		}
		response(c, http.StatusNotFound, err, nil, ErrorNotFound)
		return
	}

	// return post
	response(c, http.StatusOK, nil, post, Success)
}

// GetPostByUserID Post godoc
// @ID get_post_by_user_id
// @Router /post/user/{id} [GET]
// @Summary GetPostByUserID Post
// @Description GetPostByUserID Post
// @Tags Post
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param page query int false "page"
// @Param per_page query int false "per_page"
// @Success 200 {object} models.Response{data=object} "Posts"
// @Response 400 {object} models.Response{data=object} "Bad Request"
// @Response 404 {object} models.Response{data=object} "Not Found"
func (p *postHandler) GetPostByUserID(c *gin.Context) {

	var (
		page    string
		perPage string
	)

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

	// get posts by id from db
	posts, err := p.services.GetPostByUserID(c, int_id, &meta)

	// checking error
	if err != nil {
		// if posts not found
		if err == sql.ErrNoRows {
			err = fmt.Errorf(ErrorNotFound)
		}
		response(c, http.StatusNotFound, err, nil, ErrorNotFound)
		return
	}

	// return post
	response(c, http.StatusOK, nil, posts, Success)
}

// GetPosts Post godoc
// @ID get_posts
// @Router /post/ [GET]
// @Summary GetPosts Post
// @Description GetPosts Post
// @Tags Post
// @Accept json
// @Produce json
// @Param page query int false "page"
// @Param per_page query int false "per_page"
// @Success 200 {object} models.Response{data=object} "Posts"
// @Response 400 {object} models.Response{data=object} "Bad Request"
// @Response 404 {object} models.Response{data=object} "Not Found"
func (p *postHandler) GetPosts(c *gin.Context) {
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

	// get posts from db
	posts, err := p.services.GetPosts(c, &meta)

	// checking error
	if err != nil {
		// if error
		response(c, http.StatusBadRequest, err, nil, ErrorBadRequest)
		return
	}

	// return posts
	response(c, http.StatusOK, nil, posts, Success)
}

// UpdatePost Post godoc
// @ID update_post
// @Router /post/{id} [PUT]
// @Summary UpdatePost Post
// @Description UpdatePost Post
// @Tags Post
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param post body models.PostRequest true "post"
// @Success 200 {object} models.Response{data=object} "Post"
// @Response 400 {object} models.Response{data=object} "Bad Request"
// @Response 404 {object} models.Response{data=object} "Not Found"
func (p *postHandler) UpdatePost(c *gin.Context) {

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
		// if post not found
		response(c, http.StatusBadRequest, err, nil, ErrorBadRequest)
		return
	}

	// get model from request
	var post *models.Post

	// checking error with binding json
	if err := c.ShouldBindJSON(&post); err != nil {

		// if error
		response(c, http.StatusBadRequest, err, nil, ErrorBadRequest)
		return
	}

	post.ID = int_id

	// update post
	post, err = p.services.UpdatePost(c, post)

	// checking error
	if err != nil {
		// if post not found
		if err == sql.ErrNoRows {
			err = fmt.Errorf(ErrorNotFound)
		}
		response(c, http.StatusNotFound, err, nil, ErrorNotFound)
		return
	}

	// return post
	response(c, http.StatusOK, nil, post, Success)

}

// DeletePost Post godoc
// @ID delete_post
// @Router /post/{id} [DELETE]
// @Summary DeletePost Post
// @Description DeletePost Post
// @Tags Post
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} models.Response{data=object} "Post"
// @Response 400 {object} models.Response{data=object} "Bad Request"
// @Response 404 {object} models.Response{data=object} "Not Found"
func (p *postHandler) DeletePost(c *gin.Context) {

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
		// if post not found
		response(c, http.StatusBadRequest, err, nil, ErrorBadRequest)
		return
	}

	// delete post
	post, err := p.services.DeletePost(c, int_id)

	// checking error
	if err != nil {
		// if post not found
		if err == sql.ErrNoRows {
			err = fmt.Errorf(ErrorNotFound)
		}
		response(c, http.StatusNotFound, err, nil, ErrorNotFound)
		return
	}

	// return post
	response(c, http.StatusOK, nil, post, Success)
}
