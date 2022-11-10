package users

import (
	"context"
	"dynamodb_go/models"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
)

type RepositoryDynamo interface {
	Store(ctx context.Context, model *models.User) error
	GetOne(ctx context.Context, id string) (*models.User, error)
	UpdateEmailAndPass(ctx context.Context, email string, pass string, id string) error
	Delete(ctx context.Context, id string) error
}

type dynamoRepository struct {
	dynamo *dynamodb.DynamoDB
	table  string
}

const (
	TABLE_NAME = "Users"
)

func NewDynamoRepository(dynamo *dynamodb.DynamoDB) RepositoryDynamo {
	return &dynamoRepository{
		dynamo: dynamo,
		table:  TABLE_NAME,
	}
}

func (r *dynamoRepository) Store(ctx context.Context, u *models.User) error {
	u.Id = uuid.New().String()
	av, err := dynamodbattribute.MarshalMap(u)

	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(r.table),
	}

	_, err = r.dynamo.PutItemWithContext(ctx, input)

	if err != nil {
		return err
	}

	return nil
}

func (r *dynamoRepository) GetOne(ctx context.Context, id string) (*models.User, error) {
	result, err := r.dynamo.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(r.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	return models.ItemToUser(result.Item)
}

func (r * dynamoRepository) UpdateEmailAndPass(ctx context.Context, email, pass, id string) error {

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":p":{
				S: aws.String(pass),
			},
            ":e": {
                S: aws.String(email),
            },
        },
		TableName: aws.String(r.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(id)},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set email = :e, password = :p"),
	}

	_, err := r.dynamo.UpdateItemWithContext(ctx, input)

	if err != nil {
		return err
	}

	return nil
}

func (r *dynamoRepository) Delete(ctx context.Context, id string) error {
	_, err := r.dynamo.DeleteItemWithContext(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(r.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(id)},
		},
	})

	if err != nil {
		return err
	}

	return nil
}
