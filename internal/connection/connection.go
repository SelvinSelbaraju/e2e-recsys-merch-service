package connection

import (
	"google.golang.org/grpc"
	"log"
)

func NewConnection(url string) *grpc.ClientConn {
	log.Printf("Starting connection to inference server at url: %s", url)
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect to endpoint %s: %v", url, err)
	}
	return conn
}
