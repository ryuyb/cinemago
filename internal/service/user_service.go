package service

import (
	"cinemago/internal/model/dto"
	"cinemago/internal/model/ent"
	"cinemago/internal/pkg/utils"
	"cinemago/internal/repository"
	utils2 "cinemago/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/jinzhu/copier"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// CreateUser Create new user
func (s *UserService) CreateUser(req dto.CreateUserReq) error {
	err := utils.ParseValidatorError(validator.New().Struct(req))
	if err != nil {
		return err
	}
	if err = s.validateEmail(req.Email); err != nil {
		return err
	}
	err = s.checkUsernameAndEmailExists(req.Username, req.Email, nil)
	if err != nil {
		return err
	}
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return err
	}
	req.Password = hashedPassword
	return s.userRepo.Create(req)
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(req dto.UpdateUserReq) error {
	err := utils.ParseValidatorError(validator.New().Struct(req))
	if err != nil {
		return err
	}
	if err = s.validateEmail(req.Email); err != nil {
		return err
	}
	err = s.checkUsernameAndEmailExists(req.Username, req.Email, &req.Id)
	if err != nil {
		return err
	}
	if req.Password != "" {
		hashedPassword, err := hashPassword(req.Password)
		if err != nil {
			return err
		}
		req.Password = hashedPassword
	}
	return s.userRepo.Update(req)
}

func (s *UserService) checkUsernameAndEmailExists(username, email string, id *int) error {
	existingUser, err := s.userRepo.FindByUsername(username)
	if err == nil && existingUser != nil && (id == nil || existingUser.ID != *id) {
		return dto.BadRequest("username already exists")
	}
	if email != "" {
		existingUser, err = s.userRepo.FindByEmail(email)
		if err == nil && existingUser != nil && (id == nil || existingUser.ID != *id) {
			return dto.BadRequest("email already exists")
		}
	}
	return nil
}

func (s *UserService) validateEmail(email string) error {
	if email == "" {
		return nil
	}
	return utils.ParseValidatorError(validator.New().Var(email, "email"))
}

// GetUserByID Get user information based on ID
func (s *UserService) GetUserByID(id int) (*dto.UserResp, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	resp := &dto.UserResp{}
	err = copier.Copy(resp, user)

	if err != nil {
		log.Errorw("Copy User to UserResp failed", "id", id, "err", err, "user", user, "resp", resp)
		return nil, err
	}

	return resp, nil
}

// DeleteUser Delete user by id
func (s *UserService) DeleteUser(id int) error {
	_, err := s.userRepo.FindById(id)
	if err != nil {
		return err
	}
	return s.userRepo.DeleteById(id)
}

func (s *UserService) CheckPassword(req dto.UserLoginReq) (*ent.User, error) {
	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil || !utils2.CheckPasswordHash(req.Password, user.Password) {
		return nil, err
	}
	return user, nil
}

func hashPassword(password string) (string, error) {
	result, err := utils2.HashPassword(password)
	if err != nil {
		return "", dto.NewErrorResponse(fiber.StatusInternalServerError, err.Error())
	}
	return result, err
}
