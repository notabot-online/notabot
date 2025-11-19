package blockchain

import (
	"context"
)

type BlockchainService struct {
	rpcURL     string
	privateKey string
}

func New(rpcURL, privateKey string) *BlockchainService {
	return &BlockchainService{
		rpcURL:     rpcURL,
		privateKey: privateKey,
	}
}

func (b *BlockchainService) SubmitVerification(ctx context.Context, proof []byte) (string, error) {
	// expecting service
	return "", nil
}

func (b *BlockchainService) ListenForEvents(ctx context.Context) error {
	// expecting event bus
	return nil
}

