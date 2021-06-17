package main

import (
	"fmt"
	"os"

	"github.com/philips-software/go-hsdp-api/iam"
)

func main() {
	serviceID := os.Getenv("service_id")
	privateKey := os.Getenv("private_key")
	clientID := os.Getenv("client_id")
	clientSecret := os.Getenv("client_secret")
	region := os.Getenv("region")
	environment := os.Getenv("environment")

	fmt.Printf("client_id=%s\n", clientID)
	fmt.Printf("service_id=%s\n", serviceID)

	client, err := iam.NewClient(nil, &iam.Config{
		Region:         region,
		Environment:    environment,
		OAuth2ClientID: clientID,
		OAuth2Secret:   clientSecret,
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
	fmt.Printf("::add-mask::%s\n", client.Token())
	fmt.Printf("::set-output name=token::%s\n", client.Token())

}
