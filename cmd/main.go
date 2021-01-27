package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/zhanglt/internal/model"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var areaid = []string{"01080901", "01080903", "01080904", "01080906", "01080907", "01080908", "01080909", "01080910", "01080911", "01080912", "01080913", "01080914", "01080915", "01080916", "01080917", "01080801", "01080802", "01080803", "01080804", "01080805", "01080806", "01080807", "01080808", "01080809", "01080810", "01080601", "01080603", "01080604", "01080605", "01080606", "01080607", "01080608", "01080610", "01080611", "01080612", "01080613", "01080614", "01080615", "01080501", "01080502", "01080503", "01080504", "01080505", "01080506", "01080507", "01080508", "01080509", "01080510", "01080511", "01080512", "01080401", "01080402", "01080403", "01080404", "01080406", "01080407", "01080408", "01080409", "01080410", "01080411", "01080412", "01080201", "01080202", "01080203", "01080204", "01080205", "01080206", "01080207", "01080101", "01080102", "01080103", "01080104", "01080105", "01080106", "01080107", "01081201", "01081202", "01081203", "01081204", "01081205", "01081206", "01081207", "01081208", "01081209", "01081210", "01081211", "01081212", "01081101", "01081102", "01081103", "01081104", "01081105", "01081106", "01081107", "01081109", "01081001", "01081002", "01081003", "01081004", "01081005", "01081006", "01081007", "01081008", "01081009", "01081010", "01081011", "01081012", "01081014", "01081015", "01081016", "01080301", "01080701", "01080702", "01080703", "01080704", "01080705", "01080706", "01080707", "01080708", "01080709", "01080710", "01080711", "01080712", "01080101", "01080102", "01080103", "01080104", "01080105", "01080106", "01080107"}

func main() {
	var area = "101080511"
	var f interface{}
	var m model.Weather
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	collection := client.Database("pub").Collection("weather")

	res, err := http.Get("http://wthrcdn.etouch.cn/weather_mini?citykey=" + area)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(body, &f)
	json.Unmarshal(body, &m)

	r, err := collection.InsertOne(ctx, m)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.InsertedID)

}
