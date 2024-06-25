package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	cp "github.com/felipemalacarne/back-prod-sup/internal/category/infrastructure/persistence"
	"github.com/felipemalacarne/back-prod-sup/internal/film/application/commands"
	fp "github.com/felipemalacarne/back-prod-sup/internal/film/infrastructure/persistence"
	"github.com/felipemalacarne/back-prod-sup/utils"
)

func main() {
	lambda.Start(func(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		var cmd commands.SetFilmCategoryCommand

		db := dynamodb.New(session.Must(session.NewSession()))
		filmRepository := fp.NewDynamoDBFilmRepository(db, "films")
		categoryRepository := cp.NewDynamoDBCategoryRepository(db, "categories")

		setFilmCategoryHandler := commands.NewSetFilmCategoryHandler(filmRepository, categoryRepository)

		film, err := setFilmCategoryHandler.Handle(cmd)
		if err != nil {
			return utils.APIErrorResponse(http.StatusInternalServerError, err.Error()), nil
		}

		return utils.APISuccessResponse(film)
	})
}
