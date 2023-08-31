package customerconfig

import (
	customerconstants "Netxd_Customer/Customer_Connection/Customer_DAL_constants"
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



var Collection *mongo.Collection

func ConnectDatabase()(*mongo.Client,error) {
	clientoption := options.Client().ApplyURI(customerconstants.Connectstring)

	client, err := mongo.Connect(context.TODO(), clientoption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDb sucessfully connected")
	Collection = client.Database(customerconstants.Dbname).Collection(customerconstants.Dbcollection)

	fmt.Println("Collection Connected")
	return client,nil
}