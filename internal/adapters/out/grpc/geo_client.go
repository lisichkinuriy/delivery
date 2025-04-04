package grpc

import (
	"context"
	"errors"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"lisichkinuriy/delivery/internal/adapters/ports"
	"lisichkinuriy/delivery/internal/domain/vo"
	"lisichkinuriy/delivery/pkg/clients/geo/geosrv/geopb"
	"time"
)

type GRPCGeoClient struct {
	conn     *grpc.ClientConn
	pbClient geopb.GeoClient
	timeout  time.Duration
}

func (c *GRPCGeoClient) GetLocation(ctx context.Context, street string) (vo.Location, error) {
	req := &geopb.GetGeolocationRequest{
		Street: street,
	}

	// Делаем запрос
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	resp, err := c.pbClient.GetGeolocation(ctx, req)
	if err != nil {
		return vo.Location{}, err
	}

	// Создаем и возвращаем VO Geo
	location, err := vo.NewLocation(int(resp.Location.X), int(resp.Location.Y))
	if err != nil {
		return vo.Location{}, err
	}
	return location, nil
}

var _ ports.IGeoClient = &GRPCGeoClient{}

func NewGRPCGeoClient(host string) (*GRPCGeoClient, error) {
	if host == "" {
		return nil, errors.New("host is empty")
	}

	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	pbClient := geopb.NewGeoClient(conn)

	return &GRPCGeoClient{
		conn:     conn,
		pbClient: pbClient,
		timeout:  5 * time.Second,
	}, nil
}

func (c *GRPCGeoClient) Close() error {
	return c.conn.Close()
}
