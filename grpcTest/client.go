package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"google.golang.org/grpc"

	pb "github.com/wangle201210/goPkgLearn/grpcTest/proto"
)

const PORTCLIENT = "9001"

func main() {
	conn, err := grpc.Dial(":"+PORTCLIENT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
	defer conn.Close()
	client := pb.NewSearchServiceClient(conn)
	//tick := time.NewTicker(time.Second * 10)
	i := 0
	for  {
		//select {
		//case <-tick.C:
			i++
			fmt.Println(i)
			md5Resp, err := client.MD5(context.Background(), &pb.SearchRequest{
				Request: strconv.Itoa(i),
			})
			if err != nil {
				log.Fatalf("client.Search err: %v", err)
			}
			log.Printf("resp: %s", md5Resp.GetResponse())
		//}
	}
}
