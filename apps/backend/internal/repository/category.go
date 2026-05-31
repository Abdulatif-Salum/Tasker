package repository

import "github.com/sriniously/go-tasker/internal/server"

type CategoryRepository struct{
	  server *server.Server
}

func NewCategoryRepository(server *server.Server) *CategoryRepository{
	  return &CategoryRepository{server: server}
}