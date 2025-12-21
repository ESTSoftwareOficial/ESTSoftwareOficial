package application

import (
	"estsoftwareoficial/src/users/domain"
)

type GetTotalUsers struct {
	userRepo domain.UserRepository
}

func NewGetTotalUsers(userRepo domain.UserRepository) *GetTotalUsers {
	return &GetTotalUsers{userRepo: userRepo}
}

func (gt *GetTotalUsers) Execute() (int, error) {
	return gt.userRepo.GetTotal()
}
