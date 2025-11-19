package kyc

import (
	"context"
)

type KYCService struct {
	apiKey string
}

func New(apiKey string) *KYCService {
	return &KYCService{
		apiKey: apiKey,
	}
}

func (k *KYCService) VerifyUser(ctx context.Context, userID string) (bool, error) {
	// TODO: expecting KYC labs confirm
	return false, nil
}
