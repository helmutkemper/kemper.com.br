package MongoDB

import (
	"context"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Client     *mongo.Client
	Ctx        context.Context
	CancalFunc context.CancelFunc
}
