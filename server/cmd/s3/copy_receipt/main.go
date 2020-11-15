package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/homma509/9rece/server/log"
	"github.com/homma509/9rece/server/registry"
	"golang.org/x/xerrors"
)

func handler(ctx context.Context, event events.S3Event) error {
	log.AppLogger.Info(
		"start lambda function",
		"S3Event", event,
	)

	err := registry.Creater().ReceiptController().Store(ctx, event)
	if err != nil {
		err = xerrors.Errorf("on handler: %w", err)
		log.AppLogger.Error(
			"error lambda function",
			"Result", "failure",
			"Error", err,
		)
		return err
	}

	log.AppLogger.Info(
		"end lambda function",
		"Result", "successful",
	)

	return nil
}

func main() {
	lambda.Start(handler)
}
