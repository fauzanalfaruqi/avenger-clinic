package router

import (
	"avengers-clinic/src/action/actionDelivery"
	"avengers-clinic/src/action/actionRepository"
	"avengers-clinic/src/action/actionUsecase"
	"avengers-clinic/src/user/userDelivery"
	"avengers-clinic/src/user/userRepository"
	"avengers-clinic/src/user/userUsecase"
	"avengers-clinic/src/medicine/delivery"
	"avengers-clinic/src/medicine/repository"
	"avengers-clinic/src/medicine/usecase"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1Group *gin.RouterGroup, db *sql.DB) {
	userRepository := userRepository.NewUserRepository(db)
	userUsecase := userUsecase.NewUserUsecase(userRepository)
	userDelivery.NewUserDelivery(v1Group, userUsecase)

	actionRepository := actionRepository.NewActionRepository(db)
	actionUsecase := actionUsecase.NewActionUsecase(actionRepository)
	actionDelivery.NewActionDelivery(v1Group, actionUsecase)

	medicineR := repository.NewMedicineRepository(db)
	medicineUC := usecase.NewMedicineUsecase(medicineR)
	delivery.NewMedicineDelivery(v1Group, medicineUC)
}
