package kong

import "eternauta/pkg"

type KongClient struct {
	client pkg.IClient
}

var _ IKongClient = &KongClient{}

func NewAliasClient(client pkg.IClient) *KongClient {
	return &KongClient{
		client: client,
	}
}
