package main

import (
	"api-server/api"
	"api-server/model"
	"api-server/repo"
)

func main() {

	migrate()

	handler := api.NewAPIHandler(8080)

	handler.Run()
}

func migrate() {
	db := repo.New()

	db.DB.AutoMigrate(&model.User{})
}
