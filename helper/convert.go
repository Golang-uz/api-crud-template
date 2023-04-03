package helper

import "github.com/realtemirov/api-crud-template/models"

func RequestToModel(request models.UserRequest, id int) *models.User {

	return &models.User{
		Base: models.Base{
			ID: id,
		},
		FirstName: request.FirstName,
		LastName:  request.LastName,
		UserName:  request.UserName,
		Email:     request.Email,
		Password:  request.Password,
	}
}
