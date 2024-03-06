package main

import (
	"log"

	restApi "github.com/MaksimUlitin"
	"github.com/MaksimUlitin/pkg/handler"
	"github.com/MaksimUlitin/pkg/repository"
	"github.com/MaksimUlitin/pkg/service"
)

func main() {

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(restApi.Server)
	if err := srv.Run("8088", handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
