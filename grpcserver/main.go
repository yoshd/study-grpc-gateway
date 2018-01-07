package main

import (
	"flag"
	"net"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/yoshd/study-grpc-gateway/pb"
)

var (
	addr = flag.String("addr", "localhost:9090", "addr host:port")
)

type server struct{}

func (s *server) Echo(ctx context.Context, in *pb.StringMessage) (*pb.StringMessage, error) {
	glog.Infof("New Request: %v", in.String())
	return &pb.StringMessage{Value: in.Value}, nil
}

func main() {
	flag.Parse()
	defer glog.Flush()

	lis, err := net.Listen("tcp", *addr)

	if err != nil {
		glog.Fatal(err)
	}

	s := grpc.NewServer()

	pb.RegisterSampleServer(s, &server{})
	s.Serve(lis)
}
