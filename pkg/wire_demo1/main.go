package main

import "fmt"

type config struct {
	dbAddr string
}

type db struct {
}

type application struct {
	db *db
	config *config
}

func (app *application) start()  {
	fmt.Println("app start.")
}

func main()  {
	// 不使用 wire
	//config := newConfig("./configs")
	//db := newDB(config)
	//app := newApplication(config, db)
	//app.start()

	// 使用 wire
	app := InitApplication("./configs")
	app.start()
}