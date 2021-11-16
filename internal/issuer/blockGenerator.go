package issuer

import (
	"context"
	"encoding/hex"
	"time"
	"token-strike/internal/types/dispatcher"
	"token-strike/internal/utils/tokenstrikemock"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/justifications"
	"token-strike/tsp2p/server/lock"
	"token-strike/tsp2p/server/tokenstrike"
)

// TODO: Error handling
func (i *Issuer) startBlockGenerator(tokenID string) {
	dispather := i.invServer.Subscribe(tokenID).(dispatcher.TokenDispatcher)

	go func() {
		for {
			wait := dispather.WaitLockAction(func(l lock.Lock) (string, error) {
				block, err := i.generateLockBlock(tokenID, l)
				if err != nil {
					panic(err)
				}

				err = i.tokendb.SaveBlock(tokenID, block)
				if err != nil {
					panic(err)
				}

				resp, err := i.invServer.GetTokenStatus(context.Background(), &tokenstrike.TokenStatusReq{
					Tokenid: tokenID,
				})
				if err != nil {
					panic(err)
				}

				blockHash, err := resp.Dblock0.GetHash()
				if err != nil {
					panic(err)
				}

				block32 := [32]byte{}
				copy(block32[:], blockHash[:32])

				_ = i.invServer.Insert(
					tokenstrikemock.MempoolEntry{
						ParentHash: tokenID,
						Type:       tokenstrike.TYPE_BLOCK,
						Message:    block,
						Expiration: 123,
					},
				)
				return "", nil
			})

			<-wait
		}
	}()

	go func() {
		for {
			wait := dispather.WaitTxAction(func(tx justifications.TranferToken) (string, error) {
				block, err := i.generateTxBlock(tokenID, tx)
				if err != nil {
					panic(err)
				}

				err = i.tokendb.SaveBlock(tokenID, block)
				if err != nil {
					panic(err)
				}

				blockHash, err := block.GetHash()
				if err != nil {
					panic(err)
				}

				block32 := [32]byte{}
				copy(block32[:], blockHash[:32])

				_ = i.invServer.Insert(
					tokenstrikemock.MempoolEntry{
						ParentHash: tokenID,
						Type:       tokenstrike.TYPE_BLOCK,
						Message:    block,
						Expiration: 123,
					},
				)

				return "", nil
			})

			<-wait
		}
	}()

	dispather.Observe()

}

func (i Issuer) generateLockBlock(tokenID string, l lock.Lock) (*DB.Block, error) {
	chain, err := i.tokendb.GetChainInfoDB(tokenID)
	if err != nil {
		return nil, err
	}

	// Blocks are reversed
	lastBlock := chain.Blocks[0]

	stateBytes, err := chain.State.GetHash()
	if err != nil {
		return nil, err
	}

	lockJustification := &DB.Justification{
		Content: &DB.Justification_Lock{
			Lock: &justifications.LockToken{
				Lock: &l,
			},
		},
	}

	lastBlockHash, err := lastBlock.GetHash()
	if err != nil {
		return nil, err
	}

	block := &DB.Block{
		PrevBlock:      hex.EncodeToString(lastBlockHash),
		Justifications: []*DB.Justification{lockJustification},
		Creation:       time.Now().Unix(),
		State:          hex.EncodeToString(stateBytes),
		PktBlockHash:   hex.EncodeToString(i.config.Chain.BlockHashAtHeight(i.config.Chain.CurrentHeight())),
		PktBlockHeight: i.config.Chain.CurrentHeight(),
		Height:         lastBlock.Height + 1,
	}

	err = block.Sing(i.private)
	if err != nil {
		return nil, err
	}

	return block, nil
}

func (i Issuer) generateTxBlock(tokenID string, tx justifications.TranferToken) (*DB.Block, error) {
	chain, err := i.tokendb.GetChainInfoDB(tokenID)
	if err != nil {
		return nil, err
	}

	// Blocks are reversed
	lastBlock := chain.Blocks[0]

	stateBytes, err := chain.State.GetHash()
	if err != nil {
		return nil, err
	}

	txJustification := &DB.Justification{
		Content: &DB.Justification_Transfer{
			Transfer: &tx,
		},
	}

	lastBlockHash, err := lastBlock.GetHash()
	if err != nil {
		return nil, err
	}

	block := &DB.Block{
		PrevBlock:      hex.EncodeToString(lastBlockHash),
		Justifications: []*DB.Justification{txJustification},
		Creation:       time.Now().Unix(),
		State:          hex.EncodeToString(stateBytes),
		PktBlockHash:   hex.EncodeToString(i.config.Chain.BlockHashAtHeight(i.config.Chain.CurrentHeight())),
		PktBlockHeight: i.config.Chain.CurrentHeight(),
		Height:         lastBlock.Height + 1,
	}

	err = block.Sing(i.private)
	if err != nil {
		return nil, err
	}

	return block, nil
}

func (i *Issuer) GetGenesisBlock(tokenID string) [32]byte {
	resp, _ := i.invServer.GetTokenStatus(context.Background(), &tokenstrike.TokenStatusReq{
		Tokenid: tokenID,
	})

	blockHash, _ := resp.Dblock0.GetHash()

	block32 := [32]byte{}
	copy(block32[:], blockHash[:32])

	return block32
}
