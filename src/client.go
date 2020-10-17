package main

import (
  "log"

  "golang.org/x/net/context"
  "google.golang.org/grpc"

	"github.com/tutorialedge/go-grpc-tutorial/Mensajes"
)

func main() {
  var conn *grpc.ClientConn
  conn, err := grpc.Dial(":9000", grpc.WithInsecure(), grpc.WithBlock())
  if err != nil{
    log.Fatalf("could not connect: %s",err)
  }
  defer conn.Close()

  c := ordenes.NewOrdenServiceClient(conn)

  message := ordenes.Orden{
    Id :"caca",
    Producto:"jabon",
    Valor:1000,
    Destino:"casa-A",
    Prioridad:1,
    Codigo:0,
  }
  response,err := c.ReceivedOrden(context.Background(),&message)
  if err!= nil{
    log.Fatalf("Error when calling SayHello: %s",err)
  }

  log.Printf("Response from Server: %d",response.Codigo)
}
