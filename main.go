package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx *context.Context, Input interface{}) {
	log.Fatalf("Deployed to lambda successfully")
}

func main() {
	lambda.Start(handler)
}