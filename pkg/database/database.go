package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DBNAME = "test"

//ConnectToDB Establishes connection between user and MongoDB
func ConnectToDB() (client *mongo.Client) {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("[-] Error while connecting to the database: %v", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("[-] Ping error: %v", err)
	}

	fmt.Println("[+] Succesfully connected to the database.")

	return client
}

//CloseConnectionToDB closes the connection between client and MongoDB
func CloseConnectionToDB(client *mongo.Client) {

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatalf("Error while closing connection to the database: %v", err)
	}

	fmt.Println("[*] Connection to the database has been closed.")
}

func InsertAccountToTheDatabase(client *mongo.Client, ctx context.Context, data AccountItem) (*mongo.InsertOneResult, error) {

	collection := client.Database(DBNAME).Collection("accounts")
	insertResult, err := collection.InsertOne(ctx, data)
	return insertResult, err
}

func GetAccountFromDB(client *mongo.Client, ctx context.Context, oid primitive.ObjectID) *mongo.SingleResult {

	collection := client.Database(DBNAME).Collection("accounts")
	result := collection.FindOne(ctx, bson.M{"_id": oid})
	return result
}

func DeleteAccountFromDB(client *mongo.Client, ctx context.Context, oid primitive.ObjectID) error {

	collection := client.Database(DBNAME).Collection("accounts")
	result, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
	fmt.Println(result)

	return err

}

func UpdateAccountInDB(client *mongo.Client, ctx context.Context, oid primitive.ObjectID, update bson.M) *mongo.SingleResult {

	collection := client.Database(DBNAME).Collection("accounts")
	filter := bson.M{"_id": oid}
	result := collection.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(1))

	return result

}

func InsertProductToTheDB(ctx context.Context, client *mongo.Client, data ProductItem) (*mongo.InsertOneResult, error) {

	collection := client.Database(DBNAME).Collection(("products"))
	insertResult, err := collection.InsertOne(ctx, data)
	return insertResult, err
}

func GetProductFromDB(ctx context.Context, client *mongo.Client, oid primitive.ObjectID) *mongo.SingleResult {

	collection := client.Database(DBNAME).Collection("products")
	result := collection.FindOne(ctx, bson.M{"_id": oid})
	return result
}

func DeleteProductFromDB(ctx context.Context, client *mongo.Client, oid primitive.ObjectID) error {

	collection := client.Database(DBNAME).Collection("products")

	result, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
	fmt.Println(result)

	return err

}

func UpdateProductInDB(ctx context.Context, client *mongo.Client, oid primitive.ObjectID, update bson.M) *mongo.SingleResult {

	collection := client.Database(DBNAME).Collection("products")
	filter := bson.M{"_id": oid}
	result := collection.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(1))

	return result
}
