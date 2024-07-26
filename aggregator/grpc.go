package main

import (
	"context"

	"github.com/fulltimegodev/tolling/types"
)

type GPRCAggregatorServer struct {
	types.UnimplementedAggregatorServer
	svc Aggregator
}

func NewAggregatorGRPCServer(svc Aggregator) *GPRCAggregatorServer {
	return &GPRCAggregatorServer{
		svc: svc,
	}
}

// transport layer
// JSON -> types.Distance -> all done (same type)
// GRPC -> types.AggregateRequest -> type.Distance
// Webpack => types.WEBpack -> types.Distance

// business layer -> business layer type (main type everyone needs to convert to)

func (s *GPRCAggregatorServer) Aggregate(ctx context.Context, req *types.AggregateRequest) (*types.None, error) {
	distance := types.Distance{
		OBUID: int(req.ObuID),
		Value: req.Value,
		Unix:  req.Unix,
	}
	return &types.None{}, s.svc.AggregateDistance(distance)
}
