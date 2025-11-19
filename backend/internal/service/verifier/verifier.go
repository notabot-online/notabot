package verifier

import (
	"context"
)

type Verifier struct {
	// TODO: add dependenciess
}

func New() *Verifier {
	return &Verifier{}
}

func (v *Verifier) VerifyProof(ctx context.Context, proof []byte) (bool, error) {
	// TODO: zk verification
	return false, nil
}

