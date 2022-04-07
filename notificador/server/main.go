package main

import (
	"context"
	"log"
	"net"

	pb "github.com/oninugarte/grpc-examples/notificador/notificador"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) EnviarCorreo(ctx context.Context, in *pb.CorreoRequest) (*pb.CorreoResponse, error) {
	log.Printf("Enviando correo: %v a %v", in.Asunto, in.Destino)
	return &pb.CorreoResponse{Enviado: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNotificadorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
