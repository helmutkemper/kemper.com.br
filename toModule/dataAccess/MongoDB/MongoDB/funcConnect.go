package MongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func (e *MongoDB) Connect(connectionString string, args ...interface{}) (err error) {
	e.Client, err = mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		return
	}

	e.Ctx = context.Background()
	err = e.Client.Connect(e.Ctx)
	if err != nil {
		return
	}

	err = e.Client.Ping(e.Ctx, readpref.Primary())
	return
}
