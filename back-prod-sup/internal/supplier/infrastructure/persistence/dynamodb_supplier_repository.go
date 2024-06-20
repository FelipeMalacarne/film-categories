package persistence

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/felipemalacarne/back-prod-sup/internal/supplier/domain"
	"github.com/google/uuid"
)

type DynamoDBSupplierRepository struct {
	db        *dynamodb.DynamoDB
	tableName string
}

func NewDynamoDBSupplierRepository(db *dynamodb.DynamoDB, tableName string) *DynamoDBSupplierRepository {
	return &DynamoDBSupplierRepository{
		db:        db,
		tableName: tableName,
	}
}

func (r DynamoDBSupplierRepository) Create(supplier *domain.Supplier) (domain.Supplier, error) {
	av, err := dynamodbattribute.MarshalMap(supplier)
	if err != nil {
		return domain.Supplier{}, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: &r.tableName,
	}

	_, err = r.db.PutItem(input)
	if err != nil {
		return domain.Supplier{}, err
	}

	return *supplier, nil
}

func (r DynamoDBSupplierRepository) FindAll() ([]domain.Supplier, error) {
	scanInput := &dynamodb.ScanInput{
		TableName: &r.tableName,
	}

	result, err := r.db.Scan(scanInput)
	if err != nil {
		return nil, err
	}

	var suppliers []domain.Supplier
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &suppliers)
	if err != nil {
		return nil, err
	}

	return suppliers, nil
}

func (r DynamoDBSupplierRepository) FindByID(id uuid.UUID) (domain.Supplier, error) {
	keyCond := expression.Key("ID").Equal(expression.Value(id.String()))

	expr, err := expression.NewBuilder().WithKeyCondition(keyCond).Build()
	if err != nil {
		return domain.Supplier{}, err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 &r.tableName,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition(),
	}

	result, err := r.db.Query(queryInput)
	if err != nil {
		return domain.Supplier{}, err
	}

	if len(result.Items) == 0 {
		return domain.Supplier{}, errors.New("supplier not found")
	}

	var supplier domain.Supplier
	err = dynamodbattribute.UnmarshalMap(result.Items[0], &supplier)
	if err != nil {
		return domain.Supplier{}, err
	}

	return supplier, nil
}

func (r DynamoDBSupplierRepository) Update(supplier *domain.Supplier) (domain.Supplier, error) {
	av, err := dynamodbattribute.MarshalMap(supplier)
	if err != nil {
		return domain.Supplier{}, err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: &r.tableName,
	}

	_, err = r.db.PutItem(input)
	if err != nil {
		return domain.Supplier{}, err
	}

	return *supplier, nil
}

func (r DynamoDBSupplierRepository) Delete(id uuid.UUID) error {
	key := map[string]*dynamodb.AttributeValue{
		"ID": {
			S: aws.String(id.String()),
		},
	}

	input := &dynamodb.DeleteItemInput{
		Key:       key,
		TableName: &r.tableName,
	}

	_, err := r.db.DeleteItem(input)
	return err
}
