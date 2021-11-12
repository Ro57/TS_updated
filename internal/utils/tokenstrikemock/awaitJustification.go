package tokenstrikemock

// func (t TokenStrikeMock) AwaitJustification(tokenId string, hashOfEntity []byte) (*string, error) {
// 	// Obtained after the method dispatch is called
// 	respBlock := <-t.blockDispatcher

// 	blockBytes, err := respBlock.Block.GetHash()
// 	if err != nil {
// 		return nil, err
// 	}

// 	blockHash := hex.EncodeToString(blockBytes)

// 	number, err := getJustificationIndex(respBlock.Block.Justifications, hashOfEntity)
// 	if err != nil {
// 		return nil, err
// 	}

// 	id := idgen.Encode(blockHash, number)

// 	return &id, nil
// }
