package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(username string, password string) (string, error) // returns JWT token
	Register(fullName string, email string, password string) (string, error) // returns JWT token

	// GenerateToken(userID string, role string) (string, error)
	// ValidateToken(tokenStr string) (string, string, error) // returns userID and role
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (s *authService) Login(email string, password string) (string, error) {
	userEmail,err := s.userRepo.GetByEmail(email)
	if err != nil {
		return err.Error(), err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userEmail.Password), []byte(password))
	if err != nil {
		return "Invalid credentials", err
	}

	if userEmail.ID == "" {
		return "User not found", nil
	}

	payload := jwt.MapClaims{
		"user_id": userEmail.ID,
		"email":   userEmail.Email,
		"exp":     jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expires in 24 hours
	}

	secretKey:=os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return "JWT secret key not configured", nil
	}

	//create token
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := jwtToken.SignedString(secretKey)
	if err != nil {
		return "Failed to generate token", err
	}
	return tokenString, nil
}

func (s *authService) Register(fullName string, email string, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "Failed to hash password", err
	}

	user := &domain.User{
		FullName: fullName,
		Email:    email,
		Password: string(hashedPassword),
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return "Failed to create user", err
	}

	return s.Login(email, password) // Auto-login after registration
}