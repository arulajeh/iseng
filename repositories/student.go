package repositories

import (
	"context"
	"log"

	// "go.mongodb.org/mongo-driver/bson"
	"pertama_go/types"

	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "student"

func Insert(db *mongo.Database, student *types.Student) {
	ctx := context.Background()
	_, err := db.Collection(collectionName).InsertOne(ctx, &student)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func InsertMany(db *mongo.Database, students []interface{}) {
	ctx := context.Background()
	_, err := db.Collection(collectionName).InsertMany(ctx, students)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func FindOne(db *mongo.Database, query interface{}) (types.Student, error) {
	ctx := context.Background()
	csr := db.Collection(collectionName).FindOne(ctx, query)
	if csr.Err() != nil {
		log.Fatal(csr.Err().Error())
	}

	var student types.Student
	err := csr.Decode(&student)
	if err != nil {
		log.Fatal(err.Error())
	}

	return student, nil
}

func FindAll(db *mongo.Database, query interface{}) ([]types.Student, error) {
	ctx := context.Background()
	csr, err := db.Collection(collectionName).Find(ctx, query)
	if err != nil {
		log.Fatal(err.Error())
	}

	var students []types.Student
	err = csr.All(ctx, &students)
	if err != nil {
		log.Fatal(err.Error())
	}

	return students, nil
}

func DeleteOne(db *mongo.Database, query interface{}) {
	ctx := context.Background()
	_, err := db.Collection(collectionName).DeleteOne(ctx, query)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func DeleteMany(db *mongo.Database, query interface{}) {
	ctx := context.Background()
	_, err := db.Collection(collectionName).DeleteMany(ctx, query)
	if err != nil {
		log.Fatal(err.Error())
	}
}
