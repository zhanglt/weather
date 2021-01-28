package model

import (
	"context"
	"fmt"

	"github.com/qiniu/qmgo"
)

var url = "mongodb://10.0.0.150:27017"
var database = "pub3"
var collection = "weather"

func GetDbClient() (client *qmgo.QmgoClient, ctx context.Context) {
	context := context.Background()
	client, err := qmgo.Open(context, &qmgo.Config{Uri: url, Database: database, Coll: collection})
	//cli2, err := qmgo.Open(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017", Database: "pub", Coll: "weather"})
	if err != nil {
		fmt.Println("err")
	}
	return client, context
}
