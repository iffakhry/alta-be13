package entities

import "time"

type UserCore struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
}

type UserResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func UserCoreToResponse(dataCore UserCore) UserResponse {
	return UserResponse{
		ID:      dataCore.ID,
		Name:    dataCore.Name,
		Email:   dataCore.Email,
		Phone:   dataCore.Phone,
		Address: dataCore.Address,
	}
}

func ListUserCoreToResponse(dataCore []UserCore) []UserResponse {
	var dataResponse []UserResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, UserCoreToResponse(v))
	}
	return dataResponse
}
