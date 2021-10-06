package replicator

type IssuerToken map[string][]string

func (i IssuerToken) GetToken(issuer string) []string {
	val, found := i[issuer]
	if !found {
		return []string{}
	}
	return val
}

func (i IssuerToken) AddToken(issuer, token string) {
	i[issuer] = append(i[issuer], token)
}
