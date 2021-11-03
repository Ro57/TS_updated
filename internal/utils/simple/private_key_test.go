package simple

import (
	"crypto/ed25519"
	"testing"
	"token-strike/internal/types/address"
)

func TestSimplePrivateKey_Equal(t *testing.T) {
	seed := randomSeed(123)
	seed2 := randomSeed(123213)

	type args struct {
		private address.PrivateKey
	}
	tests := []struct {
		name string
		Key  ed25519.PrivateKey
		args args
		want bool
	}{
		{
			name: "Valid key",
			Key:  ed25519.NewKeyFromSeed(seed[:]),
			args: args{
				private: SimplePrivateKey{ed25519.NewKeyFromSeed(seed[:])},
			},
			want: true,
		},
		{
			name: "Invalid key",
			Key:  ed25519.NewKeyFromSeed(seed[:]),
			args: args{
				private: SimplePrivateKey{ed25519.NewKeyFromSeed(seed2[:])},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := SimplePrivateKey{
				Key: tt.Key,
			}
			got := p.Equal(tt.args.private)
			if got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func randomSeed(offset int) [32]byte {
	byte32 := [32]byte{}

	for i := 0; i < 32; i++ {
		byte32[i] = byte(i + offset)
	}

	return byte32
}
