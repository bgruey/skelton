package main

import (
	"api-server/api"
)

func main() {

	handler := api.NewAPIHandler(8080)

	handler.Run()
}
