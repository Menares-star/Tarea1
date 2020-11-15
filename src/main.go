package main

import (
	"log"
	"net"
	//"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"Orden"

)

type Server struct{
}

var codsum int32 =0

	func (s*Server) ReceivedOrden(ctx context.Context, message *Orden)(*Orden, error){
		codsum++
		log.Printf("Received message body from client: %s", message.Id)
		return &Orden{Id :message.Id,
		Producto:message.Producto,
		Valor:message.Valor,
		Tienda:message.Tienda,
		Destino:message.Destino,
		Prioridad:message.Prioridad,
		Codigo:codsum,
		Tipo:message.Tipo}, nil
	}

	func (s*Server) ReceivedSeguimiento(ctx context.Context, message *Seguimiento)(*Seguimiento, error){
		log.Printf("Received message seguimiento from client: %d", message.Codigo)
		return &Seguimiento{Codigo:message.Codigo,
		Estado:"En Bodega",
		}, nil
	}

func main() {

	lis, err := net.Listen("tcp", ":7000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := Server{}

	grpcServer := grpc.NewServer()

	//REGISTRO DE SERVICIOS
	RegisterOrdenServiceServer(grpcServer, &s)
	RegisterSeguimientoServiceServer(grpcServer, &s)
	//FIN REGISTRO DE SERVICIOS


	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	//SEGUIMIENTOS
	//FIN SEGUIMIENTOS
	//COMUNICACION CON CAMIONES

}
