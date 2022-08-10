package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	_ "github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	_ "net/http"
	"net/url"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/sarthak/GQLgRPC/graph/generated"
	"github.com/sarthak/GQLgRPC/graph/model"
	pb "github.com/sarthak/GQLgRPC/grpc/github.com/sarthak/GQLgRPC/grpc"
	"google.golang.org/grpc"
)

// URL is the resolver for the url field.
func (r *queryResolver) URL(ctx context.Context, input *model.Urlpath) (*model.Disecturlpath, error) {

	conn, err := grpc.Dial("localhost:8999", grpc.WithUnaryInterceptor(grpc_prometheus.UnaryClientInterceptor),
		grpc.WithStreamInterceptor(grpc_prometheus.StreamClientInterceptor), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect %v", err)
		return &model.Disecturlpath{}, err
	}

	c := pb.NewDisecturlClient(conn)
	resp, err := c.Url(ctx, &(pb.Urlpath{
		URL: *(input.URL),
	}))

	if err != nil {
		log.Fatalf("could not disect url : %v", err)
		return &model.Disecturlpath{}, err
	}
	s := resp.GetScheme()
	t := resp.GetTopleveldomain()
	se := resp.GetSecondleveldomain()
	fmt.Println(s, t, se)
	defer conn.Close()
	return &model.Disecturlpath{
		Scheme:            &s,
		Topleveldomain:    &t,
		Secondleveldomain: &se,
	}, nil
}

// Resturl is the resolver for the resturl field.
func (r *queryResolver) Resturl(ctx context.Context, input *model.Urlpath) (*model.Disecturlpath, error) {

	UrlObj, err := url.Parse(*input.URL)

	if err != nil {
		return &(model.Disecturlpath{}), err
	}

	host := strings.Split(UrlObj.Host, ".")

	return &(model.Disecturlpath{
		Scheme:            &UrlObj.Scheme,
		Topleveldomain:    &host[len(host)-1],
		Secondleveldomain: &host[len(host)-2],
	}), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
