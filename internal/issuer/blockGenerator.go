package issuer

import (
	"encoding/hex"
	"time"
	"token-strike/internal/utils/tokenstrikemock"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/justifications"
)

func (i Issuer) bootBlockGenerator() {

	go func() {
		lockChan := i.invServer.CreateNewLockChannel()
		for {
			curLock := <-lockChan
			block, _ := i.generateLockBlock(curLock)

			i.tokendb.SaveBlock(curLock.TokenID, block)
		}
	}()

	go func() {
		txChan := i.invServer.CreateNewTxChannel()
		for {
			tx := <-txChan
			block, _ := i.generateTxBlock(tx)

			i.tokendb.SaveBlock(tx.TokenID, block)
		}
	}()
}

func (i Issuer) generateLockBlock(curLock *tokenstrikemock.LockForBlock) (*DB.Block, error) {
	chain, err := i.tokendb.GetChainInfoDB(curLock.TokenID)
	if err != nil {
		return nil, err
	}

	lastBlock := chain.Blocks[len(chain.Blocks)-1]

	blockBytes, err := lastBlock.GetHash()
	if err != nil {
		return nil, err
	}

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
		PrevBlock:      hex.EncodeToString(blockBytes),
		Justifications: []*DB.Justification{lockJustification},
		Creation:       time.Now().Unix(),
		State:          hex.EncodeToString(stateBytes),
		PktBlockHash:   string(i.config.Chain.BlockHashAtHeight(i.config.Chain.CurrentHeight())),
		PktBlockHeight: i.config.Chain.CurrentHeight(),
		Height:         lastBlock.Height + 1,
	}

	return block, nil
}

func (i Issuer) generateTxBlock(curTx *tokenstrikemock.TxForBlock) (*DB.Block, error) {
	chain, err := i.tokendb.GetChainInfoDB(curTx.TokenID)
	if err != nil {
		return nil, err
	}

	lastBlock := chain.Blocks[len(chain.Blocks)-1]

	blockBytes, err := lastBlock.GetHash()
	if err != nil {
		return nil, err
	}

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
		PrevBlock:      hex.EncodeToString(blockBytes),
		Justifications: []*DB.Justification{txJustification},
		Creation:       time.Now().Unix(),
		State:          hex.EncodeToString(stateBytes),
		PktBlockHash:   string(i.config.Chain.BlockHashAtHeight(i.config.Chain.CurrentHeight())),
		PktBlockHeight: i.config.Chain.CurrentHeight(),
		Height:         lastBlock.Height + 1,
	}

	return block, nil
}
