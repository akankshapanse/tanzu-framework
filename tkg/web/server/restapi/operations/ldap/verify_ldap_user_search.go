// Code generated by go-swagger; DO NOT EDIT.

package ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// VerifyLdapUserSearchHandlerFunc turns a function with the right signature into a verify ldap user search handler
type VerifyLdapUserSearchHandlerFunc func(VerifyLdapUserSearchParams) middleware.Responder

// Handle executing the request and returning a response
func (fn VerifyLdapUserSearchHandlerFunc) Handle(params VerifyLdapUserSearchParams) middleware.Responder {
	return fn(params)
}

// VerifyLdapUserSearchHandler interface for that can handle valid verify ldap user search params
type VerifyLdapUserSearchHandler interface {
	Handle(VerifyLdapUserSearchParams) middleware.Responder
}

// NewVerifyLdapUserSearch creates a new http.Handler for the verify ldap user search operation
func NewVerifyLdapUserSearch(ctx *middleware.Context, handler VerifyLdapUserSearchHandler) *VerifyLdapUserSearch {
	return &VerifyLdapUserSearch{Context: ctx, Handler: handler}
}

/*
VerifyLdapUserSearch swagger:route POST /api/ldap/users/search ldap verifyLdapUserSearch

Validate LDAP User Search configuration
*/
type VerifyLdapUserSearch struct {
	Context *middleware.Context
	Handler VerifyLdapUserSearchHandler
}

func (o *VerifyLdapUserSearch) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewVerifyLdapUserSearchParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
