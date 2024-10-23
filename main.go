package main

import (
	"github.com/labstack/echo/v4"
	"go-ddd-sample/config"
	"go-ddd-sample/infra"
	"go-ddd-sample/interface/handler"
	"go-ddd-sample/usecase"
)

func main() {
	tr := infra.NewTaskRepository(config.NewDB())
	tu := usecase.NewTaskUsecase(tr)
	th := handler.NewTaskHandler(tu)

	e := echo.New()
	handler.InitRouting(e, th)
	e.Logger.Fatal(e.Start(":1323"))
}
