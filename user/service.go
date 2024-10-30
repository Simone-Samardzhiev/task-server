package user

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"task-server/middleware"
)

// ServiceImp is an implementation of Service.
type ServiceImp struct {
	Repository    Repository
	Authenticator middleware.Authenticator
}

// Login will check the user credentials and return an access token.
func (s *ServiceImp) Login(user *WithoutIdUser) (*string, error) {
	foundUser, err := s.Repository.GetUserByEmail(&user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrWrongCredentials
		}
		log.Printf("Error in user-ServiceImp-Login: %v", err)
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		log.Printf("Error in user-ServiceImp-Login:%v", err)
		return nil, ErrWrongCredentials
	}

	err = s.Repository.DeleteTokenByUserId(&foundUser.Id)
	if err != nil {
		log.Printf("Error in user-ServiceImp-Login:%v", err)
		return nil, err
	}

	tokenId := uuid.New()
	refreshToken, err := s.Authenticator.CreateRefreshToken(&tokenId)
	if err != nil {
		log.Printf("Error in user-ServiceImp-Login: %v", err)
		return nil, err
	}

	err = s.Repository.AddToken(&RefreshToken{
		Id:          tokenId,
		StringToken: *refreshToken,
		UserId:      foundUser.Id,
	})
	if err != nil {
		log.Printf("Error in user-ServiceImp-Login:%v", err)
		return nil, err
	}

	return refreshToken, nil
}

// Register will add a new user.
func (s *ServiceImp) Register(user *WithoutIdUser) error {
	ok, err := s.Repository.CheckEmail(&user.Email)
	if err != nil {
		log.Printf("Error in user-ServiceImp-Register: %v", err)
		return err
	}

	if !ok {
		return ErrEmailInUse
	}

	id := uuid.New()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error in user-ServiceImp-Register: %v", err)
		return err
	}

	err = s.Repository.AddUser(&User{
		Id:       id,
		Email:    user.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		log.Printf("Error in user-ServiceImp-Regitser: %v", err)
		return err
	}

	return nil
}

// RefreshTokens will use a valid refresh token to send a new access token and refresh token.
func (s *ServiceImp) RefreshTokens(tokenSting *string) (*TokenGroup, error) {
	id, err := s.Authenticator.CheckRefreshToken(tokenSting)
	if err != nil {
		log.Printf("Error in user-ServiceImp-RefreshTokens: %v", err)
		return nil, ErrInvalidToken
	}

	fetchedToken, err := s.Repository.GetTokenById(id)
	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("Error in user-ServiceImp-RefreshTokens: %v", err)
		return nil, ErrInvalidToken
	} else if err != nil {
		log.Printf("Error in user-ServiceImp-RefreshTokens: %v", err)
		return nil, err
	}

	err = s.Repository.DeleteTokenById(&fetchedToken.Id)
	if err != nil {
		log.Printf("Error in user-ServiceImp-RefreshTokens: %v", err)
		return nil, err
	}

	refreshToken, err := s.Authenticator.CreateRefreshToken(&fetchedToken.Id)
	if err != nil {
		log.Printf("Error in user-ServiceImp-RefreshTokens: %v", err)
	}

	err = s.Repository.AddToken(&RefreshToken{
		Id:          uuid.New(),
		StringToken: *refreshToken,
		UserId:      fetchedToken.Id,
	})
	if err != nil {
		log.Printf("Error in user-ServiceImp-RefreshTokens: %v", err)
		return nil, err
	}

	accessToken, err := s.Authenticator.CreateAccessToken(&fetchedToken.Id)
	if err != nil {
		log.Printf("Error in user-ServiceImp-RefreshTokens: %v", err)
		return nil, err
	}

	return &TokenGroup{
		AccessToken:  *accessToken,
		RefreshToken: *refreshToken,
	}, nil
}

func NewServiceImp(repository Repository, authenticator middleware.Authenticator) *ServiceImp {
	return &ServiceImp{
		Repository:    repository,
		Authenticator: authenticator,
	}
}
