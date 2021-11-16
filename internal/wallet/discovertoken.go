package wallet

import (
	"context"
	"errors"
	"token-strike/internal/types/dispatcher"
	"token-strike/tsp2p/server/rpcservice"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

const emptyIssue = "empty issue url collection"

func (s *Server) DiscoverToken(ctx context.Context, req *rpcservice.DiscoverTokenRequest) (*empty.Empty, error) {
	disp, ok := s.inv.Subscribe(req.ParentHash).(dispatcher.TokenDispatcher)
	if !ok {
		return &emptypb.Empty{}, errors.New("subscribe return not correct type")
	}

	s.dispatcher = disp

	s.dispatcher.Observe()

	return &emptypb.Empty{}, nil
}
