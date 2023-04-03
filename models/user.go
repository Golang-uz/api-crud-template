package models

type User struct {
	Base
	FirstName string `json:"first_name" binding:"required" db:"first_name"`
	LastName  string `json:"last_name" binding:"required" db:"last_name"`
	UserName  string `json:"user_name" binding:"required" db:"user_name"`
	Email     string `json:"email" binding:"required" db:"email"`
	Password  string `json:"password" binding:"required" db:"password"`
}

type GetAllUsersResponse struct {
	Meta Meta    `json:"meta"`
	Data []*User `json:"data"`
}

type EmailRequest struct {
	Email string `json:"email"`
}

type UserNameRequest struct {
	UserName string `json:"user_name"`
}

type UserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
