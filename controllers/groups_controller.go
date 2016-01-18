package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/philipgiuliani/rustplanner/models"
	"github.com/unrolled/render"
)

type GroupsController struct {
	AppController
	*render.Render
}

func (c *GroupsController) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	group := models.Group{
		Name: "cortex",
	}
	c.JSON(w, http.StatusOK, group)
}

func (c *GroupsController) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c.JSON(w, http.StatusOK, nil)
}
