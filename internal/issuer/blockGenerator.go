package issuer

import (
	"context"
	"encoding/hex"
	"time"
	"token-strike/internal/utils/tokenstrikemock"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/justifications"
	"token-strike/tsp2p/server/tokenstrike"
)

// TODO: Error handling
func (i *Issuer) startBlockGenerator(tokenID string) {
	dispather := i.invServer.Subscribe(tokenID)
	go func() {
		for {
			curLock := <-dispather.Lock

			block, err := i.generateLockBlock(curLock)
			if err != nil {
				panic(err)
			}

			err = i.tokendb.SaveBlock(curLock.TokenID, block)
			if err != nil {
				panic(err)
			}

			resp, err := i.invServer.GetTokenStatus(context.Background(), &tokenstrike.TokenStatusReq{
				Tokenid: curLock.TokenID,
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
					ParentHash: curLock.TokenID,
					Type:       tokenstrike.TYPE_BLOCK,
					Message:    block,
					Expiration: 123,
				})
		}
	}()

	go func() {
		for {
			tx := <-dispather.TX

			block, err := i.generateTxBlock(tx)
			if err != nil {
				panic(err)
			}

			err = i.tokendb.SaveBlock(tx.TokenID, block)
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
					ParentHash: tx.TokenID,
					Type:       tokenstrike.TYPE_BLOCK,
					Message:    block,
					Expiration: 123,
				})
		}
	}()
}

func (i Issuer) generateLockBlock(curLock *tokenstrikemock.LockEvent) (*DB.Block, error) {
	chain, err := i.tokendb.GetChainInfoDB(curLock.TokenID)
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
				Lock: &curLock.Content,
			},
		},
	}

	block := &DB.Block{
		PrevBlock:      lastBlock.GetSignature(),
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

func (i Issuer) generateTxBlock(curTx *tokenstrikemock.TxEvent) (*DB.Block, error) {
	chain, err := i.tokendb.GetChainInfoDB(curTx.TokenID)
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
			Transfer: &curTx.Content,
		},
	}

	block := &DB.Block{
		PrevBlock:      lastBlock.GetSignature(),
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
