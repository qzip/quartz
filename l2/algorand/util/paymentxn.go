package util

import (
	"context"
	"log"
)

func CreatePaymentTxn(ctx context.Context) {
	sp, err := algodClient.SuggestedParams().Do(context.Background())
	if err != nil {
		log.Fatalf("failed to get suggested params: %s", err)
	}
}
