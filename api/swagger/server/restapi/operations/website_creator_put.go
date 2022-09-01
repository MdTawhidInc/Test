// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WebsiteCreatorPutHandlerFunc turns a function with the right signature into a website creator put handler
type WebsiteCreatorPutHandlerFunc func(WebsiteCreatorPutParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WebsiteCreatorPutHandlerFunc) Handle(params WebsiteCreatorPutParams) middleware.Responder {
	return fn(params)
}

// WebsiteCreatorPutHandler interface for that can handle valid website creator put params
type WebsiteCreatorPutHandler interface {
	Handle(WebsiteCreatorPutParams) middleware.Responder
}

// NewWebsiteCreatorPut creates a new http.Handler for the website creator put operation
func NewWebsiteCreatorPut(ctx *middleware.Context, handler WebsiteCreatorPutHandler) *WebsiteCreatorPut {
	return &WebsiteCreatorPut{Context: ctx, Handler: handler}
}

/* WebsiteCreatorPut swagger:route PUT /websiteCreator/{url} websiteCreatorPut

WebsiteCreatorPut website creator put API

*/
type WebsiteCreatorPut struct {
	Context *middleware.Context
	Handler WebsiteCreatorPutHandler
}

func (o *WebsiteCreatorPut) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewWebsiteCreatorPutParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}