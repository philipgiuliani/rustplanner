package controllers

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
)

// Action defines a standard function signature for us to use when creating
// controller actions. A controller action is basically just a method attached to
// a controller.
type Action func(w http.ResponseWriter, r *http.Request, _ httprouter.Params)

// AppController is the Base controller of the App
type AppController struct{}