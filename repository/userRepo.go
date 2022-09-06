package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"golang_bootstrap_project/model"
)

type UserRepo interface {
	FindActiveUserWithUserId(userId string) (*model.UserDTO, error)
}

type UserRepoImpl struct {
	Db *sqlx.DB
}

func (r UserRepoImpl) FindActiveUserWithUserId(userId string) (*model.UserDTO, error) {
	var users model.UserDao
	query := fmt.Sprintf(`
		SELECT u.id as id, coalesce(u.first_name, '') as first_name, coalesce(u.last_name, '') as last_name, u.user_id as user_id, 
			       coalesce(u.email, concat(u.user_id, '')) as email, 
		FROM users u
		LEFT JOIN entitlements e on u.id = e.user_id
		WHERE u.is_active=True AND LOWER(u.user_id) = LOWER($1)
		ORDER BY u.user_id`)

	err := r.Db.Select(&users, query, userId)
	if err != nil {
		return nil, err
	}

	transformedReturningUsers := r.transformToUserEntitlementDTO(users)

	return transformedReturningUsers, nil
}

func (r UserRepoImpl) transformToUserEntitlementDTO(user model.UserDao) *model.UserDTO {
	userDTO := &model.UserDTO{}
	userDTO.UserId = user.UserId
	userDTO.Id = user.Id
	userDTO.FirstName = user.FirstName
	userDTO.LastName = user.LastName
	userDTO.Email = user.Email
	return userDTO
}
