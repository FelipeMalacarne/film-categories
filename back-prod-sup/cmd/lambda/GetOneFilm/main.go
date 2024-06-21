package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/felipemalacarne/back-prod-sup/internal/film/application/queries"
	"github.com/felipemalacarne/back-prod-sup/internal/film/infrastructure/persistence"
	"github.com/felipemalacarne/back-prod-sup/utils"
	"github.com/google/uuid"
)

func main() {
	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var query queries.GetOneFilmQuery

		id, err := uuid.Parse(request.PathParameters["id"])
		if err != nil {
			return utils.APIErrorResponse(http.StatusBadRequest, "Invalid ID"), nil
		}

		query.ID = id

		db := dynamodb.New(session.Must(session.NewSession()))

		filmRepository := persistence.NewDynamoDBFilmRepository(db, "films")
		getOneFilmHandler := queries.NewGetOneFilmHandler(filmRepository)

		film, err := getOneFilmHandler.Handle(query)
		if err != nil {
			return utils.APIErrorResponse(http.StatusBadRequest, err.Error()), nil
		}

		return utils.APISuccessResponse(film)
	})
}
