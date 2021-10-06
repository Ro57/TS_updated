package replicatorrpc

type IssuerTokens map[string][]string

func (i IssuerTokens) GetToken(issuer string) []string {
	val, found := i[issuer]
	if !found {
		return []string{}
	}
	return val
}

func (i IssuerTokens) AddToken(issuer, token string) {
	i[issuer] = append(i[issuer], token)
}
