package utils_test

import (
	"testing"

	"token-strike/tsp2p/server/DB"

	"google.golang.org/protobuf/runtime/protoiface"
)

func (suite *TestSuite) TestInsert() {

	type args struct {
		hash       string
		message    protoiface.MessageV1
		msgType    uint32
		expiration int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "valid test",
			args: args{
				hash:       "test Hash",
				message:    &DB.State{},
				expiration: 12312312,
			},
			want: "dGVzdCBIYXNoL1wvXDEyMzEyMzEy",
		},
	}
	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {

			if got := suite.tokenStrike.Insert(
				tt.args.hash,
				tt.args.msgType,
				tt.args.message,
				tt.args.expiration,
			); got != tt.want {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
