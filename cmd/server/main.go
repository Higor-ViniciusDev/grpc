package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/Higor-ViniciusDev/grpc/internal/database"
	"github.com/Higor-ViniciusDev/grpc/internal/pb"
	"github.com/Higor-ViniciusDev/grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "file:./db.sqlite")

	if err != nil {
		panic(err)
	}

	categoriaDB := database.NewCategoria(db)
	CategoriaService := service.NewCategoriaService(categoriaDB)

	grpcServe := grpc.NewServer()
	pb.RegisterCategoriaServiceServer(grpcServe, CategoriaService)
	reflection.Register(grpcServe)

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	log.Printf("server listening at %v", listen.Addr())

	services := grpcServe.GetServiceInfo()
	for name := range services {
		log.Println("✅ Serviço registrado:", name)
	}

	if err := grpcServe.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
