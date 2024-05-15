package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type RetMsg struct {
	Message string `json:"message"`
}

func handler(ctx context.Context, Input interface{}) (RetMsg, error) {
	return RetMsg{
		Message: "Action Worked.",
	}, nil
}

func main() {
	lambda.Start(handler)
}