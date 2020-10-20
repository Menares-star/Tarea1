# Tarea1
Grupo Gama Corp
Martín Menares 201573536-0
Gabriel Pezoa 201673562-3

Dado la complejidad de la Tarea y la poca asesoría dada en la ayudantía donde el ayudante
dijo con sus propias palabras de no sabia nada de Grpc, donde implicaba aprender un lenguaje nuevo y donde
no se consideró que los alumnos del ramos son personas y tienen otras responsabilidades, el grupo sólo
logró avanzar con la comunicación Cliente-Logística donde se pide al usuario ingresar un tiempo entre ordenes, luego el programa elige aleatoriamente la orden que desea enviar, el servidor de Logística asigna un código de seguimiento para cada paquete (Orden), se lleva un registro de los ordenes enviadas por el cliente, se establece un tiempo para generar consultas de seguimiento donde Logística sólo responde que el estado de los paquetes es "En Bodega". Dicho lo anterior sólo se ocupan las siguientes máquinas Virtuales:

ssh root@dist73
4Q3L8hhL

ssh root@dist74
WZX9zfXC

Para probar la simulación se debe acceder al siguiente directorio para cada MV: cd Tarea1/src/

Luego en la MV dist73 utilizar el comando: go run main.go

y luego para la MV dist74 utilizar el comando: go run client.go
