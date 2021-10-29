package issuer

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
	address2 "token-strike/internal/types/address"
	"token-strike/internal/types/users"

	"token-strike/internal/utils/address"
	"token-strike/internal/utils/config"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/tokenstrike"

	"github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
)

type SimpleIssuer struct {
	config    *config.Config
	invServer tokenstrike.TokenStrikeClient

	private address2.PrivateKey
	address address2.Address
}

var _ users.Issuer = &SimpleIssuer{}

func CreateIssuer(cfg *config.Config, pk address2.PrivateKey, target string) (users.Issuer, error) {
	grpcConn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &SimpleIssuer{
		private: pk,
		address: address.NewSimpleAddress(pk.GetPublicKey()),

		config:    cfg,
		invServer: tokenstrike.NewTokenStrikeClient(grpcConn),
	}, nil
}

func (s *SimpleIssuer) IssueToken(owners []*DB.Owner, expiration int32) (string, error) {
	token := &DB.Token{
		Count:        10,
		Expiration:   expiration,
		Creation:     time.Now().Unix(),
		IssuerPubkey: s.address.String(),
		Urls: []string{
			"http://localhost:3333/token1",
		},
	}

	state := &DB.State{
		Token:  token,
		Owners: owners,
		Locks:  nil,
	}

	stateBytes, err := state.GetHash()
	if err != nil {
		return "", err
	}

	block := &DB.Block{
		PrevBlock:      "",
		Justifications: nil,
		Creation:       time.Now().Unix(),
		State:          hex.EncodeToString(stateBytes),
		PktBlockHash:   string(s.config.Chain.BlockHashAtHeight(s.config.Chain.CurrentHeight())),
		PktBlockHeight: s.config.Chain.CurrentHeight(),
		Height:         0,
	}

	err = block.Sing(s.private)
	if err != nil {
		return "", err
	}

	blockSigned, err := proto.Marshal(block)
	if err != nil {
		return "", err
	}

	blockHash := sha256.Sum256(blockSigned)
	tokenID := hex.EncodeToString(blockHash[:])

	err = s.config.DB.SaveIssuerTokenDB(tokenID, s.address.String())
	if err != nil {
		return "", err
	}

	return tokenID, s.config.DB.IssueTokenDB(tokenID, token, block, state)
}
