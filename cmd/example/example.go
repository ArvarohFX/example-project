package main

import (
	"github.com/12storeez/logzer"
	"github.com/ArvarohFX/example-project/config"
	"github.com/ArvarohFX/example-project/internal/app/example/service"
	"github.com/ArvarohFX/example-project/internal/app/example/usecase"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	cfg := config.GetConfig()
	logger := logzer.New(cfg.Server.Debug)

	logger.Infow("service is starting...", zap.Bool("debug", cfg.Server.Debug))

	uc := usecase.New(cfg, logger)
	svc := service.New(cfg, logger, uc)

	app := fiber.New()
	app.Get("/", svc.User().Create)
}

//type AnimalInterface interface {
//	SetName(name string)
//}
//
//type Animal struct {
//	FirstName string
//	LastName  string
//}
//
//type Rabbit struct {
//	Animal
//	PawsCount int
//}
//
//type Cat struct {
//	Animal
//	TailLength int
//}
//
//func (a *Animal) SetName(name string) {
//	a.FirstName = name
//}
//
//func main2() {
//	cat := &Cat{
//		TailLength: 14,
//	}
//
//	rabbit := &Rabbit{
//		PawsCount: 4,
//	}
//
//	cat.SetName("Bobik")
//	rabbit.SetName("Krol")
//}
