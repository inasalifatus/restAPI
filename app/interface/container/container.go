package container

import (
	"restAPI/app/adapter/database"
	"restAPI/app/repository/order"
	"restAPI/app/usecase"
)

type Container struct {
	Usecase usecase.Usecase
}

func SetupContainer() Container {
	db := database.SetupDatabase().DB

	repo := order.NewOrderRepository(db)
	usecase := usecase.NewUsecase(repo)

	return Container{
		Usecase: usecase,
	}

}
