package MongoDBLanguage

import (
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func (e *MongoDBLanguage) Set(id, name string) (err error) {
	_, err = e.ClientLanguage.InsertOne(e.Ctx, bson.M{"_id": id, "name": name})
	if err != nil {
		log.Printf("MongoDBLanguage.Set().error: %v", err.Error())
		return
	}

	return
}
