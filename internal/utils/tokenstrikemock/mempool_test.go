package tokenstrikemock

import (
	"testing"

	"token-strike/tsp2p/server/DB"

	"google.golang.org/protobuf/runtime/protoiface"
)

func Test_mempoolImpl_Insert(t *testing.T) {

	type args struct {
		hash       string
		message    protoiface.MessageV1
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
		t.Run(tt.name, func(t *testing.T) {
			m := NewMempoolImpl()
			if got := m.Insert(tt.args.hash, tt.args.message, tt.args.expiration); got != tt.want {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
