package main

import (
	"context"
	"fmt"

	"github.com/viajae/mp-sdk-go/pkg/config"
	"github.com/viajae/mp-sdk-go/pkg/preference"
)

func main() {
	cfg, err := config.New("{{ACCESS_TOKEN}}")
	if err != nil {
		fmt.Println(err)
		return
	}

	client := preference.NewClient(cfg)

	preferenceID := "123"

	resource, err := client.Get(context.Background(), preferenceID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resource)
}
