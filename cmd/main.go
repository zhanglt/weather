package main

import (
	"time"

	"github.com/zhanglt/weather/internal/model"
	"gopkg.in/mgo.v2/bson"
)

var logger = model.GetLogger()

func main() {
	client, ctx := model.GetDbClient()
	logger.Println("-----开始同步：", time.Now(), "-----")
	t1 := time.Now()
	for name, area := range model.Areaid {
		filter := bson.M{"areaid": area} //查询条件
		up, ok := model.UpdateWeather(ctx, client, area)
		if ok != nil {
			logger.Println("更新错误：", name, area, ok)
		} else {
			result, err := client.Upsert(ctx, filter, up)
			if err != nil {
				logger.Println("更新失败错误信息：", area, "|", err)
			}
			logger.Println("更新信息：", area, result.MatchedCount, ":", result.ModifiedCount, ":", result.UpsertedCount)
		}

	}
	t2 := time.Now()
	logger.Println("=====结束同步：", time.Now(), "=====", t2.Sub(t1))
}
