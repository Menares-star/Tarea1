package main

import (
  "os"
  "fmt"
  "log"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "encoding/csv"
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
  f,err := os.Open("pymes.csv")
  if err != nil{
    log.Printf("error abriendo el archivo: %v",err)
  }
  defer f.Close()

  r := csv.NewReader(f)
  r.Comma = ','
  r.Comment = '#'
  r.FieldsPerRecord = 6

  rawData,err := r.ReadAll()
  if err!= nil{
    log.Printf("error leyendo la informacion del archivo: %v",err)
  }
  fmt.Println(rawData)

  message := ordenes.Orden{
    Id :"caca",
    Producto:"jabon",
    Valor:1000,
    Tienda:"Beta",
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
