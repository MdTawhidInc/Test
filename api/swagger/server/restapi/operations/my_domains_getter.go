// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// MyDomainsGetterHandlerFunc turns a function with the right signature into a my domains getter handler
type MyDomainsGetterHandlerFunc func(MyDomainsGetterParams) middleware.Responder

// Handle executing the request and returning a response
func (fn MyDomainsGetterHandlerFunc) Handle(params MyDomainsGetterParams) middleware.Responder {
	return fn(params)
}

// MyDomainsGetterHandler interface for that can handle valid my domains getter params
type MyDomainsGetterHandler interface {
	Handle(MyDomainsGetterParams) middleware.Responder
}

// NewMyDomainsGetter creates a new http.Handler for the my domains getter operation
func NewMyDomainsGetter(ctx *middleware.Context, handler MyDomainsGetterHandler) *MyDomainsGetter {
	return &MyDomainsGetter{Context: ctx, Handler: handler}
}

/* MyDomainsGetter swagger:route GET /my/domains/{nickname} myDomainsGetter

MyDomainsGetter my domains getter API

*/
type MyDomainsGetter struct {
	Context *middleware.Context
	Handler MyDomainsGetterHandler
}

func (o *MyDomainsGetter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewMyDomainsGetterParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}