package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/tutorialedge/go-grpc-tutorial/Mensajes"
)

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := ordenes.Server{}

	grpcServer := grpc.NewServer()

	ordenes.RegisterOrdenServiceServer(grpcServer, &s)
	ordenes.RegisterSeguimientoServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
