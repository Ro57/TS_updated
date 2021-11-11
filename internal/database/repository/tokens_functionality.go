package repository

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"token-strike/internal/database"
	"token-strike/internal/errors"
	"token-strike/internal/utils/idgen"
	"token-strike/tsp2p/server/DB"
	"token-strike/tsp2p/server/lock"
	"token-strike/tsp2p/server/replicator"

	"github.com/golang/protobuf/proto"
	"go.etcd.io/bbolt"
)

func (b *Bbolt) IssueTokenDB(tokenID string, tokenInfo *DB.Token, block *DB.Block, state *DB.State) error {
	return b.db.Update(func(tx *bbolt.Tx) error {
		rootBucket, err := tx.CreateBucketIfNotExists(database.TokensKey)
		if err != nil {
			return err
		}

		tokenBucket, err := rootBucket.CreateBucket([]byte(tokenID))
		if err != nil {
			return err
		}

		// put token info
		tokenBytes, err := proto.Marshal(tokenInfo)
		if err != nil {
			return err
		}

		err = tokenBucket.Put(database.InfoKey, tokenBytes)
		if err != nil {
			return err
		}

		// put state info
		marshaledState, err := proto.Marshal(state)
		if err != nil {
			return err
		}

		errPut := tokenBucket.Put(database.StateKey, marshaledState)
		if errPut != nil {
			return errPut
		}

		// put genesis state
		errPut = tokenBucket.Put(database.GenesisStateKey, marshaledState)
		if errPut != nil {
			return errPut
		}

		err = tokenBucket.Put(database.RootHashKey, []byte(""))
		if err != nil {
			return err
		}

		// TODO: remove after do refactoring
		//if string(tokenBucket.Get(database.RootHashKey)) != block.PrevBlock {
		//	return fmt.Errorf(
		//		"invalid hash of the previous block want %s but get %s",
		//		tokenBucket.Get(database.RootHashKey),
		//		block.PrevBlock,
		//	)
		//}

		// set root hash key
		blockHash, err := block.GetHash()
		if err != nil {
			return err
		}

		err = tokenBucket.Put(database.RootHashKey, blockHash)
		if err != nil {
			return err
		}

		// save block
		blockBytes, errMarshal := proto.Marshal(block)
		if errMarshal != nil {
			return errMarshal
		}

		err = tokenBucket.Put(database.GenesisBlockKey, blockBytes)
		if err != nil {
			return err
		}

		chainBucket, err := tokenBucket.CreateBucketIfNotExists(database.ChainKey)
		if err != nil {
			return err
		}

		return chainBucket.Put(blockHash, blockBytes)
	})
}

func (b *Bbolt) LockToken(tokenID string, lock *lock.Lock) error {
	return b.db.Update(func(tx *bbolt.Tx) error {
		rootBucket := tx.Bucket(database.TokensKey)
		if rootBucket == nil {
			return errors.RootBucketNotFoundErr
		}

		tokenBucket := rootBucket.Bucket([]byte(tokenID))
		if tokenBucket == nil {
			return errors.TokenNotFoundErr
		}

		stateBytes := tokenBucket.Get(database.StateKey)
		if stateBytes == nil {
			return errors.StateNotFoundErr
		}

		var state DB.State
		err := proto.Unmarshal(stateBytes, &state)
		if err != nil {
			return err
		}

		state.Locks = append(state.Locks, lock)

		senderIndex := state.GetOwnerIndexByHolder(lock.Sender, state.Owners)
		if senderIndex == nil {
			return fmt.Errorf("holder with name %v not found in state", lock.Sender)
		}

		state.Owners[*senderIndex].Count = state.Owners[*senderIndex].Count - lock.Count

		stateBytes, err = proto.Marshal(&state)
		if err != nil {
			return nil
		}

		return tokenBucket.Put(database.StateKey, stateBytes)
	})
}

func (b *Bbolt) ApplyJustification(tokenID string, justification *DB.Justification) error {
	switch j := justification.Content.(type) {
	case *DB.Justification_Lock:
		return b.LockToken(tokenID, j.Lock.Lock)
	case *DB.Justification_Transfer:
		return b.TransferTokens(tokenID, j.Transfer.Lock)
	default:
		return nil
	}
}

func (b *Bbolt) TransferTokens(tokenID, lockID string) error {
	return b.db.Update(func(tx *bbolt.Tx) error {
		rootBucket, err := tx.CreateBucketIfNotExists(database.TokensKey)
		if err != nil {
			return err
		}

		tokenBucket := rootBucket.Bucket([]byte(tokenID))
		if tokenBucket == nil {
			return fmt.Errorf("token with id %v not found", tokenID)
		}

		stateBytes := tokenBucket.Get(database.StateKey)
		if stateBytes == nil {
			return errors.StateNotFoundErr
		}

		chainBucket := tokenBucket.Bucket(database.ChainKey)
		if chainBucket == nil {
			return errors.ChainBucketNotFoundErr
		}

		blockHash, justificationIndex, err := idgen.Decode(lockID)
		if err != nil {
			return err
		}

		var state DB.State
		err = proto.Unmarshal(stateBytes, &state)
		if err != nil {
			return err
		}

		blockBytes := chainBucket.Get([]byte(*blockHash))
		if blockBytes == nil {
			return fmt.Errorf(
				"block not found by signature=%v",
				blockHash,
			)
		}

		var block DB.Block
		err = proto.Unmarshal(blockBytes, &block)
		if err != nil {
			return err
		}

		lockJustification, ok := block.Justifications[*justificationIndex].Content.(*DB.Justification_Lock)
		if !ok {
			return fmt.Errorf("lockFromState not found in blockL %v", block)
		}

		lockHash, err := lockJustification.Lock.Lock.GetHash()
		if err != nil {
			return err
		}

		lockHashIndex := state.GetLockIndexByHash(hex.EncodeToString(lockHash), state.Locks)
		if lockHashIndex == nil {
			return fmt.Errorf("not found lockFromState %v in state", lockID)
		}

		lockFromState := state.Locks[*lockHashIndex]
		recipientIndex := state.GetOwnerIndexByHolder(lockFromState.Recipient, state.Owners)
		if recipientIndex == nil {
			index := len(state.Owners)
			recipientIndex = &index
			state.Owners = append(state.Owners, &DB.Owner{HolderWallet: lockFromState.Recipient, Count: state.Locks[*lockHashIndex].Count})
		} else {
			// Remove lockFromState
			state.Locks = append(state.Locks[:*lockHashIndex], state.Locks[*lockHashIndex+1:]...)

			// Change balance
			state.Owners[*recipientIndex].Count = state.Owners[*recipientIndex].Count + lockFromState.Count
		}

		stateBytes, err = proto.Marshal(&state)
		if err != nil {
			return nil
		}

		return tokenBucket.Put(database.StateKey, stateBytes)
	})
}

func (b *Bbolt) SaveIssuerTokenDB(name, issuer string) error {
	return b.db.Update(func(tx *bbolt.Tx) error {
		rootBucket, err := tx.CreateBucketIfNotExists(database.TokensKey)
		if err != nil {
			return err
		}

		tokens := rootBucket.Get(database.IssuerTokens)
		if tokens == nil {
			tokens, _ = json.Marshal(replicator.IssuerToken{})
		}

		var issuerTokens replicator.IssuerToken
		errUnmarshal := json.Unmarshal(tokens, &issuerTokens)
		if errUnmarshal != nil {
			return errUnmarshal
		}

		issuerTokens.AddToken(issuer, name)

		issuerTokensBytes, errMarshal := json.Marshal(issuerTokens)
		if errMarshal != nil {
			return errMarshal
		}

		return rootBucket.Put(database.IssuerTokens, issuerTokensBytes)
	})
}
