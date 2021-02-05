package grpc
/*
"github.com/idylicaro/full-cycle-immersive/codepix-go/application/grpc/pb"
	"github.com/idylicaro/full-cycle-immersive/codepix-go/application/usecase"
	"github.com/idylicaro/full-cycle-immersive/codepix-go/infrastructure/repository"
	"google.golang.org/grpc/reflection"
*/
import (
	"fmt"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpcServer := grpc.NewServer()

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}

	log.Printf("gRPC server has been started on port %d", port)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}
}