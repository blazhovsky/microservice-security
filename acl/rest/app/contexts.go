// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "acl-manager": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/JormungandrK/microservice-security/acl/rest
// --out=$(GOPATH)/src/github.com/JormungandrK/microservice-security/acl/rest
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// CreatePolicyAclContext provides the acl createPolicy action context.
type CreatePolicyAclContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *ACLPolicyPayload
}

// NewCreatePolicyAclContext parses the incoming request URL and body, performs validations and creates the
// context used by the acl controller createPolicy action.
func NewCreatePolicyAclContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreatePolicyAclContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreatePolicyAclContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *CreatePolicyAclContext) Created(r *ACLPolicy) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/jormungandr-acl-policy+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 201, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreatePolicyAclContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreatePolicyAclContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// DeletePolicyAclContext provides the acl deletePolicy action context.
type DeletePolicyAclContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	PolicyID string
}

// NewDeletePolicyAclContext parses the incoming request URL and body, performs validations and creates the
// context used by the acl controller deletePolicy action.
func NewDeletePolicyAclContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeletePolicyAclContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeletePolicyAclContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramPolicyID := req.Params["policyId"]
	if len(paramPolicyID) > 0 {
		rawPolicyID := paramPolicyID[0]
		rctx.PolicyID = rawPolicyID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *DeletePolicyAclContext) OK(r *ACLPolicy) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/jormungandr-acl-policy+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeletePolicyAclContext) NotFound(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *DeletePolicyAclContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// GetAclContext provides the acl get action context.
type GetAclContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	PolicyID string
}

// NewGetAclContext parses the incoming request URL and body, performs validations and creates the
// context used by the acl controller get action.
func NewGetAclContext(ctx context.Context, r *http.Request, service *goa.Service) (*GetAclContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := GetAclContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramPolicyID := req.Params["policyId"]
	if len(paramPolicyID) > 0 {
		rawPolicyID := paramPolicyID[0]
		rctx.PolicyID = rawPolicyID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *GetAclContext) OK(r *ACLPolicy) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/jormungandr-acl-policy+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *GetAclContext) NotFound(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *GetAclContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// ManageAccessAclContext provides the acl manage-access action context.
type ManageAccessAclContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *AccessPolicyPayload
}

// NewManageAccessAclContext parses the incoming request URL and body, performs validations and creates the
// context used by the acl controller manage-access action.
func NewManageAccessAclContext(ctx context.Context, r *http.Request, service *goa.Service) (*ManageAccessAclContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ManageAccessAclContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ManageAccessAclContext) OK(r *ACLPolicy) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/jormungandr-acl-policy+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ManageAccessAclContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ManageAccessAclContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// UpdatePolicyAclContext provides the acl updatePolicy action context.
type UpdatePolicyAclContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	PolicyID string
	Payload  *ACLPolicyPayload
}

// NewUpdatePolicyAclContext parses the incoming request URL and body, performs validations and creates the
// context used by the acl controller updatePolicy action.
func NewUpdatePolicyAclContext(ctx context.Context, r *http.Request, service *goa.Service) (*UpdatePolicyAclContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UpdatePolicyAclContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramPolicyID := req.Params["policyId"]
	if len(paramPolicyID) > 0 {
		rawPolicyID := paramPolicyID[0]
		rctx.PolicyID = rawPolicyID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdatePolicyAclContext) OK(r *ACLPolicy) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/jormungandr-acl-policy+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *UpdatePolicyAclContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdatePolicyAclContext) NotFound(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *UpdatePolicyAclContext) InternalServerError(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}
