package utils_test

import (
	"fmt"
	idgen "token-strike/internal/utils/idgen"
)

func (suite *TestSuite) TestSign() {
	wantBlockHash := "testHash"
	wantJustificationNumber := 3
	wantID := fmt.Sprint(wantBlockHash, ":", wantJustificationNumber)

	id := idgen.Encode(wantBlockHash, wantJustificationNumber)
	suite.Equal(id, wantID)

	blockHash, number, err := idgen.Decode(id)
	suite.NoError(err)

	suite.Equal(*blockHash, wantBlockHash)
	suite.Equal(*number, wantJustificationNumber)

}
