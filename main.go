package main

import (
	"./config"
	"./db"

	"./framework"
	"net/http"
)

func main()  {


	dbinit := db.InitializeDB()
	defer func() {
		if dbinit != nil {
			db, _ := config.App.DB.DB()
			db.Close()
		}
	}()
	core := framework.NewCore()
	core.Use(framework.Recovery())
	core.Use(framework.Cost())
	framework.RegisterRouter(core)
	server := &http.Server{
		Addr:		":8080",
		Handler:    core,
	}

	server.ListenAndServe()
}