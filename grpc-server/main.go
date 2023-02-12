package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	// import module using buf
	"github.com/firacloudtech/grpc-echo-benchmark/db"
	orderv1 "github.com/firacloudtech/grpc-echo-benchmark/gen/go/order/v1"
	productv1 "github.com/firacloudtech/grpc-echo-benchmark/gen/go/product/v1"
	"github.com/google/uuid"
	"github.com/lib/pq"
	_ "github.com/lib/pq"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

var (
	serverAddr  = "127.0.0.1"
	port        = 5002
	gatewayPort = 3001
	wg          sync.WaitGroup
)

type combinedServer struct {
	productv1.UnimplementedProductServiceServer
	orderv1.UnimplementedOrderServiceServer
}

func NewServer() *combinedServer {
	return &combinedServer{}
}

func main() {
	// initialize DB
	db.InitDB()
	defer db.Db.Close()

	wg.Add(2)
	go func() {
		if err := run(); err != nil {
			log.Fatal(err)
		}
		wg.Done()
	}()
	go func() {

		if err := runGrpcGateway(); err != nil {
			log.Fatal(err)
		}

		wg.Done()
	}()

	wg.Wait()

}

func run() error {
	grpcAddr := fmt.Sprintf("%v:%d", serverAddr, port)
	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %v", grpcAddr, err)

	}
	defer listener.Close()

	// Register the services
	s := grpc.NewServer()
	productv1.RegisterProductServiceServer(s, &combinedServer{})
	orderv1.RegisterOrderServiceServer(s, &combinedServer{})

	log.Println("Listening on port", grpcAddr)
	if err := s.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)

	}

	return nil

}

func runGrpcGateway() error {

	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames:   true,
			EmitUnpopulated: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	// grpc gateway server
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%v:%v", serverAddr, port), grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}
	defer conn.Close()

	gwmux := runtime.NewServeMux(jsonOption)
	err = productv1.RegisterProductServiceHandler(context.Background(), gwmux, conn)

	if err != nil {
		log.Fatalln("Failed to register product service:", err)
	}

	err = orderv1.RegisterOrderServiceHandler(context.Background(), gwmux, conn)

	if err != nil {
		log.Fatalln("Failed to register order service:", err)
	}

	// Create a new ServeMux for serving the Swagger-UI files
	swaggerMux := http.NewServeMux()
	swaggerMux.Handle("/", gwmux)

	fs := http.FileServer(http.Dir("../gen/openapiv2"))
	swaggerMux.Handle("/docs/", http.StripPrefix("/docs/", fs))

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%v", gatewayPort),
		Handler: swaggerMux,
	}

	log.Printf("Serving gRPC-gateway on http://%v%v", serverAddr, gwServer.Addr)

	err = gwServer.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to serve gRPC gateway server: %w", err)
	}

	return nil
}

// ----------------------------------------------------------------

// a function that creates a product and return the response
func (s *combinedServer) CreateProduct(ctx context.Context, req *productv1.CreateProductRequest) (*productv1.CreateProductResponse, error) {

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	product := &db.Product{
		ID:          id.String(),
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Price:       float64(req.GetPrice()),
		Category:    req.GetCategory(),
		ImageUrl:    req.GetUrl(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	queries := db.New(db.Db)

	params := db.CreateProductParams{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    product.Category,
		ImageUrl:    product.ImageUrl,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}

	result, err := queries.CreateProduct(ctx, params)

	log.Printf("index is: %v", result)
	if err, ok := err.(*pq.Error); ok {
		fmt.Println("pq error:", err.Code.Name())
	}

	log.Printf("Received a request to create a product: %v\n", result)

	if err != nil {
		fmt.Print("Unable to save to db: " + err.Error())
		return nil, err
	}

	return &productv1.CreateProductResponse{
		Id:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		Category:    result.Category,
		Url:         result.ImageUrl,
	}, nil

}

func (s *combinedServer) ReadProduct(ctx context.Context, req *productv1.ReadProductRequest) (*productv1.ReadProductResponse, error) {

	id := req.GetId()

	log.Printf("Received a request to read a product: %v\n", id)

	return &productv1.ReadProductResponse{
		Name: "Product 1",
	}, nil

}

func (s *combinedServer) CreateOrder(ctx context.Context, req *orderv1.CreateOrderRequest) (*orderv1.CreateOrderResponse, error) {
	name := req.GetName()

	log.Printf("Received a request to create a product: %v\n", name)

	return &orderv1.CreateOrderResponse{
		Message: "Order created successfully",
	}, nil

}
