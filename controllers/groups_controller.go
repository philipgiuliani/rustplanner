package controllers

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/philipgiuliani/rustplanner/models"
	"github.com/unrolled/render"
	"upper.io/db"
)

type GroupsController struct {
	AppController
	*render.Render
	DB db.Database
}

func (c *GroupsController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var groups models.Groups

	groupsCollection, err := c.DB.Collection("groups")
	if err != nil {
		log.Fatalf("DB.Collection(): %q\n", err)
	}

	res := groupsCollection.Find()
	res.All(&groups)

	c.JSON(w, http.StatusOK, groups)
}

func (c *GroupsController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c.JSON(w, http.StatusOK, nil)
}
