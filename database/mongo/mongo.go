package mongo

import (
	"api/config/setting"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoHandler struct {
	client *mongo.Client
}

var ConnectMongo MongoHandler

func  OpenConnect(){
	credential := options.Credential{
		Username: setting.MongoConfig.Username,
		Password: setting.MongoConfig.Password,
	}

	clientOptions := options.Client().ApplyURI("mongodb://" + setting.MongoConfig.Host).SetAuth(credential)
	clientOptions.SetMaxPoolSize(5000) //最大連線數
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	ConnectMongo.client = client
}


func Collection(collectionName string) *mongo.Collection {
	return ConnectMongo.client.Database(setting.MongoConfig.DBName).Collection(collectionName)
}


