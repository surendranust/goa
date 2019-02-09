// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// secured_service gRPC client
//
// Command:
// $ goa gen goa.design/goa/examples/security/design -o
// $(GOPATH)/src/goa.design/goa/examples/security

package client

import (
	"context"

	goa "goa.design/goa"
	secured_servicepb "goa.design/goa/examples/security/gen/grpc/secured_service/pb"
	goagrpc "goa.design/goa/grpc"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli secured_servicepb.SecuredServiceClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the secured_service service
// servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: secured_servicepb.NewSecuredServiceClient(cc),
		opts:    opts,
	}
}

// Signin calls the "Signin" function in secured_servicepb.SecuredServiceClient
// interface.
func (c *Client) Signin() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildSigninFunc(c.grpccli, c.opts...),
			EncodeSigninRequest,
			DecodeSigninResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goagrpc.DecodeError(err)
		}
		return res, nil
	}
}

// Secure calls the "Secure" function in secured_servicepb.SecuredServiceClient
// interface.
func (c *Client) Secure() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildSecureFunc(c.grpccli, c.opts...),
			EncodeSecureRequest,
			DecodeSecureResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goagrpc.DecodeError(err)
		}
		return res, nil
	}
}

// DoublySecure calls the "DoublySecure" function in
// secured_servicepb.SecuredServiceClient interface.
func (c *Client) DoublySecure() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildDoublySecureFunc(c.grpccli, c.opts...),
			EncodeDoublySecureRequest,
			DecodeDoublySecureResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goagrpc.DecodeError(err)
		}
		return res, nil
	}
}

// AlsoDoublySecure calls the "AlsoDoublySecure" function in
// secured_servicepb.SecuredServiceClient interface.
func (c *Client) AlsoDoublySecure() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildAlsoDoublySecureFunc(c.grpccli, c.opts...),
			EncodeAlsoDoublySecureRequest,
			DecodeAlsoDoublySecureResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goagrpc.DecodeError(err)
		}
		return res, nil
	}
}
