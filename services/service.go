package services

import (
	"github.com/realtemirov/api-crud-template/config"
	"github.com/realtemirov/api-crud-template/storage"
	"github.com/rs/zerolog"
)

type Services struct {
	storage     storage.StorageI
	cnf         *config.Config
	log         zerolog.Logger
	UserService UserService
	PostService *PostService
}

func NewService(cnf *config.Config, storage storage.StorageI, log zerolog.Logger) *Services {
	return &Services{
		storage:     storage,
		cnf:         cnf,
		log:         log,
		UserService: NewUserService(storage.User(), log),
		PostService: NewPostService(storage.Post(), log),
	}
}
