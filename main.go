// Package main RustPlanner API.
//
//     Schemes: http
//     Host: localhost
//     BasePath: /api/v1
//     Version: 0.0.1
//     Contact: Philip Giuliani<info@philipg.net> http://philipg.net
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"log"
	"os"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	"github.com/philipgiuliani/rustplanner/controllers"
	"github.com/unrolled/render"
	"upper.io/db.v2"
	"upper.io/db.v2/sqlite"
)

var dbSettings = sqlite.ConnectionURL{
	Database: "test.db",
}

func main() {
	// render
	r := render.New(render.Options{})

	// middleware
	n := negroni.New(
		negroni.NewLogger(),
		negroni.NewRecovery(),
	)

	// database
	sess, err := db.Open(sqlite.Adapter, dbSettings)
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	// router
	router := httprouter.New()

	groupsController := controllers.NewGroupsController(r, sess)
	router.GET("/api/v1/groups.json", groupsController.Index)
	router.POST("/api/v1/groups.json", groupsController.Create)

	n.UseHandler(router)
	n.Run(":" + strconv.Itoa(getPort()))
}

func getPort() int {
	defaultPort := 3000
	envPort := os.Getenv("PORT")
	if envPort == "" {
		return defaultPort
	}

	port, err := strconv.Atoi(envPort)
	if err != nil {
		return defaultPort
	}
	return port
}
