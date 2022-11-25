package entities

import "time"

type UserCore struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Phone     string
	Address   string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Role     string `json:"role" form:"role"`
}

type UserResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Role    string `json:"role"`
}

func UserCoreToResponse(dataCore UserCore) UserResponse {
	return UserResponse{
		ID:      dataCore.ID,
		Name:    dataCore.Name,
		Email:   dataCore.Email,
		Phone:   dataCore.Phone,
		Address: dataCore.Address,
		Role:    dataCore.Role,
	}
}

func ListUserCoreToResponse(dataCore []UserCore) []UserResponse {
	var dataResponse []UserResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, UserCoreToResponse(v))
	}
	return dataResponse
}
