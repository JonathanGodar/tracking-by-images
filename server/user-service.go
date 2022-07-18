package main

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/JonathanGodar/go-web-gin/models"
	"github.com/JonathanGodar/go-web-gin/server/myoto"
	"github.com/golang-jwt/jwt"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	db     *sql.DB
	logger *zap.Logger
}

func (s *userService) Initialize(db *sql.DB, logger *zap.Logger) {
	s.db = db
	s.logger = logger
}

type UserClaims struct {
	User models.User `json:"user"`

	jwt.StandardClaims
}

var secret = []byte("dev secret")

func (s userService) AddUser(ctx context.Context, req myoto.AddUserRequest) (*myoto.AddUserResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("Email and password must be set")
	}

	var user models.User
	user.Email = req.Email
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)

	if err != nil {
		return nil, err
	}
	user.PasswordHash = string(hash)
	err = user.Insert(ctx, s.db, boil.Infer())

	return &myoto.AddUserResponse{
		User: user,
	}, err
}

func (s userService) GetAccessToken(ctx context.Context, req myoto.GetAccessTokenRequest) (*myoto.GetAccessTokenResponse, error) {
	user, err := models.Users(qm.Where("email = ?", req.Email)).One(ctx, s.db)

	unauthErr := errors.New("Unauthorized")
	if err != nil {
		return nil, unauthErr
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, unauthErr
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		User: *user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return nil, err
	}

	return &myoto.GetAccessTokenResponse{
		Token: tokenString,
	}, nil
}

func (s userService) GetMe(ctx context.Context, _ interface{}) (*myoto.GetMeResponse, error) {
	user, ok := ctx.Value(SignedInUserKey).(models.User)

	if !ok {
		s.logger.Debug("Non authenticated user tried GetMe")
		return nil, errors.New("Unauthorized")
	}

	fetchedUser, err := models.FindUser(ctx, s.db, user.ID)
	if err != nil {
		return nil, err
	}

	return &myoto.GetMeResponse{
		User: *fetchedUser,
	}, nil
}
