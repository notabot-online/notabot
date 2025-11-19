package domain

import "time"

type Transaction struct {
	ID            string
	VerificationID string
	TxHash        string
	Status        string
	CreatedAt     time.Time
}


