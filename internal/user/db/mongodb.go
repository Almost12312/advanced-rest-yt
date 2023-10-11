package db

import (
	"advanced-rest-yt/internal/apperror"
	"advanced-rest-yt/internal/user"
	"advanced-rest-yt/pkg/logging"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func NewStorage(database *mongo.Database, collectionName string, logger *logging.Logger) user.Storage {
	return &db{
		collection: database.Collection(collectionName),
		logger:     logger,
	}
}

func (d *db) Create(ctx context.Context, user user.User) (userId string, err error) {
	d.logger.Debug("create user")
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		d.logger.Infof("cant get user, err: %v", err)
		return userId, fmt.Errorf("failed to create user")
	}

	d.logger.Debug("convert InsertedID to primitive.ObjectID")
	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		d.logger.Trace(user)
		d.logger.Tracef("cant type assert, objectID: %s error: %v", id, err)
		return userId, fmt.Errorf("failed to type assertion, error: %v", err)
	}

	return id.Hex(), err
}

func (d *db) FindAll(ctx context.Context) (u []user.User, err error) {
	cur, err := d.collection.Find(ctx, bson.M{})
	//err = result.Err()
	if cur.Err() != nil {
		return u, fmt.Errorf("cant find all users. error is: %v", err)
	}

	err = cur.All(ctx, &u)
	if err != nil {
		return u, fmt.Errorf("cant decode all users. error is: %v", err)
	}

	return u, nil
}

func (d *db) FindOne(ctx context.Context, id string) (u user.User, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("failed to convert objID to hex. hex is: %s, error is: %v", id, err)
	}

	filter := bson.M{"_id": oid}
	result := d.collection.FindOne(ctx, filter)

	err = result.Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return u, apperror.ErrNotFound
		} else {
			return u, fmt.Errorf("cant find user. hex is: %s, error is: %v", id, err)
		}
	}

	err = result.Decode(&u)
	if err != nil {
		return u, fmt.Errorf("cant decode user. hex is: %s, error is: %v", id, err)
	}

	return u, nil

}

func (d *db) Update(ctx context.Context, user user.User) (err error) {
	objectID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return fmt.Errorf("cant convert hex to id, userID: %s", user.ID)
	}

	filter := bson.M{"_id": objectID}

	userBytes, err := bson.Marshal(user)
	if err != nil {
		return fmt.Errorf("cant marshal user, error: %v", err)
	}

	var updateUserObj bson.M
	err = bson.Unmarshal(userBytes, &updateUserObj)
	if err != nil {
		return fmt.Errorf("cant unmarshal user, error: %v", err)
	}

	delete(updateUserObj, "_id")

	update := bson.M{"$set": updateUserObj}

	result, err := d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("cant update user, userID: %s, error: %s", user.ID, err)
	}

	if result.MatchedCount == 0 {
		return apperror.ErrNotFound
	}

	d.logger.Tracef("Matched %d docs, modify %d", result.MatchedCount, result.ModifiedCount)

	return err
}

func (d *db) Delete(ctx context.Context, id string) (err error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("cant convert to objID, hex: %s, error: %v", id, err)
	}

	filter := bson.M{"_id": objID}

	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("cant delete, objID: %s, error: %v", objID, err)
	}

	if result.DeletedCount == 0 {
		return apperror.ErrNotFound
	}

	d.logger.Tracef("Deleted %d docs", result.DeletedCount)
	return err
}
