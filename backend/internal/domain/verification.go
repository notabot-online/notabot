package domain

import "time"

type VerificationStatus string

const (
	StatusPending   VerificationStatus = "pending"
	StatusVerified   VerificationStatus = "verified"
	StatusRejected   VerificationStatus = "rejected"
	StatusProcessing VerificationStatus = "processing"
)

type Verification struct {
	ID        string
	UserID    string
	Status    VerificationStatus
	Proof     []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}


