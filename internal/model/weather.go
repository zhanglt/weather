package model

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/qiniu/qmgo"
	"gopkg.in/mgo.v2/bson"
)

var logger = GetLogger()
var client, ctx = GetDbClient()

func getWeatherNew(area int) Weather {

	var wn Weather
	res, err := http.Get(WeatherURL + strconv.Itoa(area))
	if err != nil {
		fmt.Println("获取天气信息错误，区域编码：", area, "错误信息：", err)
		logger.Println("获取天气信息错误，区域编码：", area, "错误信息：", err)
		return wn
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("读取天气信息错误，区域编码：", area, "错误信息：", err)
		logger.Println("读取天气信息错误，区域编码：", area, "错误信息：", err)
	}
	json.Unmarshal(body, &wn)
	return wn
}

func getWeatherF(area int, client *qmgo.QmgoClient, ctx context.Context) Weather_f {
	var wf Weather_f

	filter := bson.M{"areaid": area}         //查询条件
	err := client.Find(ctx, filter).One(&wf) //从mongodb数据库中查询区域对应的天气信息并反序列化
	if err != nil {
		fmt.Println("mongoDB查询天气信息错误，区域编码：", area, "错误信息：", err)
		logger.Println("mongoDB查询天气信息错误，区域编码：", area, "错误信息：", err)
	}
	return wf
}

func UpdateWeather(area int, client *qmgo.QmgoClient, ctx context.Context) Weather_f {
	wn := getWeatherNew(area)
	wf := getWeatherF(area, client, ctx)
	wf.Date = time.Now()
	for i := 0; i < 3; i++ {
		//更新信息
		wf.Date = time.Now()
		//wf.Content.F.F0 = "202101271907"
		windPower := wn.Data.Forecast[i].Fengli
		windPower = windPower[9 : len(windPower)-1]
		windPower = windPower[:len(windPower)-5]
		tempertureH := wn.Data.Forecast[i].High
		tempertureH = tempertureH[6 : len(tempertureH)-1]
		tempertureH = tempertureH[:len(tempertureH)-2]
		tempertureL := wn.Data.Forecast[i].Low
		tempertureL = tempertureL[6 : len(tempertureL)-1]
		tempertureL = tempertureL[:len(tempertureL)-2]
		wf.Content.F.F1[i].Fa = ConvertWeather(wn.Data.Forecast[i].Type)
		wf.Content.F.F1[i].Fb = ConvertWeather(wn.Data.Forecast[i].Type)
		wf.Content.F.F1[i].Fc = tempertureH
		wf.Content.F.F1[i].Fd = tempertureL
		wf.Content.F.F1[i].Fe = ConvertWind(wn.Data.Forecast[i].Fengxiang)
		wf.Content.F.F1[i].Ff = ConvertWind(wn.Data.Forecast[i].Fengxiang)
		wf.Content.F.F1[i].Fg = windPower
		wf.Content.F.F1[i].Fh = windPower
	}
	return wf
}
