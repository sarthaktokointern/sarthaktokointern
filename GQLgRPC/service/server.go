package main

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	pb "github.com/sarthak/GQLgRPC/grpc/github.com/sarthak/GQLgRPC/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
)

type Disecturlserver struct {
	pb.UnimplementedDisecturlServer
}

func (d *Disecturlserver) Url(ctx context.Context, in *pb.Urlpath) (*pb.Disecturlpath, error) {
	UrlObj, err := url.Parse(in.URL)

	if err != nil {
		return &(pb.Disecturlpath{}), err
	}

	host := strings.Split(UrlObj.Host, ".")

	return &(pb.Disecturlpath{
		Scheme:            UrlObj.Scheme,
		Topleveldomain:    host[len(host)-1],
		Secondleveldomain: host[len(host)-2],
	}), nil
}

func main() {

	lis, err := net.Listen("tcp", ":8999")

	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	s := grpc.NewServer(grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)

	pb.RegisterDisecturlServer(s, &(pb.UnimplementedDisecturlServer{}))
	log.Printf("server listening at : %v", lis.Addr())
	grpc_prometheus.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}

	http.Handle("/metrics", promhttp.Handler())

}
