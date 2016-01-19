package controllers

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/philipgiuliani/rustplanner/models"
	"github.com/unrolled/render"
	"upper.io/db.v2"
)

type GroupsController struct {
	BaseController
	*render.Render
	DB db.Database
}

func NewGroupsController(render *render.Render, sess db.Database) *GroupsController {
	return &GroupsController{Render: render, DB: sess}
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
	col, err := c.DB.Collection("groups")
	if err != nil {
		log.Fatalf("DB.Collection(): %q\n", err)
	}

	var groups models.Groups
	col.Find().Where("name LIKE ?", "cortex%").All(&groups)

	c.JSON(w, http.StatusOK, map[string]interface{}{
		"groups": groups,
	})
}

// swagger:route POST /groups Create group
//
// This will create a new group
//
//     Produces:
//     - application/json
//
//     Responses:
//       200: group
func (c *GroupsController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	group := &models.Group{
		Name:     "cortex2",
		Password: "new",
	}

	col, err := c.DB.Collection("groups")
	if err != nil {
		log.Fatalf("DB.Collection(): %q\n", err)
	}

	created, err := col.Append(group)
	if err != nil {
		log.Fatalf("col.Append: %q\n", err)
	}

	c.JSON(w, http.StatusOK, map[string]interface{}{
		"group": created,
	})
}
