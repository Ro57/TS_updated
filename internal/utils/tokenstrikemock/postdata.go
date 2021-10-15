package tokenstrikemock

import (
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
		err error
	)

	switch req.Data.(type) {
	case *tokenstrike.Data_Block:
		block := req.GetBlock()
		err, resp.Warning = validateBlock(block)
	case *tokenstrike.Data_Lock:
		lock := req.GetLock()
		err, resp.Warning = validateLock(lock)
	default:
		return nil,  errors.New("unknown data type")
	}

	if err != nil {
		resp.Error = err.Error()
		return resp, err
	}

	return resp, nil
}

//todo place here checking for ret error with warnings
func validateBlock(block *DB.Block) (err error,warnings []string){
	return nil, nil
}

//todo place here checking for ret error with warnings
func validateLock(block *lock.Lock) (err error,warnings []string){
	return nil, nil
}
