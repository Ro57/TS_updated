package tokenstrikemock

import (
	"bytes"
	"context"
	"errors"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/lock"
	"token-strike/tsp2p/server/tokenstrike"
)

func (t TokenStrikeMock) PostData(ctx context.Context, req *tokenstrike.Data) (*tokenstrike.PostDataResp, error) {
	if req == nil {
		return nil, errors.New("nil data")
	}

	var (
		resp *tokenstrike.PostDataResp
		err  error
	)

	switch req.Data.(type) {
	case *tokenstrike.Data_Block:
		block := req.GetBlock()
		err, resp.Warning = validateBlock(block)
	case *tokenstrike.Data_Lock:
		lock := req.GetLock()
		err, resp.Warning = t.validateLock(lock)
	default:
		return nil, errors.New("unknown data type")
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

//todo place here checking for ret warnings
func validateBlock(block *DB.Block) (err error, warnings []string) {
	return nil, nil
}

//todo place here checking for ret warnings
func (t TokenStrikeMock) validateLock(lock *lock.Lock) (err error, warnings []string) {
	privateKey := SimpleAddressSchemeMock{}.GenerateKey([32]byte{})

	//Is token issued by issuer ?
	b0 := lock.GetPktBlockHash()
	isIsSigByIsaac := privateKey.Address().CheckSig(b0, []byte(lock.GetSignature()))
	if !isIsSigByIsaac {
		return errors.New("its not sign by Isaac"), nil
	}

	//Is pkt block height within 10 blocks old, not in the future ?
	curHeight := SimplePktChainMock{}.CurrentHeight()
	if int32(lock.PktBlockHeight) <= (curHeight - 10) {
		return errors.New("pkt block height too old"), nil
	}
	//Is pkt block hash correct ?
	pktBlockHash := SimplePktChainMock{}.BlockHashAtHeight(int32(lock.PktBlockHeight))
	if bytes.Compare(lock.PktBlockHash, pktBlockHash) != 0 {
		return errors.New("pkt block hash not correct"), nil
	}
	//Does sender have count tokens it their possession ? todo need to explain
	//Is expires_pkt_height at least 2 blocks in the future ? todo  need to explain
	//Is signature valid for sender address ? todo need to explain

	return nil, nil
}
