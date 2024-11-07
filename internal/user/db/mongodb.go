package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"rest_api/internal/user"
	"rest_api/pkg/logging"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func (d *db) Create(ctx context.Context, user *user.User) (string, error) {
	d.logger.Debug("Create user")
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		d.logger.Errorf("error creating user: %v", err)
		return "", fmt.Errorf("error creating user: %v", err)
	}
	d.logger.Debug("Convert InsertedID obj to hex")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	d.logger.Trace(user)
	return "", fmt.Errorf("error failed to convert obj to hex: %s", oid)
}

func (d *db) FindOne(ctx context.Context, id string) (u *user.User, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		d.logger.Errorf("error finding user: %v, HEX: %s", err, id)
		return nil, fmt.Errorf("error finding user: %v, HEX: %s", err, id)
	}
	filter := bson.M{"_id": oid}
	d.collection.FindOne(ctx, filter)
}

func (d *db) Update(ctx context.Context, user *user.User) error {
	//TODO implement me
	panic("implement me")
}

func (d *db) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
