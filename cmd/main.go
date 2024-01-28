package main

import (
	"database/sql"
	"dontpad-storage/internal/api"
	"dontpad-storage/internal/dontpad"
	"dontpad-storage/internal/files"
	"dontpad-storage/internal/user"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

func main() {
	r := gin.Default()

	db, err := sql.Open("sqlite3", "./resources/data.db")
	if err != nil {
		panic(err)
	}

	clientHttp := http.Client{}

	encrypter := files.NewEncrypter()
	userData := user.NewData(db)
	dontPad := dontpad.NewDontPad(&clientHttp)
	processor := files.NewProcessor(encrypter, dontPad, userData)
	fileHandler := files.NewHandler(processor)
	routes := api.NewRoutes(r, fileHandler)

	routes.SetupRoutes()
	if err != nil {
		panic(err)
	}

	err = r.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
