package repository

import (
	"cinemago/internal/model/dto"
	"cinemago/internal/model/ent"
	"cinemago/internal/model/ent/user"
	"cinemago/internal/pkg/utils"
	"context"
)

type UserRepository struct {
	*Client
}

func NewUserRepository(client *Client) *UserRepository {
	return &UserRepository{Client: client}
}

func (r *UserRepository) Create(userReq dto.CreateUserReq) error {
	createBuilder := r.User.Create().
		SetUsername(userReq.Username).
		SetPassword(userReq.Password).
		SetEmail(userReq.Email)
	return utils.ParseEntError(createBuilder.Exec(context.Background()))
}

func (r *UserRepository) Update(req dto.UpdateUserReq) error {
	updateBuilder := r.User.UpdateOneID(req.Id).
		SetUsername(req.Username).
		SetPassword(req.Password).
		SetEmail(req.Email)
	return utils.ParseEntError(updateBuilder.Exec(context.Background()))
}

func (r *UserRepository) FindById(id int) (*ent.User, error) {
	u, err := r.User.Get(context.Background(), id)
	if err != nil {
		return nil, utils.ParseEntError(err)
	}
	return u, nil
}

func (r *UserRepository) FindByUsername(username string) (*ent.User, error) {
	u, err := r.User.Query().Where(user.UsernameEQ(username)).Only(context.Background())
	if err != nil {
		return nil, utils.ParseEntError(err)
	}
	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*ent.User, error) {
	u, err := r.User.Query().Where(user.EmailEQ(email)).Only(context.Background())
	if err != nil {
		return nil, utils.ParseEntError(err)
	}
	return u, err
}

func (r *UserRepository) DeleteById(id int) error {
	return utils.ParseEntError(r.User.DeleteOneID(id).Exec(context.Background()))
}
