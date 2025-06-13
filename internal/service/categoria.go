package service

import (
	"context"

	"github.com/Higor-ViniciusDev/grpc/internal/database"
	"github.com/Higor-ViniciusDev/grpc/internal/pb"
)

type CategoriaService struct {
	pb.UnimplementedCategoriaServiceServer
	CategoriaDB *database.Categoria
}

func NewCategoriaService(db *database.Categoria) *CategoriaService {
	return &CategoriaService{
		CategoriaDB: db,
	}
}

func (c *CategoriaService) CriarCategoria(_ context.Context, in *pb.CreateCategoria) (*pb.CategoriaResponse, error) {
	categoria, err := c.CategoriaDB.Create(in.Nome, in.Descricao)

	if err != nil {
		return nil, err
	}

	CategoriaService := &pb.Categoria{
		Id:        categoria.ID,
		Nome:      categoria.Nome,
		Descricao: categoria.Descricao,
	}

	return &pb.CategoriaResponse{
		Categoria: CategoriaService,
	}, nil
}
