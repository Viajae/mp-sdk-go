package main

import (
	"context"
	"fmt"

	"github.com/viajae/mp-sdk-go/pkg/config"
	"github.com/viajae/mp-sdk-go/pkg/invoice"
)

func main() {
	cfg, err := config.New("{{ACCESS_TOKEN}}")
	if err != nil {
		fmt.Println(err)
		return
	}

	client := invoice.NewClient(cfg)

	invoiceID := "123"

	resource, err := client.Get(context.Background(), invoiceID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resource)
}
