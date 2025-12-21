package application

import (
	"estsoftwareoficial/src/users/domain"
	"estsoftwareoficial/src/users/domain/entities"
	"time"
)

type OAuthService struct {
	userRepo domain.UserRepository
}

func NewOAuthService(userRepo domain.UserRepository) *OAuthService {
	return &OAuthService{userRepo: userRepo}
}

func (os *OAuthService) FindOrCreateOAuthUser(email, provider, oauthID, firstName, lastName string) (*entities.User, error) {
	user, err := os.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	newUser := &entities.User{
		FirstName:        firstName,
		LastName:         lastName,
		Email:            email,
		RegistrationDate: time.Now(),
		RoleID:           3,
		OAuthProvider:    &provider,
		OAuthID:          &oauthID,
	}

	return os.userRepo.Save(newUser)
}
