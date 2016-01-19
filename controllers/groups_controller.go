package controllers

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/philipgiuliani/rustplanner/models"
)

type GroupsController struct {
	BaseController
}

// swagger:route GET /groups Lists groups
//
// This will show all registered groups by default
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: groups
func (c *GroupsController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	groupsCollection, err := c.DB.Collection("groups")
	if err != nil {
		log.Fatalf("DB.Collection(): %q\n", err)
	}

	var groups models.Groups
	res := groupsCollection.Find()
	res.All(&groups)

	c.JSON(w, http.StatusOK, map[string]interface{}{
		"groups": groups,
	})
}

// swagger:route POST /groups Creates a new group
//
// This will create a new group
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: group
func (c *GroupsController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c.JSON(w, http.StatusOK, nil)
}
