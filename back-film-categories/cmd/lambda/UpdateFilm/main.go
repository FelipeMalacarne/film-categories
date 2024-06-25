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
	"github.com/google/uuid"
)

func main() {
	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var cmd commands.UpdateFilmCommand

		id, err := uuid.Parse(request.PathParameters["id"])
        if err != nil {
            return utils.APIErrorResponse(http.StatusBadRequest, "Invalid ID"), nil
        }

		if err := json.Unmarshal([]byte(request.Body), &cmd); err != nil {
			return utils.APIErrorResponse(http.StatusBadRequest, "Failed to unmarshal request body"), nil
		}

        cmd.ID = id

		if cmd.Name == nil && cmd.Description == nil && cmd.Duration == nil && cmd.ReleaseDate == nil {
			return utils.APIErrorResponse(http.StatusBadRequest, "At least one field must be filled"), nil
		}

		db := dynamodb.New(session.Must(session.NewSession()))

		filmRepository := persistence.NewDynamoDBFilmRepository(db, "films")
		updateFilmHandler := commands.NewUpdateFilmHandler(filmRepository)

		film, err := updateFilmHandler.Handle(cmd)
		if err != nil {
			return utils.APIErrorResponse(http.StatusBadRequest, err.Error()), nil
		}

		return utils.APISuccessResponse(film)
	})
}
