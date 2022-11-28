package delivery

import "be13/clean/features/user"

type UserResponse struct {
	ID       uint   `json:"id" form:"id"`
	FullName string `json:"full_name" form:"full_name"`
	Email    string `json:"email" form:"email"`
}

func FromCore(data user.Core) UserResponse {
	return UserResponse{
		ID:       data.ID,
		FullName: data.Name,
		Email:    data.Email,
	}
}
