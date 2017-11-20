package kgrpc

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kaiguangyun/kgrpc/helper"
	"github.com/kaiguangyun/kgrpc/debug"
	"github.com/kaiguangyun/kgrpc/auth"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"net"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	// server info
	ServerName    string
	ServerPath    string
	ServerAddress string
	// grpc server pointer
	GrpcServer *grpc.Server
	// ssl
	SslSecure   bool
	sslCertFile string
	sslKeyFile  string
	// grpc gateway
	GrpcGateway    bool
	GatewayAddress string
	// auth
	ServerAuth bool
}

var ServerConfig Config

func init() {
	ServerConfig.ServerPath, _ = os.Getwd()
	ServerConfig.ServerAddress = net.JoinHostPort(helper.GetEnv("ServerAddress"), helper.GetEnv("ServerPort"))
	ServerConfig.GatewayAddress = net.JoinHostPort(helper.GetEnv("ServerAddress"), helper.GetEnv("ServerGatewayPort"))
	ServerConfig.GrpcGateway = strings.ToLower(helper.GetEnv("ServerGateway")) == "true"
	ServerConfig.SslSecure = strings.ToLower(helper.GetEnv("ServerSSL")) == "true"
	ServerConfig.sslCertFile = ServerConfig.ServerPath + helper.GetEnv("ServerSSLCertFile")
	ServerConfig.sslKeyFile = ServerConfig.ServerPath + helper.GetEnv("ServerSSLKeyFile")
	ServerConfig.ServerName = helper.GetEnv("ServerName")
	ServerConfig.ServerAuth = strings.ToLower(helper.GetEnv("ServerENV")) != "local"
}

// GetServer
func GetServer() *grpc.Server {
	// options
	var opts []grpc.ServerOption
	if ServerConfig.SslSecure {
		creds, err := credentials.NewServerTLSFromFile(ServerConfig.sslCertFile, ServerConfig.sslKeyFile)
		if err != nil {
			debug.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	// interceptor
	opts = registerInterceptor(opts)
	// new server
	ServerConfig.GrpcServer = grpc.NewServer(opts...)

	return ServerConfig.GrpcServer
}

// RunServer
func RunServer() {
	listen, err := net.Listen("tcp", ServerConfig.ServerAddress)
	if err != nil {
		debug.Fatalf("Net failed to listen: %v", err)
	}
	// output server address
	helper.Outputf("Service address is %v", ServerConfig.ServerAddress)
	// run server
	err = ServerConfig.GrpcServer.Serve(listen)
	if err != nil {
		debug.Fatalf("Grpc failed to serve: %v", err)
	}
}

// GetGateway
func GetGateway() (ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption, cancel context.CancelFunc) {
	endpoint = ServerConfig.ServerAddress
	ctx = context.Background()
	ctx, cancel = context.WithCancel(ctx)
	//defer cancel()
	mux = runtime.NewServeMux()
	opts = []grpc.DialOption{grpc.WithInsecure()}

	return ctx, mux, endpoint, opts, cancel
}

// RunGateway
func RunGateway(mux *runtime.ServeMux, cancel context.CancelFunc) {
	defer cancel()

	helper.Outputf("Grpc gatewate listen on %v", ServerConfig.GatewayAddress)
	err := http.ListenAndServe(ServerConfig.GatewayAddress, mux)
	if err != nil {
		debug.Fatalf("Gateway failed to ListenAndServe: %v", err)
	}
}

// GetClient
func GetClient() (conn *grpc.ClientConn) {
	var opts []grpc.DialOption
	var creds credentials.TransportCredentials
	var err error
	if ServerConfig.SslSecure {
		creds, err = credentials.NewClientTLSFromFile(ServerConfig.sslCertFile, ServerConfig.ServerName)
		if err != nil {
			debug.Fatalf("did not connect: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	// output server address
	//helper.Outputf("Client address is %v", ServerConfig.ServerAddress)
	// client connection
	conn, err = grpc.Dial(ServerConfig.ServerAddress, opts...)
	if err != nil {
		debug.Fatalf("Grpc failed to dial: %v", err)
	}
	return conn
}

// register interceptor
func registerInterceptor(opts []grpc.ServerOption) []grpc.ServerOption {
	// register auth interceptor
	if ServerConfig.ServerAuth {
		opts = append(opts, authInterceptor())
	}
	return opts
}

// auth interceptor : jwt token auth
func authInterceptor() grpc.ServerOption {
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if err = checkAuth(ctx, info); err != nil {
			return
		}
		return handler(ctx, req)
	}
	return grpc.UnaryInterceptor(interceptor)
}

// auth
func checkAuth(ctx context.Context, info *grpc.UnaryServerInfo) error {
	// excluded method
	excludedMethodSlice := strings.Split(strings.ToLower(helper.GetEnv("AuthExcludedMethod")), ",")
	if len(excludedMethodSlice) > 0 {
		infoFullMethod := strings.ToLower(info.FullMethod)
		for _, method := range excludedMethodSlice {
			if strings.HasSuffix(infoFullMethod, method) {
				return nil
			}
		}
	}
	// jwt token
	tokenString := ""                  // tokenString
	//md, _ := metadata.FromContext(ctx) // metadata
	md, _ := metadata.FromIncomingContext(ctx) // metadata
	// grpc authorization token : md["authorization"]
	// gateway authorization token : md["grpcgateway-authorization"]
	if tokenSlice, ok := md["authorization"]; ok {
		tokenString = tokenSlice[0]
	}
	jwtClaims, err := auth.JwtGetClaims(tokenString)
	if err != nil {
		return helper.GrpcError(codes.Unauthenticated, auth.ErrTokenInvalid.Error())
	}
	if ! auth.JwtValidateToken(tokenString, auth.GetSecret(jwtClaims.Uuid)) {
		return helper.GrpcError(codes.Unauthenticated, auth.ErrTokenInvalid.Error())
	}
	return nil
}
