package lightningnode

import "token-strike/internal/blockchainnode"

var _ blockchainnode.BlockchainNode = Node{}

type Node struct {
	client interface{} //todo change it from void interface to lnd client
}


func New() *Node{
	return &Node{}
}

func (n *Node) Connect(){
	//todo make this sa
}