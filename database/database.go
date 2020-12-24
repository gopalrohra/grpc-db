package database

import (
	pb "github.com/gopalrohra/grpcdb/grpc_database"
)

type Database interface {
	CreateDatabase(*pb.Database) (*pb.DatabaseResult, error)
	CreateTable(*pb.TableRequest) (*pb.TableResponse, error)
	ExecuteSelect(*pb.SelectQuery) (*pb.SelectQueryResult, error)
}
