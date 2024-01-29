package main

import (
	"fmt"
	"log"
	"scandemo/pkg/api"
)

func main() {
	client, err := api.NewClient(api.DockerNetWork, api.DockerAddress)
	if err != nil {
		log.Println(err)
		return
	}
	history, err := client.GetHistory("nginx")
	if err != nil {
		log.Println(err)
		return
	}
	for _, h := range history {
		fmt.Printf("%+v\n", h)
	}

	image, err := client.GetImageMeta("nginx")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("%+v\n", *image)
}
