package mongodb

import (
	"advanced-rest-yt/pkg/logging"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewClient(ctx context.Context, host, port, username, password, database, authDB string, logger *logging.Logger) (db *mongo.Database, err error) {
	uri := fmt.Sprintf("mongodb://%s:%s", host, port)
	opts := options.Client().ApplyURI(uri)

	if username != "" && password != "" {
		if authDB == "" {
			authDB = database
		}

		cred := options.Credential{
			AuthSource: authDB,
			Username:   username,
			Password:   password,
		}
		opts.SetAuth(cred)
	}

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		logger.Debugf("cant connect, err is %v", err)
		return nil, fmt.Errorf("cant connect to mongodb, err: %v", err)
	}

	err = client.Ping(ctx, &readpref.ReadPref{})
	if err != nil {
		logger.Debugf("cant ping mongodb, error is %v", err)
		return nil, fmt.Errorf("cant connect to mongodb, err: %v", err)
	}

	return client.Database(database), err
}
