package main

import (
	"context"
	"log"
	"time"

	pb "github.com/oninugarte/grpc-examples/notificador/notificador"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	//Conexión al server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error en la conexion: %v", err)
	}
	defer conn.Close()
	c := pb.NewNotificadorClient(conn)

	//Enviamos petición
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.EnviarCorreo(ctx, &pb.CorreoRequest{Destino: "ugarteonin@gmail.com", Asunto: "Prueba", Mensaje: "Este es un mensaje de prueba"})
	if err != nil {
		log.Fatalf("Error en la petición: %v", err)
	}

	//imprimmimos respuesta
	log.Println("¿El correo fue enviado?", r.Enviado)
}
