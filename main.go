package main

import (
	"log"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	"github.com/philipgiuliani/rustplanner/controllers"
	"github.com/unrolled/render"
	"upper.io/db"
	"upper.io/db/sqlite"
)

func main() {
	// render, middleware
	r := render.New(render.Options{})

	n := negroni.New(
		negroni.NewLogger(),
		negroni.NewRecovery(),
	)

	// database
	dbSettings := sqlite.ConnectionURL{
		Database: "test.db",
	}

	sess, err := db.Open(sqlite.Adapter, dbSettings)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	// router
	router := httprouter.New()

	groupsController := &controllers.GroupsController{Render: r, DB: sess}
	router.GET("/api/groups.json", groupsController.Index)
	router.POST("/api/groups.json", groupsController.Create)

	n.UseHandler(router)
	n.Run(":3000")
}
