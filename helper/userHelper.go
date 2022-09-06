package helper

import (
	"golang_bootstrap_project/model"
)

type UserHelper interface {
	TransformToUserEntitlementDTO(user model.UserDao) *model.UserDTO
}

type UserHelperImpl struct {
}

func (r UserHelperImpl) TransformToUserEntitlementDTO(user model.UserDao) *model.UserDTO {
	userDTO := &model.UserDTO{}
	userDTO.UserId = user.UserId
	userDTO.Id = user.Id
	userDTO.FirstName = user.FirstName
	userDTO.LastName = user.LastName
	userDTO.Email = user.Email
	return userDTO
}
