package middleware
import (
	"github.com/pmccau/TriviaGo/server/datamanagement"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

// GenerateQuestionBson is a helper function to generate BSON format question data for entry into db
func GenerateQuestionBson(text string, answer string, tags []string) bson.D {
	output := bson.D{
		{Key: "text", Value: text},
		{Key: "answer", Value: answer},
		{Key: "tags", Value: bson.A{tags}},
  }
  return output
}

// GetDb is a helper func to get db
func GetDb(name string, client *mongo.Client) *mongo.Database {
	return client.Database(name)
}

// GetCollection is a helper func to get collection
func GetCollection(collectionName string, db *mongo.Database) *mongo.Collection {
	return db.Collection(collectionName)
}

// PingClient is used as a debugging function to ping the client to test connection
func PingClient(ctx context.Context, client *mongo.Client) {
	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("SUCCESS!")
	}
}

// AddDocument adds a single document to a collection
func AddDocument(ctx context.Context, collection *mongo.Collection, document bson.D) error {
	result, err := collection.InsertOne(ctx, document)
	datamanagement.Check(err)
	fmt.Println(result)
	return err
}

// GetAllDocuments retrieves all documents within a collection and save them to an array
func GetAllDocuments(ctx context.Context, collection *mongo.Collection) []bson.M {
	var saveArr []bson.M
	cursor, err := collection.Find(ctx, bson.M{})
	datamanagement.Check(err)
	err = cursor.All(ctx, &saveArr)
	datamanagement.Check(err)
	return saveArr
}

// ConnectToMongo establishes connection to mongo DB and return the client/context
func ConnectToMongo(uri string) (*mongo.Client, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	datamanagement.Check(err)
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	err = client.Connect(ctx)
	datamanagement.Check(err)
	PingClient(ctx, client)
	return client, ctx
}