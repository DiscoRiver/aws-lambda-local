package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func HandleRequest(ctx context.Context) (events.APIGatewayProxyResponse, error) {
	TestZendeskParam := "TestZendesk"
	Decrypt := true
	parameter := ssm.GetParameterInput{
		Name:           &TestZendeskParam,
		WithDecryption: &Decrypt,
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := ssm.New(sess)

	paramOut, err := svc.GetParameter(&parameter)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: *paramOut.Parameter.Value, StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
