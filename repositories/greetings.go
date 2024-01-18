package repositories

import (
	"context"
	"log"

	// "go.mongodb.org/mongo-driver/bson"
	"pertama_go/types"
	"pertama_go/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const collectionGreetings = "greetings"

func InsertGreetings(db *mongo.Database, greetings *types.Greetings) {
	ctx := context.Background()
	_, err := db.Collection(collectionGreetings).InsertOne(ctx, &greetings)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func InsertManyGreetings(db *mongo.Database, greetingss []interface{}) {
	ctx := context.Background()
	_, err := db.Collection(collectionGreetings).InsertMany(ctx, greetingss)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func FindOneGreetings(db *mongo.Database, query interface{}) (types.Greetings, error) {
	ctx := context.Background()
	csr := db.Collection(collectionGreetings).FindOne(ctx, query)
	if csr.Err() != nil {
		log.Fatal(csr.Err().Error())
	}

	var greetings types.Greetings
	err := csr.Decode(&greetings)
	if err != nil {
		log.Fatal(err.Error())
	}

	return greetings, nil
}

func FindAllGreetings(db *mongo.Database, query interface{}, pagination *types.Pagination) ([]types.Greetings, error) {
	ctx := context.Background()
	option := options.Find()
	offset := utils.GetOffset(pagination)
	option.SetLimit(int64(pagination.Limit))
	option.SetSkip(int64(offset))
	option.SetSort(bson.D{{Key: "created_at", Value: -1}})
	csr, err := db.Collection(collectionGreetings).Find(ctx, query, option)
	if err != nil {
		log.Fatal(err.Error())
	}

	var greetingss []types.Greetings
	err = csr.All(ctx, &greetingss)
	if err != nil {
		log.Fatal(err.Error())
	}

	return greetingss, nil
}

func DeleteOneGreetings(db *mongo.Database, query interface{}) {
	ctx := context.Background()
	_, err := db.Collection(collectionGreetings).DeleteOne(ctx, query)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func DeleteManyGreetings(db *mongo.Database, query interface{}) {
	ctx := context.Background()
	_, err := db.Collection(collectionGreetings).DeleteMany(ctx, query)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func CountGreetings(db *mongo.Database, query interface{}) int64 {
	ctx := context.Background()
	count, err := db.Collection(collectionGreetings).CountDocuments(ctx, query)
	if err != nil {
		log.Fatal(err.Error())
	}
	return count
}
