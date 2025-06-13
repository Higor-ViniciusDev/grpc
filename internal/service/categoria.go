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

func (c *CategoriaService) CriarCategoria(_ context.Context, in *pb.CreateCategoria) (*pb.Categoria, error) {
	categoria, err := c.CategoriaDB.Create(in.Nome, in.Descricao)

	if err != nil {
		return nil, err
	}

	CategoriaService := &pb.Categoria{
		Id:        categoria.ID,
		Nome:      categoria.Nome,
		Descricao: categoria.Descricao,
	}

	return CategoriaService, nil
}

func (c *CategoriaService) ListaCategorias(_ context.Context, _ *pb.Blank) (*pb.ListaDeCategorias, error) {
	cateSlice, err := c.CategoriaDB.FindAll()

	if err != nil {
		return nil, err
	}

	var sliCate []*pb.Categoria

	for _, row := range cateSlice {
		cateResponse := &pb.Categoria{
			Id:        row.ID,
			Nome:      row.Nome,
			Descricao: row.Descricao,
		}

		sliCate = append(sliCate, cateResponse)
	}

	return &pb.ListaDeCategorias{
		Categoria: sliCate,
	}, nil
}
