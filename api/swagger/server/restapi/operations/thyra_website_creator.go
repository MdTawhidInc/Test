// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ThyraWebsiteCreatorHandlerFunc turns a function with the right signature into a thyra website creator handler
type ThyraWebsiteCreatorHandlerFunc func(ThyraWebsiteCreatorParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ThyraWebsiteCreatorHandlerFunc) Handle(params ThyraWebsiteCreatorParams) middleware.Responder {
	return fn(params)
}

// ThyraWebsiteCreatorHandler interface for that can handle valid thyra website creator params
type ThyraWebsiteCreatorHandler interface {
	Handle(ThyraWebsiteCreatorParams) middleware.Responder
}

// NewThyraWebsiteCreator creates a new http.Handler for the thyra website creator operation
func NewThyraWebsiteCreator(ctx *middleware.Context, handler ThyraWebsiteCreatorHandler) *ThyraWebsiteCreator {
	return &ThyraWebsiteCreator{Context: ctx, Handler: handler}
}

/* ThyraWebsiteCreator swagger:route GET /thyra/websiteCreator/{resource} thyraWebsiteCreator

ThyraWebsiteCreator thyra website creator API

*/
type ThyraWebsiteCreator struct {
	Context *middleware.Context
	Handler ThyraWebsiteCreatorHandler
}

func (o *ThyraWebsiteCreator) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewThyraWebsiteCreatorParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}