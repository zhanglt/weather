package main

import (
	"fmt"
	"time"

	"github.com/zhanglt/weather/internal/model"
	"gopkg.in/mgo.v2/bson"
)

var logger = model.GetLogger()

func main() {
	client, ctx := model.GetDbClient()
	logger.Println("-----开始同步：", time.Now(), "-----")
	for k, v := range model.Areaid {
		filter := bson.M{"areaid": v} //查询条件
		up := model.UpdateWeather(v, client, ctx)

		result, err := client.Upsert(ctx, filter, up)
		if err != nil {
			fmt.Println("err:", err)
			logger.Println("更新失败错误信息：", v, "|", err)
		}
		logger.Println("更新信息：", k, result.MatchedCount, ":", result.ModifiedCount, ":", result.UpsertedCount)

	}
	logger.Println("=====结束同步：", time.Now(), "=====")
}
