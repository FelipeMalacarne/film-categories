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
)

func main() {
	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var query queries.ListFilmsQuery

		db := dynamodb.New(session.Must(session.NewSession()))

		filmRepository := persistence.NewDynamoDBFilmRepository(db, "films")
		listFilmsHandler := queries.NewListFilmsHandler(filmRepository)

		films, err := listFilmsHandler.Handle(query)
		if err != nil {
			return utils.APIErrorResponse(http.StatusInternalServerError, err.Error()), nil
		}
		return utils.APISuccessResponse(films)
	})
}
