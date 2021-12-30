package main

import (
	"./framework"
	"./orm"
	"net/http"
)



func main()  {

	orm.DB = orm.InitializeDB()
	defer func() {

	}()
	orm.InitTable(orm.DB)
	//orm.CreateData(db)




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