# Financial Cloud SDK for Go

fincloud-sdk-for-go provides Go packages for managing and using Financial Cloud services.
It officially supports the last two major releases of Go. Older versions of
Go will be kept running in CI until they no longer work due to changes in any
of the SDK's external dependencies. The CHANGELOG will be updated when a
version of Go is removed from CI.

## Installation

```bash
$ go get -u github.com/samjegal/fincloud-sdk-for-go
```

If you need to install Go, follow [the official instructions](https://golang.org/dl/).

## Authentication

Typical SDK operations must be authenticated and authorized. The _Authorizer_
interface allows use of any auth style in requests, such as inserting such as Access and Secret Keys
Authorization header token received.

```go
func authorize(sender autorest.Sender, serviceId string) autorest.Authorizer {
	httpMethod := "POST"
	requestUrl := "/sms/v2/services/" + serviceId + "/messages"

	auth := authentication.Builder{
		AccessKeyId: "{{ACCESS_KEY}}",
		SecretKey:   "{{SECRET_KEY}}",
		HttpMethod:  httpMethod,
		RequestURL:  requestUrl,
	}

	env, err := authentication.DetermineEnvironment("FINCLOUD")
	if err != nil {
		panic("Authentication determine set environment failed...")
	}

	config, err := auth.Build()
	if err != nil {
		panic("Authentication build failed...")
	}


	a, err := config.GetAuthorizationToken(sender, env.ResourceManagerEndpoint)
	if err != nil {
		panic("Get Authorization token failed...")
	}

	return a
}
```

## Example

Apply the following general steps to use packages in this repo. For more on authentication and the Authorizer interface see the next section.

1. Import a package from the services directory.
2. Create and authenticate a client with a New\*Client func, e.g. c := sens.NewSMSClientWithBaseURI(...).
3. Invoke API methods using the client, e.g. res, err := c.CreateOrUpdate(...).
4. Handle responses and errors.

For example, to create a new virtual network (substitute your own values for strings in angle brackets):

```go
package main

import (
	"context"
	"fmt"

	"github.com/Azure/go-autorest/autorest"
	"github.com/samjegal/fincloud-sdk-for-go/services/sens"
	"github.com/samjegal/go-fincloud-helpers/authentication"
	"github.com/samjegal/go-fincloud-helpers/sender"
)

func main()  {
	serviceId := "ncp:sms:fkr:256601028628:portal"

	sender := sender.BuildSender("FINCLOUD")
	client := sens.NewSMSClientWithBaseURI("https://sens.apigw.fin-ntruss.com")

	client.Authorizer = authorize(sender, serviceId)
	client.Sender = sender

	param := sens.SMSMessageRequestParameter{
		Type:            stringPtr("SMS"),
		CountryCode:     stringPtr("82"),
		From:            stringPtr("{{SENDER_PHONE_NUMBER}}"),
		Subject:         stringPtr("subject"),
		ContentType:     stringPtr("COMM"),
		Content:         stringPtr("sms"),
		Messages:        &[]sens.SMSMessageRequestParameterMessagesItem{
			{
				Subject: stringPtr("Test Subject"),
				Content: stringPtr("Test Message for SMS"),
				To:      stringPtr("{{OWN_PHONE_NUMBER}}"),
			},
		},
	}

	ctx := context.Background()
	resp, err := client.Request(ctx, serviceId, param)
	if err != nil {
		panic("Send request SMS faild...")
	}
}
```

## Contributing

This repository is inspired by Microsoft Azure SDK for Go. The code is based on [azure-sdk-for-go](https://github.com/Azure/azure-sdk-for-go)
