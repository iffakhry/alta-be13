package delivery

import "be13/clean/features/user"

type UserResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Role    string `json:"role"`
}

func fromCore(dataCore user.Core) UserResponse {
	return UserResponse{
		ID:      dataCore.ID,
		Name:    dataCore.Name,
		Email:   dataCore.Email,
		Phone:   dataCore.Phone,
		Address: dataCore.Address,
		Role:    dataCore.Role,
	}
}

func fromCoreList(dataCore []user.Core) []UserResponse {
	var dataResponse []UserResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
