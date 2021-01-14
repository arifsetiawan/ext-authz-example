package main

import (
	"context"
	"log"
	"net"

	ext_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	ext_authz_v3 "github.com/envoyproxy/go-control-plane/envoy/service/auth/v3"
	ext_type_v3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	"google.golang.org/genproto/googleapis/rpc/code"
	rpc_status "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
)

// AuthorizationServer ...
type AuthorizationServer struct{}

func (a *AuthorizationServer) denied(body string) *ext_authz_v3.CheckResponse {
	return &ext_authz_v3.CheckResponse{
		Status: &rpc_status.Status{
			Code: int32(code.Code_PERMISSION_DENIED),
		},
		HttpResponse: &ext_authz_v3.CheckResponse_DeniedResponse{
			DeniedResponse: &ext_authz_v3.DeniedHttpResponse{
				Status: &ext_type_v3.HttpStatus{
					Code: ext_type_v3.StatusCode(ext_type_v3.StatusCode_Forbidden),
				},
				Body: body,
			},
		},
	}
}

func (a *AuthorizationServer) ok() *ext_authz_v3.CheckResponse {
	return &ext_authz_v3.CheckResponse{
		Status: &rpc_status.Status{
			Code: int32(code.Code_OK),
		},
		HttpResponse: &ext_authz_v3.CheckResponse_OkResponse{
			OkResponse: &ext_authz_v3.OkHttpResponse{
				Headers: []*ext_core_v3.HeaderValueOption{},
			},
		},
	}
}

// Check ...
func (a *AuthorizationServer) Check(ctx context.Context, req *ext_authz_v3.CheckRequest) (*ext_authz_v3.CheckResponse, error) {
	log.Printf("Check %+v", req.Attributes)

	return a.ok(), nil
}

func main() {
	// create a TCP listener on port 4000
	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("listening on %s", listener.Addr())

	grpcServer := grpc.NewServer()
	authServer := &AuthorizationServer{}
	ext_authz_v3.RegisterAuthorizationServer(grpcServer, authServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
