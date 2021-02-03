package controllers

import "github.com/Gys/goa"

// JsController implements the js resource.
type JsController struct {
	*goa.Controller
}

// NewJs creates a js controller.
func NewJs(service *goa.Service) *JsController {
	return &JsController{Controller: service.NewController("JsController")}
}
