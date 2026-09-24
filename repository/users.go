package repository

import (
	"context"
	"log"
	"lpnapi/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	DB *pgxpool.Pool
}

func (r *UserRepository) CreateUser(user *model.User) (*model.User, error) {
	var createdUser *model.User
	query := `insert into users() values () on conflict do nothing returning user`
	err := r.DB.QueryRow(context.Background(), query).Scan(&user)
	if err != nil {
		log.Printf("Error creating new user record: %s", err.Error())
		return &model.User{}, err
	}
	return createdUser, err
}

func (r *UserRepository) FetchUser(userId string)

func (r *UserRepository) UpdateUser(user *model.User)

func (r *UserRepository) DeleteUser(userId string)
