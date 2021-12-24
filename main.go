package main

import (
	"./framework"
	"net/http"
)

func main()  {
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