package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	tr := infra.NewTaskRepository(config.NewDB())
	tu := infra.NewTaskUsecase(tr)
	th := infra.NewTaskHandler(tu)

	e := echo.New()
	hander.InitRouting(e, th)
	e.Logger.Fatal(e.Start(":1323"))
}
