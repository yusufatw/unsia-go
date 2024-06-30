package main

import (
	"fmt"
	"net"
	"unsia/controllers"
	"unsia/pb/cities"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":7080")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	grpcServer := grpc.NewServer()

	cityServer := controllers.City{}
	cities.RegisterCitiesServiceServer(grpcServer, &cityServer)

	fmt.Println("running server grpc")
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %s", err)
		return
	}
}
