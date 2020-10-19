package ordenes
 import (
   "log"
   "golang.org/x/net/context"
 )

 type Server struct{
 }

var codsum int32 =0

   func (s*Server) ReceivedOrden(ctx context.Context, message *Orden)(*Orden, error){
     codsum++
     log.Printf("Received message body from client: %s", message.Id)
     return &Orden{Id :"caca",
     Producto:"jabon",
     Valor:1000,
     Tienda:"Beta",
     Destino:"casa-A",
     Prioridad:1,
     Codigo:codsum,
     Tipo:"pymes"}, nil
   }

   func (s*Server) ReceivedSeguimiento(ctx context.Context, message *Seguimiento)(*Seguimiento, error){
     log.Printf("Received message seguimiento from client: %i", message.Codigo)
     return &Seguimiento{Codigo:message.Codigo,
     Estado:"En Bodega",
     }, nil
   }
