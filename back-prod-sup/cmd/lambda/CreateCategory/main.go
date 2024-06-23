package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/felipemalacarne/back-prod-sup/internal/category/application/commands"
	"github.com/felipemalacarne/back-prod-sup/internal/category/infrastructure/persistence"
	"github.com/felipemalacarne/back-prod-sup/utils"
)

func main() {
	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var cmd commands.CreateCategoryCommand

		if err := json.Unmarshal([]byte(request.Body), &cmd); err != nil {
			return utils.APIErrorResponse(http.StatusBadRequest, "Failed to unmarshal request body"), nil
		}

		db := dynamodb.New(session.Must(session.NewSession()))

		categoryRepository := persistence.NewDynamoDBCategoryRepository(db, "categories")
		createCategoryHandler := commands.NewCreateCategoryHandler(categoryRepository)

		category, err := createCategoryHandler.Handle(cmd)
		if err != nil {
			return utils.APIErrorResponse(http.StatusBadRequest, err.Error()), nil
		}

		return utils.APISuccessResponse(category)
	})
}
