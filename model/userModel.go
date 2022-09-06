package model

type UserDTO struct {
	Id        int32   `db:"id" json:"id,omitempty"`
	FirstName *string `db:"first_name" json:"firstName,omitempty"`
	LastName  *string `db:"last_name" json:"lastName,omitempty"`
	UserId    string  `db:"user_id" json:"userId,omitempty"`
	Email     *string `db:"email" json:"emailAddress,omitempty"`
}

type UserDao struct {
	Id        int32   `db:"id"`
	FirstName *string `db:"first_name" json:"firstName,omitempty"`
	LastName  *string `db:"last_name" json:"lastName,omitempty"`
	UserId    string  `db:"user_id" json:"userId,omitempty"`
	Email     *string `db:"email" json:"emailAddress,omitempty"`
}
