package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/felipemalacarne/back-film-categories/utils"
	"github.com/felipemalacarne/back-film-categories/internal/supplier/application/commands"
	"github.com/felipemalacarne/back-film-categories/internal/supplier/infrastructure/persistence"
)

func main() {
	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var cmd commands.CreateSupplierCommand

		if err := json.Unmarshal([]byte(request.Body), &cmd); err != nil {
			return utils.APIErrorResponse(http.StatusBadRequest, "Failed to unmarshal request body"), nil
		}

		sess := session.Must(session.NewSession())
		db := dynamodb.New(sess)

		supplierRepository := persistence.NewDynamoDBSupplierRepository(db, "suppliers")
		createSupplierHandler := commands.NewCreateSupplierHandler(supplierRepository)

		supplier, err := createSupplierHandler.Handle(cmd)
		if err != nil {
			return utils.APIErrorResponse(http.StatusInternalServerError, err.Error()), nil
		}

		return utils.APISuccessResponse(supplier)
	})
}
