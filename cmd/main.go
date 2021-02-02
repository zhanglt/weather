package main

//0 */2 * * * date >> /home/qw/weather
import (
	"context"
	"log"
	"time"

	"github.com/qiniu/qmgo"
	"github.com/zhanglt/weather/internal/model"
	"go.uber.org/fx"
	"gopkg.in/mgo.v2/bson"
)

func bootInvoke(ctx context.Context, client *qmgo.QmgoClient, conf *model.Config, logger *log.Logger) {
	t1 := time.Now()
	count := 0
	for name, area := range conf.Area.Area {
		filter := bson.M{"areaid": area} //查询条件
		up, ok := model.UpdateWeather(ctx, client, conf, area, logger)
		if ok != nil {
			logger.Println("更新错误：", name, area, ok)
		} else {
			result, err := client.Upsert(ctx, filter, up)
			if err != nil {
				logger.Println("更新失败错误信息：", area, "|", err)
			}
			count++
			if conf.Writable.LogLevel == "DEBUG" {
				logger.Println("更新信息：", area, result.MatchedCount, ":", result.ModifiedCount, ":", result.UpsertedCount)
			}
		}

	}
	t2 := time.Now()
	logger.Println(time.Now().Format("2006/1/2 15:04:05"), "共同步：", count, "条数据,用时：", t2.Sub(t1))

}
func main() {
	fx.New(
		fx.Provide(model.ProvideConfig),
		fx.Provide(model.ProvideLog),
		fx.Provide(model.ProvideDbClient),
		fx.Invoke(bootInvoke),
	)
}
