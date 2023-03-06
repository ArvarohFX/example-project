package main

import (
	"fmt"
	"github.com/12storeez/logzer"
	"github.com/ArvarohFX/example-project/config"
	db "github.com/ArvarohFX/example-project/internal"
	"github.com/ArvarohFX/example-project/internal/app/example/service"
	"github.com/ArvarohFX/example-project/internal/app/example/usecase"
	"github.com/ArvarohFX/example-project/internal/repository"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func main() {
	cfg := config.GetConfig()
	logger := logzer.New(cfg.Server.Debug)

	logger.Infow("service is starting...", zap.Bool("debug", cfg.Server.Debug))

	dbStorage, err := db.NewDBStorage(cfg, logger, false)
	if err != nil {
		logger.Fatal(err)
	}

	store := repository.New(cfg, logger, dbStorage)
	uc := usecase.New(cfg, logger, store)
	svc := service.New(cfg, logger, uc)

	app := fiber.New()
	app.Post("/user", svc.User().Create)
	//app.Get("/about", svc.Metadata().About().Get)

	logger.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.Server.Port)))
}

//string
//int32
//int64
//int
//float32
//float64
//float
//bool
//rune
//byte
//
//map[string]int  {"param": 2, "param2": 4}
//[]string
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
