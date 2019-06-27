package main

import (
	"context"
	"flag"
	"github.com/golang/glog"
	api "github.com/tianhanfangyan/grpc-api-demo/api/proto"
	"google.golang.org/grpc"
	"log"
)

var (
	port = ":5000"
)

func printGet(ctx context.Context, in *api.StringMessage, client api.RestServiceClient) {
	out, err := client.Get(ctx, in)

	if err != nil {
		glog.Fatalf("the server must ")
	}

	log.Println(out)
}

func printPost(ctx context.Context, in *api.StringMessage, client api.RestServiceClient) {
	out, err := client.Post(ctx, in)

	if err != nil {
		glog.Fatalf("the server must ")
	}

	log.Println(out)
}

func main() {

	flag.Parse()

	conn, err := grpc.Dial(port, grpc.WithInsecure())

	if err != nil {
		glog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	client := api.NewRestServiceClient(conn)

	printGet(ctx, &api.StringMessage{Value: port}, client)
	printPost(ctx, &api.StringMessage{Value: port}, client)

}
