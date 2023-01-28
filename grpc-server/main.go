package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	// import module using buf
	orderv1 "buf.build/gen/go/firacloudtech/grpc-echo-benchmark/protocolbuffers/go/order/v1"
	productv1 "buf.build/gen/go/firacloudtech/grpc-echo-benchmark/protocolbuffers/go/product/v1"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	orderGrpc "buf.build/gen/go/firacloudtech/grpc-echo-benchmark/grpc/go/order/v1/orderv1grpc"
	productGrpc "buf.build/gen/go/firacloudtech/grpc-echo-benchmark/grpc/go/product/v1/productv1grpc"
)

var (
	port        = 5002
	gatewayPort = 3001
)

type combinedServer struct {
	productGrpc.UnimplementedProductServiceServer
	orderGrpc.UnimplementedOrderServiceServer
}

func NewServer() *combinedServer {
	return &combinedServer{}
func NewServer() *combinedServer {
	return &combinedServer{}
}

func main() {

	if err := run(); err != nil {
		log.Fatal(err)
	}

	// wait till the grpc server is ready
	done := make(chan struct{})

	go func() {

		if err := runGrpcGateway(); err != nil {
			log.Fatal(err)
		}
		close(done)
	}()
	<-done

}

func run() error {
	grpcAddr := fmt.Sprintf("127.0.0.1:%d", port)
	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %v", grpcAddr, err)

	}
	defer listener.Close()

	// Register the services
	server := grpc.NewServer()
	productGrpc.RegisterProductServiceServer(server, &combinedServer{})
	orderGrpc.RegisterOrderServiceServer(server, &combinedServer{})

	log.Println("Listening on port", grpcAddr)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)

	}

	return nil
}

func runGrpcGateway() error {
	// grpc gateway server

	gwmux := runtime.NewServeMux()
	gwAddr := fmt.Sprintf("127.0.0.1:%d", gatewayPort)
	gwServer := &http.Server{
		Addr:    gwAddr,
		Handler: gwmux,
	}

	conn, err := grpc.DialContext(context.Background(), fmt.Sprintf("127.0.0.1:%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock)

	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	log.Println("Serving gRPC-gateway on http://localhost:", gatewayPort)

	if err := gwServer.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to serve gRPC gateway server: %w", err)
	}

	return nil
}

// a function that creates a product and return the response
func (s *combinedServer) CreateProduct(ctx context.Context, req *productv1.CreateRequest) (*productv1.CreateResponse, error) {

	name := req.GetName()

	log.Printf("Got a request to create a product: %v\n", name)

	return &productv1.CreateResponse{
		Message: "Success",
	}, nil

}

func (s *combinedServer) CreateOrder(ctx context.Context, req *orderv1.CreateRequest) (*orderv1.CreateResponse, error) {

	name := req.GetName()

	log.Printf("Got a request to create a order: %v\n", name)

	return &orderv1.CreateResponse{}, nil
	return &orderv1.CreateResponse{}, nil

}
