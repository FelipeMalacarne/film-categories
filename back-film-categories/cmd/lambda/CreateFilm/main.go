package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/felipemalacarne/back-film-categories/internal/film/application/commands"
	"github.com/felipemalacarne/back-film-categories/internal/film/infrastructure/persistence"
	"github.com/felipemalacarne/back-film-categories/utils"
)

func main() {
	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var cmd commands.CreateFilmCommand

		if err := json.Unmarshal([]byte(request.Body), &cmd); err != nil {
			return utils.APIErrorResponse(http.StatusBadRequest, "Failed to unmarshal request body"), nil
		}

		db := dynamodb.New(session.Must(session.NewSession()))

		filmRepository := persistence.NewDynamoDBFilmRepository(db, "films")
		createFilmHandler := commands.NewCreateFilmHandler(filmRepository)

		supplier, err := createFilmHandler.Handle(cmd)
		if err != nil {
			return utils.APIErrorResponse(http.StatusBadRequest, err.Error()), nil
		}

		return utils.APISuccessResponse(supplier)
	})
}
