package services

import (
	"errors"
	"rest-api/internal/models"
	"rest-api/internal/repositories"
	"rest-api/internal/utils"
	"time"

	"github.com/go-playground/validator/v10"
)

type UserService struct {
	userRepo  *repositories.UserRepository
	validator *validator.Validate
}

func NewUserService(userRepo *repositories.UserRepository, validator *validator.Validate) *UserService {
	return &UserService{
		userRepo:  userRepo,
		validator: validator,
	}
}

func (s *UserService) CreateUser(req *models.CreateUserRequest) (*models.UserResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return nil, err
	}

	// Check if user already exists
	_, err := s.userRepo.GetByEmail(req.Email)
	if err == nil {
		return nil, errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return &models.UserResponse{
		ID:        user.ID.Hex(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *UserService) Login(req *models.LoginRequest, jwtSecret string, jwtExpire int) (*models.LoginResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return nil, err
	}

	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(user.ID.Hex(), user.Email, jwtSecret, jwtExpire)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		Token: token,
		User: models.UserResponse{
			ID:        user.ID.Hex(),
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		},
	}, nil
}

func (s *UserService) GetUserByID(id string) (*models.UserResponse, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &models.UserResponse{
		ID:        user.ID.Hex(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *UserService) UpdateUser(id string, req *models.UpdateUserRequest) (*models.UserResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return nil, err
	}

	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		// Check if email is already taken by another user
		existingUser, err := s.userRepo.GetByEmail(req.Email)
		if err == nil && existingUser.ID.Hex() != id {
			return nil, errors.New("email is already taken")
		}
		user.Email = req.Email
	}

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return &models.UserResponse{
		ID:        user.ID.Hex(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *UserService) DeleteUser(id string) error {
	return s.userRepo.Delete(id)
}

func (s *UserService) GetAllUsers(page, limit int) (*models.PaginatedResponse, error) {
	page, limit = utils.GetPaginationParams(page, limit)
	offset := utils.CalculateOffset(page, limit)

	users, total, err := s.userRepo.GetAll(offset, limit)
	if err != nil {
		return nil, err
	}

	userResponses := make([]models.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = models.UserResponse{
			ID:        user.ID.Hex(),
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		}
	}

	return &models.PaginatedResponse{
		Success: true,
		Message: "Users retrieved successfully",
		Data:    userResponses,
		Page:    page,
		Limit:   limit,
		Total:   total,
	}, nil
}
