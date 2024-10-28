package repository

import (
	"context"
	"errors"
	"github.com/volkankocaali/bi-taksi-case/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Register(ctx context.Context, user domain.User) error
	Login(ctx context.Context, username string) (domain.User, error)
	CheckUserExists(ctx context.Context, username string) (bool, error)
}

type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(collection *mongo.Collection) *MongoUserRepository {
	return &MongoUserRepository{
		collection: collection,
	}
}

func (repo *MongoUserRepository) Register(ctx context.Context, user domain.User) error {
	_, err := repo.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (repo *MongoUserRepository) Login(ctx context.Context, username string) (domain.User, error) {
	var user domain.User

	err := repo.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.User{}, errors.New("user not found")
		}

		return domain.User{}, err
	}

	return user, nil
}

func (repo *MongoUserRepository) CheckUserExists(ctx context.Context, username string) (bool, error) {
	var user domain.User

	err := repo.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
