package main

import (
	"github.com/idylicaro/full-cycle-immersive/codepix-go/application/grpc"
	"github.com/idylicaro/full-cycle-immersive/codepix-go/infrastructure/db"
	"github.com/jinzhu/gorm"
	"os"
)

var database *gorm.DB

func main(){
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}