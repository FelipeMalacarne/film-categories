package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/felipemalacarne/back-prod-sup/internal/supplier/application/queries"
	"github.com/felipemalacarne/back-prod-sup/internal/supplier/infrastructure/persistence"
)

func LambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var query queries.ListSuppliersQuery

	db := dynamodb.New(session.Must(session.NewSession()))

	supplierRepository := persistence.NewDynamoDBSupplierRepository(db, "suppliers")
	listSuppliersHandler := queries.NewListSuppliersHandler(supplierRepository)

	suppliers, err := listSuppliersHandler.Handle(query)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	responseBody, err := json.Marshal(suppliers)
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
