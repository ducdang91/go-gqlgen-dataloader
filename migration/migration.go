package migration

import (
	"github.com/ducdang91/go-gqlgen-dataloader/config"
	"github.com/ducdang91/go-gqlgen-dataloader/graph/model"
)

func MigrateTable() {
	db := config.GetDB()
	db.AutoMigrate(model.User{}, model.Transaction{}, model.TransactionDetail{})
}