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
	"github.com/google/uuid"
)

func main() {
	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var cmd commands.UpdateCategoryCommand

		id, err := uuid.Parse(request.PathParameters["id"])
		if err != nil {
			return utils.APIErrorResponse(http.StatusBadRequest, "Invalid ID"), nil
		}

		if err := json.Unmarshal([]byte(request.Body), &cmd); err != nil {
			return utils.APIErrorResponse(http.StatusBadRequest, "Failed to unmarshal request body"), nil
		}

		cmd.ID = id

		db := dynamodb.New(session.Must(session.NewSession()))

		categoryRepository := persistence.NewDynamoDBCategoryRepository(db, "categories")
		updateCategoryHandler := commands.NewUpdateCategoryHandler(categoryRepository)

		category, err := updateCategoryHandler.Handle(cmd)
		if err != nil {
			return utils.APIErrorResponse(http.StatusBadRequest, err.Error()), nil
		}

		return utils.APISuccessResponse(category)
	})
}
