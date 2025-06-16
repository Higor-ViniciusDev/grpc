package service

import (
	"context"
	"io"

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

func (c *CategoriaService) GetCategoriaById(_ context.Context, in *pb.GetCategoriaByIdInput) (*pb.Categoria, error) {
	categoriaDB, err := c.CategoriaDB.FindById(in.Id)

	if err != nil {
		return nil, err
	}
	CategoriaService := &pb.Categoria{
		Id:        categoriaDB.ID,
		Nome:      categoriaDB.Nome,
		Descricao: categoriaDB.Descricao,
	}

	return CategoriaService, nil
}

func (c *CategoriaService) CriarCategoriaStream(stream pb.CategoriaService_CriarCategoriaStreamServer) error {
	categorias := &pb.ListaDeCategorias{}

	for {
		categoria, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(categorias)
		}

		if err != nil {
			return err
		}

		categoriaResult, err := c.CategoriaDB.Create(categoria.Nome, categoria.Descricao)

		if err != nil {
			return err
		}

		categorias.Categoria = append(categorias.Categoria, &pb.Categoria{
			Id:        categoriaResult.ID,
			Nome:      categoriaResult.Nome,
			Descricao: categoriaResult.Descricao,
		})
	}
}

func (c *CategoriaService) CriarCategoriaStreamBI(stream pb.CategoriaService_CriarCategoriaStreamBIServer) error {
	for {
		categoria, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		categoriaResult, err := c.CategoriaDB.Create(categoria.Nome, categoria.Descricao)

		if err != nil {
			return err
		}

		err = stream.Send(&pb.Categoria{
			Id:        categoriaResult.ID,
			Nome:      categoriaResult.Nome,
			Descricao: categoriaResult.Descricao,
		})
	}
}
