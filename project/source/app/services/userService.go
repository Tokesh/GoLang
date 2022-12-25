package services

import "projectGoLang/source/domain/entity"

func (s Service) SignUpService(user entity.User) error {
	return s.CreateUser(user)
}

func (s Service) FindUserId(user entity.User) entity.User {
	return s.FindUserID(user)
}

func (s Service) FindUserPass(user entity.User) entity.User {
	return s.FindUserPassword(user)
}

func (s Service) FindUserByIdService(user_id int) (entity.User, error) {
	return s.FindUserByID(user_id)
}
