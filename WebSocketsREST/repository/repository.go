package repository

import (
	"context"
	"rest_ws/models"
)

/*  Esta es la implementacion general de la capa de base de datos , se utiliza el patron de diseño Repository
Esta implementacion es agnostica a la base de datos que se utilice, ya que solo describe los metodos que se van a utilizar
y no la logica de estos.

Para definir la logica se debe crear una implementacion especifica, la cual tiene que implementar
todos los metodos de esta interfaz Repository
*/

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error
	FindUserById(ctx context.Context, id string) (*models.User, error)
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	InsertPost(ctx context.Context, post *models.Post) error
	GetPostById(ctx context.Context, id string) (*models.Post, error)
	UpdatePost(ctx context.Context, post *models.Post) error
	DeletePost(ctx context.Context, id string) error
	ListPosts(ctx context.Context, page uint64) ([]*models.Post, error)
	Close() error
}

var implementation Repository

func SetRepository(repo Repository) {
	implementation = repo
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func FindUserById(ctx context.Context, id string) (*models.User, error) {
	return implementation.FindUserById(ctx, id)
}

func FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.FindUserByEmail(ctx, email)
}

func InsertPost(ctx context.Context, post *models.Post) error {
	return implementation.InsertPost(ctx, post)
}

func GetPostById(ctx context.Context, id string) (*models.Post, error) {
	return implementation.GetPostById(ctx, id)
}

func UpdatePost(ctx context.Context, post *models.Post) error {
	return implementation.UpdatePost(ctx, post)
}

func DeletePost(ctx context.Context, id string) error {
	return implementation.DeletePost(ctx, id)
}

func ListPosts(ctx context.Context, page uint64) ([]*models.Post, error) {
	return implementation.ListPosts(ctx, page)
}
