package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	cp "github.com/felipemalacarne/back-film-categories/internal/category/infrastructure/persistence"
	"github.com/felipemalacarne/back-film-categories/internal/film/application/queries"
	fp "github.com/felipemalacarne/back-film-categories/internal/film/infrastructure/persistence"
	"github.com/felipemalacarne/back-film-categories/utils"
)

func main() {
	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var query queries.ListFilmsQuery

		db := dynamodb.New(session.Must(session.NewSession()))

		filmRepository := fp.NewDynamoDBFilmRepository(db, "films")
		categoryRepository := cp.NewDynamoDBCategoryRepository(db, "categories")
		listFilmsHandler := queries.NewListFilmsHandler(filmRepository, categoryRepository)

		films, err := listFilmsHandler.Handle(query)
		if err != nil {
			return utils.APIErrorResponse(http.StatusInternalServerError, err.Error()), nil
		}
		return utils.APISuccessResponse(films)
	})
}
