package main

import (
	"log"
	"net"
	//"fmt"
	"google.golang.org/grpc"

	pb "github.com/Menares-star/Tarea1/src/Mensajes"
)

func main() {

	lis, err := net.Listen("tcp", ":7000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := ordenes.Server{}

	grpcServer := grpc.NewServer()

	pb.RegisterOrdenesServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	//SEGUIMIENTOS
	//FIN SEGUIMIENTOS
	//COMUNICACION CON CAMIONES

}
