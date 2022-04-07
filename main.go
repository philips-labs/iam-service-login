package main

import (
	"fmt"
	"os"

	"github.com/philips-software/go-hsdp-api/iam"
)

func main() {
	fmt.Printf("initializing from environment (%d)\n", len(os.Args))
	serviceID := os.Getenv("service_id")
	privateKey := os.Getenv("private_key")
	region := os.Getenv("region")
	environment := os.Getenv("environment")
	if len(os.Args) == 5 { // Read from command line instead
		fmt.Printf("initializing from command line arguments\n")
		region = os.Args[1]
		environment = os.Args[2]
		serviceID = os.Args[3]
		privateKey = os.Args[4]
	}
	fmt.Printf("service_id=%s\n", serviceID)

	client, err := iam.NewClient(nil, &iam.Config{
		Region:      region,
		Environment: environment,
	})
	if err != nil {
		fmt.Printf("::set-output name=message::NewClient: %v\n", err)
		os.Exit(1)
	}
	err = client.ServiceLogin(iam.Service{
		ServiceID:  serviceID,
		PrivateKey: privateKey,
	})
	if err != nil {
		fmt.Printf("::set-output name=message::ServiceLogin: %v\n", err)
		os.Exit(1)
	}
	token, err := client.Token()
	if err != nil {
		fmt.Printf("::set-output name=message::ServiceLogin: %v\n", err)
		return
	}
	fmt.Printf("::add-mask::%s\n", token)
	fmt.Printf("::set-output name=token::%s\n", token)
	fmt.Printf("::set-output name=message::Login success\n")
}
