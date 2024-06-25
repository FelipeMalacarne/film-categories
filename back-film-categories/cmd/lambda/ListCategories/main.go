package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/felipemalacarne/back-film-categories/internal/category/application/queries"
	"github.com/felipemalacarne/back-film-categories/internal/category/infrastructure/persistence"
	"github.com/felipemalacarne/back-film-categories/utils"
)

func main() {
	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var query queries.ListCategoriesQuery

		db := dynamodb.New(session.Must(session.NewSession()))

		categoryRepository := persistence.NewDynamoDBCategoryRepository(db, "categories")
		listCategoriesHandler := queries.NewListCategoriesHandler(categoryRepository)

		categories, err := listCategoriesHandler.Handle(query)
		if err != nil {
			return utils.APIErrorResponse(http.StatusBadRequest, err.Error()), nil
		}

		return utils.APISuccessResponse(categories)
	})
}
