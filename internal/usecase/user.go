package usecase

import "vk-api/internal/entity"

type UserRepo interface {
	User(userID string) (entity.User, error)
	FriendsByID(userID string) ([]entity.User, error)
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(r UserRepo) *UserUsecase {
	return &UserUsecase{r}
}

func (u *UserUsecase) GetUser(userID string) (entity.User, error) {
	return u.repo.User(userID)
}

func (u *UserUsecase) GetUserFriendsByID(userID string) ([]entity.User, error) {
	return u.repo.FriendsByID(userID)
}
