package main

import (
	"context"
	"fmt"

	"github.com/viajae/mp-sdk-go/pkg/config"
	"github.com/viajae/mp-sdk-go/pkg/customer"
)

func main() {
	accessToken := "{{ACCESS_TOKEN}}"

	cfg, err := config.New(accessToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := customer.NewClient(cfg)

	request := customer.SearchRequest{
		Filters: map[string]string{
			"EMAIL": "{{EMAIL}}",
		},
	}

	resource, err := client.Search(context.Background(), request)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resource)
}
