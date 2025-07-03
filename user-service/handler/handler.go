package handler

import (
	"context"
	// "fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/saloni111/RealTimeDocColabPlatform/user-service/model"
	pb "github.com/saloni111/RealTimeDocColabPlatform/user-service/proto"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Server struct {
	pb.UnimplementedUserServiceServer
	UserModel *model.UserModel
}

type UserClaims struct {
	UserID string
	jwt.StandardClaims
}

var jwtKey = []byte("my_secret_key_for_the_sake_of_simplicity_its_just_a_long_string")

func (s *Server) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	userId := uuid.New().String()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := model.User{
		UserID:   userId,
		Email:    req.Email,
		Password: string(hashedPassword),
		UserName: req.Name,
	}

	err = s.UserModel.CreateUser(ctx, &user)

	if err != nil {
		return nil, err
	}

	return &pb.RegisterUserResponse{UserId: userId}, nil
}

func (s *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	user, err := s.UserModel.GetUserByEmail(ctx, req.Email)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if err != nil {
		return nil, err
	}

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &UserClaims{
		UserID: user.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return nil, err
	}

	return &pb.LoginUserResponse{Token: tokenString}, nil
}

func (s *Server) GetUserProfile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	user, err := s.UserModel.GetUserById(ctx, req.UserId)

	if err != nil {
		return nil, err
	}

	return &pb.GetUserProfileResponse{
		UserId: user.UserID,
		Email:  user.Email,
		Name:   user.UserName,
	}, nil
}
