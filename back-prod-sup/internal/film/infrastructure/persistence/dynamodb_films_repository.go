package persistence

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/felipemalacarne/back-prod-sup/internal/film/domain"
	"github.com/google/uuid"
)

type DynamoDBFilmRepository struct {
	db        *dynamodb.DynamoDB
	tableName string
}

func NewDynamoDBFilmRepository(db *dynamodb.DynamoDB, tableName string) *DynamoDBFilmRepository {
	return &DynamoDBFilmRepository{
		db:        db,
		tableName: tableName,
	}
}

func (r *DynamoDBFilmRepository) Create(film *domain.Film) (domain.Film, error) {
	av, err := dynamodbattribute.MarshalMap(toDynamoFilm(film))
	if err != nil {
		return domain.Film{}, err
	}
	input := &dynamodb.PutItemInput{
		TableName: &r.tableName,
		Item:      av,
	}
	_, err = r.db.PutItem(input)
	if err != nil {
		return domain.Film{}, err
	}
	return *film, nil
}

func (r *DynamoDBFilmRepository) FindAll() ([]domain.Film, error) {
	scanInput := &dynamodb.ScanInput{
		TableName: &r.tableName,
	}
	result, err := r.db.Scan(scanInput)
	if err != nil {
		return nil, err
	}
	var dynamoFilms []dynamoFilm
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &dynamoFilms)
	if err != nil {
		return nil, err
	}
	var films []domain.Film
	for _, df := range dynamoFilms {
		films = append(films, *toFilm(&df))
	}
	return films, nil
}

func (r *DynamoDBFilmRepository) FindByID(id uuid.UUID) (domain.Film, error) {
	input := &dynamodb.GetItemInput{
		TableName: &r.tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id.String()),
			},
		},
	}
	result, err := r.db.GetItem(input)
	if err != nil {
		return domain.Film{}, err
	}
	if result.Item == nil {
		return domain.Film{}, nil
	}
	var df dynamoFilm
	err = dynamodbattribute.UnmarshalMap(result.Item, &df)
	if err != nil {
		return domain.Film{}, err
	}
	return *toFilm(&df), nil
}

func  (r *DynamoDBFilmRepository) Update(film *domain.Film) (domain.Film, error) {
    av, err := dynamodbattribute.MarshalMap(toDynamoFilm(film))
    if err != nil {
        return domain.Film{}, err
    }
    input := &dynamodb.PutItemInput{
        TableName: &r.tableName,
        Item:      av,
    }
    _, err = r.db.PutItem(input)
    if err != nil {
        return domain.Film{}, err
    }
    return *film, nil
}

func (r *DynamoDBFilmRepository) Delete(id uuid.UUID) error {
    input := &dynamodb.DeleteItemInput{
        TableName: &r.tableName,
        Key: map[string]*dynamodb.AttributeValue{
            "id": {
                S: aws.String(id.String()),
            },
        },
    }
    _, err := r.db.DeleteItem(input)
    return err
}
