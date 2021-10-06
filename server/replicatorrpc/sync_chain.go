package replicatorrpc

import (
	"io"

	"token-strike/tsp2p/server/replicator"
)

func (s *Server) SyncChain(stream replicator.Replicator_SyncChainServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		// save received blocks
		errUpdate := s.db.SyncBlock(msg.Name, msg.Blocks)
		if errUpdate != nil {
			return errUpdate
		}
	}
}
