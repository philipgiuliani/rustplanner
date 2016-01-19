package controllers

import (
	"github.com/unrolled/render"
	"upper.io/db"
)

// AppController is the Base controller of the App
type BaseController struct {
	*render.Render
	DB db.Database
}
