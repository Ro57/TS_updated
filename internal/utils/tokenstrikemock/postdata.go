package tokenstrikemock

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/lock"
	"token-strike/tsp2p/server/tokenstrike"

	"github.com/golang/protobuf/proto"
)

func (t TokenStrikeMock) PostData(ctx context.Context, req *tokenstrike.Data) (*tokenstrike.PostDataResp, error) {
	if req == nil {
		return nil, errors.New("nil data")
	}

	resp := &tokenstrike.PostDataResp{}
	lockEl := &lock.Lock{}
	blockEl := &DB.Block{}
	transferEl := &tokenstrike.TransferTokens{}

	var err error

	switch req.Data.(type) {
	case *tokenstrike.Data_Block:
		blockEl = req.GetBlock()
		resp.Warning, err = t.validateBlock(blockEl)
	case *tokenstrike.Data_Lock:
		lockEl = req.GetLock()
		resp.Warning, err = t.validateLock(*lockEl)
	case *tokenstrike.Data_Transfer:
		transferEl = req.GetTransfer()
		resp.Warning, err = t.validateTransfer(*transferEl)
	default:
		return nil, errors.New("unknown data type")
	}
	if err != nil {
		return nil, err
	}

	return resp, t.sendDataToSubscribers(req)
}

//TODO: place here checking for ret error with warnings
func (t TokenStrikeMock) validateBlock(block *DB.Block) (warnings []string, err error) {
	err = t.validateBlockInv(block)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

//TODO: place here checking for ret error with warnings
func (t TokenStrikeMock) validateTransfer(transfer tokenstrike.TransferTokens) (warnings []string, err error) {
	validatorErrors := []error{
		t.validateTransferLock(transfer),
		t.validateTransferHtlc(transfer),
	}

	for _, err := range validatorErrors {
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

//TODO: place here checking for ret error with warnings
func validateTransfer(block *DB.Block) (warnings []string, err error) {

	return nil, nil
}

//TODO: place here checking for ret warnings
func (t TokenStrikeMock) validateLock(reqLock lock.Lock) (warnings []string, err error) {
	validatorErrors := []error{
		t.validateLockInv(&reqLock),
		t.validateLockSignature(reqLock),
		t.validateLockHashCorrect(reqLock),
		t.validateLockIssuer(reqLock),
		t.validateLockHeight(reqLock),
		t.validateLockPktHeight(reqLock),
		t.validateLockSenderOwnedTokens(reqLock),
	}

	for _, err := range validatorErrors {
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

// Is token issued by issuer ?
func (t TokenStrikeMock) validateLockIssuer(lock lock.Lock) error {
	lockBytes, err := proto.Marshal(&lock)
	if err != nil {
		return err
	}

	tokenID := t.getTokenID(lockBytes)

	issuerStore, err := t.bboltDB.GetIssuerTokens()
	if err != nil {
		return err
	}

	tokenSlice := issuerStore[t.issuer.String()]

	if !isContainToken(tokenID, tokenSlice) {
		return errors.New("token not in storage")
	}

	return nil
}

// Is pkt block height within 10 blocks old, not in the future ?
func (t TokenStrikeMock) validateLockHeight(lock lock.Lock) error {
	curHeight := t.pktChain.CurrentHeight()

	if int32(lock.PktBlockHeight) <= (curHeight - 10) {
		return errors.New("pkt block height too old")
	}

	return nil
}

//Is pkt block hash correct ?
func (t TokenStrikeMock) validateLockHashCorrect(lock lock.Lock) error {
	pktBlockHash := t.pktChain.BlockHashAtHeight(int32(lock.PktBlockHeight))

	if bytes.Compare(lock.PktBlockHash, pktBlockHash) != 0 {
		return errors.New("pkt block hash not correct")
	}

	return nil
}

// Is signature valid for sender address ?
func (t TokenStrikeMock) validateLockSignature(lock lock.Lock) error {
	senderAddress, err := t.addressScheme.ParseAddr(lock.Sender)
	if err != nil {
		return err
	}

	sig, err := hex.DecodeString(lock.GetSignature())
	if err != nil {
		return err
	}

	unsignedLock := lock
	unsignedLock.Signature = ""

	unsignedLockBytes, err := proto.Marshal(&unsignedLock)
	if err != nil {
		return err
	}

	isSigByIsaac := senderAddress.CheckSig(unsignedLockBytes, sig)
	if !isSigByIsaac {
		return errors.New("it's not sign by sender")
	}

	return nil
}

// Is expires_pkt_height at least 2 blocks in the future ?
func (t TokenStrikeMock) validateLockPktHeight(lock lock.Lock) error {
	if lock.ProofCount-t.pktChain.CurrentHeight() < 2 {
		return errors.New("lock expired")
	}

	return nil
}

// Does sender have count tokens it their possession ?
func (t TokenStrikeMock) validateLockSenderOwnedTokens(lock lock.Lock) error {
	lockBytes, err := proto.Marshal(&lock)
	if err != nil {
		return err
	}

	tokenID := t.getTokenID(lockBytes)

	chain, err := t.bboltDB.GetChainInfoDB(tokenID)
	if err != nil {
		return err
	}

	senderInfo := getOwner(lock.Sender, chain.State.Owners)

	if senderInfo.Count < lock.Count {
		return fmt.Errorf("Sender %v has %v token(s) but tried to lock %v", senderInfo.HolderWallet, senderInfo.Count, lock.Count)
	}

	return nil
}

// Does the lock exist?
func (t TokenStrikeMock) validateTransferLock(transfer tokenstrike.TransferTokens) error {
	transferBytes, err := proto.Marshal(&transfer)
	if err != nil {
		return err
	}

	tokenID := t.getTokenID(transferBytes)

	chain, err := t.bboltDB.GetChainInfoDB(tokenID)
	if err != nil {
		return err
	}

	if _, err := getLock(transfer.Lock, chain.State.Locks); err != nil {
		return err
	}

	return nil
}

// Is the htlc secret correct for the hash?
func (t TokenStrikeMock) validateTransferHtlc(transfer tokenstrike.TransferTokens) error {
	transferBytes, err := proto.Marshal(&transfer)
	if err != nil {
		return err
	}

	tokenID := t.getTokenID(transferBytes)

	chain, err := t.bboltDB.GetChainInfoDB(tokenID)
	if err != nil {
		return err
	}

	lock, err := getLock(transfer.Lock, chain.State.Locks)
	if err != nil {
		return err
	}

	htlcSingleHash := sha256.Sum256(transfer.Htlc)
	htlcDoubleHash := sha256.Sum256(htlcSingleHash[:])
	htlcHash := hex.EncodeToString(htlcDoubleHash[:])

	if lock.HtlcSecretHash != htlcHash {
		return errors.New("htlc secret incorrect")
	}
	return nil
}

// Type of Inv correct?
func (t TokenStrikeMock) validateLockInv(checkedLock *lock.Lock) error {
	lockHash, err := proto.Marshal(checkedLock)
	if err != nil {
		return err
	}

	inv := t.getInv(lockHash)
	if inv.Type != tokenstrike.TYPE_LOCK {
		return fmt.Errorf("type of justification want lock(1) but get %v", inv.Type)
	}

	return nil

}

// Type of Inv correct?
func (t TokenStrikeMock) validateBlockInv(block *DB.Block) error {
	blockHash, err := proto.Marshal(block)
	if err != nil {
		return err
	}

	inv := t.getInv(blockHash)
	if inv.Type != tokenstrike.TYPE_BLOCK {
		return fmt.Errorf("type of justification want block(2) but get %v", inv.Type)
	}

	return nil

}

func (t TokenStrikeMock) getInv(data []byte) tokenstrike.Inv {
	dataHash := sha256.Sum256(data)
	entity := hex.EncodeToString(dataHash[:])

	inv := t.invCache[entity]

	return inv
}

func (t TokenStrikeMock) getTokenID(data []byte) string {
	dataHash := sha256.Sum256(data)
	entity := hex.EncodeToString(dataHash[:])

	inv := t.invCache[entity]

	return hex.EncodeToString(inv.Parent)
}

func getOwner(ownerName string, owners []*DB.Owner) *DB.Owner {
	for _, owner := range owners {
		if owner.HolderWallet == ownerName {
			return owner
		}
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

func getLock(lockHash []byte, lockSlice []*lock.Lock) (*lock.Lock, error) {
	for _, lock := range lockSlice {
		lockBytes, err := proto.Marshal(lock)
		if err != nil {
			return nil, err
		}

		curLockHash := sha256.Sum256(lockBytes)
		if bytes.Equal(curLockHash[:], lockHash) {
			return lock, nil
		}
	}

	return nil, errors.New("lock not found")
}
