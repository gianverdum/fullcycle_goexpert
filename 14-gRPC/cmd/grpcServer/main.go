package main

import (
	"database/sql"
	"net"
	"path/filepath"

	"github.com/gianverdum/fullcycle_goexpert/14-gRPC/internal/database"
	"github.com/gianverdum/fullcycle_goexpert/14-gRPC/internal/pb"
	"github.com/gianverdum/fullcycle_goexpert/14-gRPC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	absPath, err := filepath.Abs("./db.sqlite")
    if err != nil {
        panic(err)
    }

	db, err := sql.Open("sqlite3", absPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
