package main

import (
	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	"github.com/philipgiuliani/rustplanner/controllers"
	"github.com/unrolled/render"
)

func main() {
	r := render.New(render.Options{})

	n := negroni.New(
		negroni.NewLogger(),
		negroni.NewRecovery(),
	)

	router := httprouter.New()

	groupsController := &controllers.GroupsController{Render: r}
	router.GET("/api/groups.json", groupsController.Index)
	router.POST("/api/groups.json", groupsController.Create)

	n.UseHandler(router)
	n.Run(":3000")
}
