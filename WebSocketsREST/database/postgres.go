package database

/*
	Implementacion especifica del repositorio para la base de datos de PostgresSQL
	Aqui se incluyen las funciones que se encargan de interactuar con la base de datos de PostgresSQL
	Tambien se crea una conexion a la base de datos y se cierra al final
*/

import (
	"context"
	"database/sql"
	"log"
	"rest_ws/models"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db: db}, nil
}

func (p *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := p.db.ExecContext(ctx, "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)",
		user.Id, user.Email, user.Password)
	return err
}

func (p *PostgresRepository) FindUserById(ctx context.Context, id string) (*models.User, error) {
	rows, err := p.db.QueryContext(ctx, "SELECT id, email FROM users WHERE id = $1", id)

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var user = models.User{}

	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email); err == nil {
			return &user, nil
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *PostgresRepository) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	rows, err := p.db.QueryContext(ctx, "SELECT id, email, password FROM users WHERE email = $1", email)

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var user = models.User{}

	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email, &user.Password); err == nil {
			return &user, nil
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *PostgresRepository) InsertPost(ctx context.Context, post *models.Post) error {
	_, err := p.db.ExecContext(ctx, "INSERT INTO posts (id, content, created_at, user_id) VALUES ($1, $2, $3, $4)",
		post.Id, post.Content, post.CreatedAt, post.UserID)
	return err
}

func (p *PostgresRepository) GetPostById(ctx context.Context, id string) (*models.Post, error) {
	rows, err := p.db.QueryContext(ctx, "SELECT id, content, created_at, user_id FROM posts WHERE id = $1", id)

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var post = models.Post{}

	for rows.Next() {
		if err = rows.Scan(&post.Id, &post.Content, &post.CreatedAt, &post.UserID); err == nil {
			return &post, nil
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &post, nil
}

func (p *PostgresRepository) UpdatePost(ctx context.Context, post *models.Post) error {
	_, err := p.db.ExecContext(ctx, "UPDATE posts SET content = $1 WHERE id = $2", post.Content, post.Id)
	return err
}

func (p *PostgresRepository) DeletePost(ctx context.Context, id string) error {
	_, err := p.db.ExecContext(ctx, "DELETE FROM posts WHERE id = $1", id)
	return err
}

func (p *PostgresRepository) ListPosts(ctx context.Context, page uint64) ([]*models.Post, error) {
	rows, err := p.db.QueryContext(ctx, "SELECT id, content, created_at, user_id FROM posts LIMIT $1 OFFSET $2", 10, page*10)

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var posts []*models.Post

	for rows.Next() {
		var post = models.Post{}
		if err = rows.Scan(&post.Id, &post.Content, &post.CreatedAt, &post.UserID); err == nil {
			posts = append(posts, &post)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *PostgresRepository) Close() error {
	return p.db.Close()
}
