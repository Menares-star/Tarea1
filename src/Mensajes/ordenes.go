package ordenes
 import (
   "log"
   "golang.org/x/net/context"
 )

 type Server struct{
 }

   func (s*Server) ReceivedOrden(ctx context.Context, message *Orden)(*Orden, error){
     log.Printf("Received message body from client: %s", message.Id)
     return &Orden{Id :"caca",
     Producto:"jabon",
     Valor:1000,
     Destino:"casa-A",
     Prioridad:1,
     Codigo:1,}, nil
   }
