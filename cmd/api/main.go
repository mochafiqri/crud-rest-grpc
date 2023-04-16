package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mochafiqri/simple-crud/controllers"
	"github.com/mochafiqri/simple-crud/infrastructures"
	"github.com/mochafiqri/simple-crud/proto_gen"
	"github.com/mochafiqri/simple-crud/repository"
	"github.com/mochafiqri/simple-crud/usecases"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {
	var e = echo.New()
	e.GET("/", func(e echo.Context) error {
		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "hello word",
		})
	})

	db, err := infrastructures.InitMysql()
	if err != nil {
		panic(err)
	}

	rds, err := infrastructures.InitRedis()
	if err != nil {
		panic(err)
	}

	var repoContent = repository.NewContentRepo(db, rds)
	var ucContent = usecases.NewContentUseCase(repoContent)
	var contentApi = controllers.NewHandler(ucContent)
	contentApi.Routes(e)

	netListen, err := net.Listen("tcp", ":8181")
	if err != nil {
		panic(err)
	}

	var contentGrpc = controllers.NewContentGrpc(ucContent)

	var grpcServer = grpc.NewServer()
	proto_gen.RegisterContentServiceServer(grpcServer, contentGrpc)

	//rest
	go func() {
		log.Println("Start HTTP Server :8080")
		e.Logger.Fatal(e.Start(":8080"))
	}()

	//grpc
	go func() {
		log.Println("Start GRPC Server ", netListen.Addr())
		err = grpcServer.Serve(netListen)
		if err != nil {
			panic(err)
		}
	}()

	select {}
}
