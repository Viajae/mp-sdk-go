package main

import (
	"context"
	"fmt"
	"github.com/viajae/mp-sdk-go/pkg/config"
	"github.com/viajae/mp-sdk-go/pkg/order"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"
	orderID := "{{ORDER_ID}}"

	c, err := config.New(accessToken)
	if err != nil {
		fmt.Println("Error creating config:", err)
		return
	}

	client := order.NewClient(c)

	refundRequest := order.RefundRequest{
		Transactions: []order.RefundTransaction{
			{
				ID:     "{{TRANSACTION_ID}}",
				Amount: "25.00",
			},
		},
	}

	orderRefunded, err := client.Refund(context.Background(), orderID, &refundRequest)
	if err != nil {
		fmt.Println("Error in refund transaction:", err)
		return
	}

	fmt.Println("Partial refund completed successfully", orderRefunded)
}
