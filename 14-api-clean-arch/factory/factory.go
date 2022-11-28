package factory

import (
	userDelivery "be13/clean/features/user/delivery"
	userRepo "be13/clean/features/user/repository"
	userService "be13/clean/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userRepoFactory := userRepo.New(db)
	// userRepoFactory := userRepo.NewRaw(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory, e)

}
