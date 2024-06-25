package persistence

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/felipemalacarne/back-film-categories/internal/supplier/domain"
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

func (r *DynamoDBSupplierRepository) Create(supplier *domain.Supplier) (domain.Supplier, error) {
	av, err := dynamodbattribute.MarshalMap(toDynamoSupplier(supplier))
	if err != nil {
		return domain.Supplier{}, err
	}

	input := &dynamodb.PutItemInput{
		TableName: &r.tableName,
		Item:      av,
	}

	_, err = r.db.PutItem(input)
	if err != nil {
		return domain.Supplier{}, err
	}

	return *supplier, nil
}

func (r *DynamoDBSupplierRepository) FindAll() ([]domain.Supplier, error) {
	scanInput := &dynamodb.ScanInput{
		TableName: &r.tableName,
	}

	result, err := r.db.Scan(scanInput)
	if err != nil {
		return nil, err
	}

	var dynamoSuppliers []dynamoSupplier
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &dynamoSuppliers)
	if err != nil {
		return nil, err
	}

	var suppliers []domain.Supplier
	for _, ds := range dynamoSuppliers {
		supplier := toSupplier(&ds)
		suppliers = append(suppliers, *supplier)
	}

	return suppliers, nil
}

func (r *DynamoDBSupplierRepository) FindByID(id uuid.UUID) (domain.Supplier, error) {
	keyCond := expression.Key("id").Equal(expression.Value(id.String()))

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

	var ds dynamoSupplier
	err = dynamodbattribute.UnmarshalMap(result.Items[0], &ds)
	if err != nil {
		return domain.Supplier{}, err
	}

	return *toSupplier(&ds), nil
}

func (r *DynamoDBSupplierRepository) Update(supplier *domain.Supplier) (domain.Supplier, error) {
	av, err := dynamodbattribute.MarshalMap(toDynamoSupplier(supplier))
	if err != nil {
		return domain.Supplier{}, err
	}

	input := &dynamodb.PutItemInput{
		TableName: &r.tableName,
		Item:      av,
	}

	_, err = r.db.PutItem(input)
	if err != nil {
		return domain.Supplier{}, err
	}

	return *supplier, nil
}

func (r *DynamoDBSupplierRepository) Delete(id uuid.UUID) error {
	key := map[string]*dynamodb.AttributeValue{
		"id": {
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
