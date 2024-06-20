package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/felipemalacarne/back-prod-sup/internal/supplier/application/commands"
	"github.com/felipemalacarne/back-prod-sup/internal/supplier/infrastructure/persistence"
)

func LambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var cmd commands.CreateSupplierCommand
	if err := json.Unmarshal([]byte(request.Body), &cmd); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid request payload",
		}, nil
	}

	sess := session.Must(session.NewSession())
	db := dynamodb.New(sess)

	supplierRepository := persistence.NewDynamoDBSupplierRepository(db, "suppliers")
	createSupplierHandler := commands.NewCreateSupplierHandler(supplierRepository)

	supplier, err := createSupplierHandler.Handle(cmd)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	responseBody, err := supplier.MarshalJSON()
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(responseBody),
	}, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
