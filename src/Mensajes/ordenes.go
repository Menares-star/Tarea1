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
     log.Printf("Received message seguimiento from client: %i", message.Codigo)
     return &Seguimiento{Codigo:message.Codigo,
     Estado:"En Bodega",
     }, nil
   }

  
