package tokenstrikemock

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/lock"
	"token-strike/tsp2p/server/tokenstrike"

	"github.com/golang/protobuf/proto"
)

func (t TokenStrikeMock) PostData(ctx context.Context, req *tokenstrike.Data) (*tokenstrike.PostDataResp, error) {
	if req == nil {
		return nil, errors.New("nil data")
	}

	var (
		resp    *tokenstrike.PostDataResp
		lockEl  *lock.Lock
		blockEl *DB.Block
		err     error
	)

	switch req.Data.(type) {
	case *tokenstrike.Data_Block:
		blockEl = req.GetBlock()
		err, resp.Warning = validateBlock(blockEl)
	case *tokenstrike.Data_Lock:
		lockEl = req.GetLock()
		err, resp.Warning = t.validateLock(lockEl)
	default:
		return nil, errors.New("unknown data type")
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

//TODO: place here checking for ret error with warnings
func validateBlock(block *DB.Block) (warnings []string, err error) {
	return nil, nil
}

//TODO: place here checking for ret warnings
func (t TokenStrikeMock) validateLock(lock *lock.Lock) (warnings []string, err error) {

	t.validateLockSignature(lock)

	//Does sender have count tokens it their possession ? TODO: need to explain
	//Is expires_pkt_height at least 2 blocks in the future ? TODO:  need to explain
	//Is signature valid for sender address ? TODO: need to explain

	return nil, nil
}

// Is token issued by issuer ?
func (t TokenStrikeMock) validateLockIssuer(lock *lock.Lock) error {
	return nil
}

// Is pkt block height within 10 blocks old, not in the future ?
func (t TokenStrikeMock) validateLockHeight(lock *lock.Lock) error {
	curHeight := t.pktChain.CurrentHeight()
	if int32(lock.PktBlockHeight) <= (curHeight - 10) {
		return errors.New("pkt block height too old")
	}

	return nil
}

//Is pkt block hash correct ?
func (t TokenStrikeMock) validateLockHashCorrect(lock *lock.Lock) error {
	pktBlockHash := t.pktChain.BlockHashAtHeight(int32(lock.PktBlockHeight))

	if bytes.Compare(lock.PktBlockHash, pktBlockHash) != 0 {
		return errors.New("pkt block hash not correct")
	}

	return nil
}

// Is signature valid for sender address ?
func (t TokenStrikeMock) validateLockSignature(lock *lock.Lock) error {
	senderAddress, err := t.addressScheme.ParseAddr(lock.Sender)
	if err != nil {
		return err
	}

	sig, err := hex.DecodeString(lock.GetSignature())
	if err != nil {
		return err
	}

	lock.Signature = ""

	unsignedLock, err := proto.Marshal(lock)
	if err != nil {
		return err
	}

	isSigByIsaac := senderAddress.CheckSig(unsignedLock, sig)
	if !isSigByIsaac {
		return errors.New("it's not sign by sender")
	}

	return nil
}

func isContainToken(tokenName string, tokenSlice []string) bool {
	for _, token := range tokenSlice {
		if token == tokenName {
			return true
		}
	}

	return false
}
