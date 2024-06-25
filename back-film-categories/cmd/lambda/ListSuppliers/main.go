package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/felipemalacarne/back-film-categories/utils"
	"github.com/felipemalacarne/back-film-categories/internal/supplier/application/queries"
	"github.com/felipemalacarne/back-film-categories/internal/supplier/infrastructure/persistence"
)

func LambdaHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var query queries.ListSuppliersQuery

	db := dynamodb.New(session.Must(session.NewSession()))

	supplierRepository := persistence.NewDynamoDBSupplierRepository(db, "suppliers")
	listSuppliersHandler := queries.NewListSuppliersHandler(supplierRepository)

	suppliers, err := listSuppliersHandler.Handle(query)
	if err != nil {
		return utils.APIErrorResponse(http.StatusInternalServerError, err.Error()), nil
	}

	return utils.APISuccessResponse(suppliers)
}

func main() {
	lambda.Start(LambdaHandler)
}
