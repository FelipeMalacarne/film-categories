package persistence

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/felipemalacarne/back-film-categories/internal/category/domain"
	"github.com/google/uuid"
)

type DynamoDBCategoryRepository struct {
	db        *dynamodb.DynamoDB
	tableName string
}

func NewDynamoDBCategoryRepository(db *dynamodb.DynamoDB, tableName string) *DynamoDBCategoryRepository {
	return &DynamoDBCategoryRepository{
		db:        db,
		tableName: tableName,
	}
}

func (r *DynamoDBCategoryRepository) Create(category *domain.Category) (domain.Category, error) {
	av, err := dynamodbattribute.MarshalMap(toDynamoCategory(category))
	if err != nil {
		return domain.Category{}, err
	}
	input := &dynamodb.PutItemInput{
		TableName: &r.tableName,
		Item:      av,
	}
	_, err = r.db.PutItem(input)
	if err != nil {
		return domain.Category{}, err
	}
	return *category, nil
}

func (r *DynamoDBCategoryRepository) FindAll() ([]domain.Category, error) {
	scanInput := &dynamodb.ScanInput{
		TableName: &r.tableName,
	}
	result, err := r.db.Scan(scanInput)
	if err != nil {
		return nil, err
	}
	var dynamoCategories []dynamoCategory
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &dynamoCategories)
	if err != nil {
		return nil, err
	}
	var categories []domain.Category
	for _, dc := range dynamoCategories {
		categories = append(categories, *toCategory(&dc))
	}
	return categories, nil
}

func (r *DynamoDBCategoryRepository) FindByID(id uuid.UUID) (domain.Category, error) {
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
		return domain.Category{}, err
	}
	if result.Item == nil {
		return domain.Category{}, nil
	}
	var dc dynamoCategory
	err = dynamodbattribute.UnmarshalMap(result.Item, &dc)
	if err != nil {
		return domain.Category{}, err
	}
	return *toCategory(&dc), nil
}

func (r *DynamoDBCategoryRepository) Update(category *domain.Category) (domain.Category, error) {
	av, err := dynamodbattribute.MarshalMap(toDynamoCategory(category))
	if err != nil {
		return domain.Category{}, err
	}
	input := &dynamodb.PutItemInput{
		TableName: &r.tableName,
		Item:      av,
	}
	_, err = r.db.PutItem(input)
	if err != nil {
		return domain.Category{}, err
	}
	return *category, nil
}

func (r *DynamoDBCategoryRepository) Delete(id uuid.UUID) error {
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
