package model

type Weather_f struct {
	Areaid  int     `json:"areaid"`
	Type    string  `json:"type"`
	Content Content `json:"content"`
}
type Content struct {
	C C `json:"c"`
	F F `json:"f"`
}
type F struct {
	F1 []F1   `json:"f1"`
	F0 string `json:"f0"` //预报发布时间 201203061100
}
type C struct {
	C1  string  `json:"c1"`  //区域 ID 101010100
	C2  string  `json:"c2"`  //城市英文名 beijing
	C3  string  `json:"c3"`  //城市中文名 北京
	C4  string  `json:"c4"`  //城市所在市英文名 beijing
	C5  string  `json:"c5"`  //城市所在市中文名 北京
	C6  string  `json:"c6"`  //城市所在省英文名 beijing
	C7  string  `json:"c7"`  //城市所在省中文名 北京
	C8  string  `json:"c8"`  //城市所在国家英文名 china
	C9  string  `json:"c9"`  //城市所在国家中文名 中国
	C10 string  `json:"c10"` //城市级别 1
	C11 string  `json:"c11"` //城市区号 010
	C12 string  `json:"c12"` //邮编 100000
	C13 float64 `json:"c13"` //经度 116.391
	C14 float64 `json:"c14"` //纬度 39.904
	C15 string  `json:"c15"` //海拔 33
	C16 string  `json:"c16"` //雷达站号 AZ9010
	C17 string  `json:"c17"` //时区 +8
}
type F1 struct {
	Fa string `json:"fa"` //白天天气现象编号 01
	Fb string `json:"fb"` //晚上天气现象编号 01
	Fc string `json:"fc"` //白天天气温度(摄氏度) 11
	Fd string `json:"fd"` //晚上天气温度(摄氏度) 0
	Fe string `json:"fe"` //白天风向编号 4
	Ff string `json:"ff"` //晚上风向编号 4
	Fg string `json:"fg"` //白天风力编号 1
	Fh string `json:"fh"` //晚上风力编号 0
	Fi string `json:"fi"` //日出日落时间(中间用|分割)
}
type Weather struct {
	Data   Data   `json:"data"`
	Status int    `json:"status"`
	Desc   string `json:"desc"`
}
type Data struct {
	Yesterday Yesterday  `json:"yesterday"`
	City      string     `json:"city"`
	Forecast  []Forecast `json:"forecast"`
	Ganmao    string     `json:"ganmao"`
	Wendu     string     `json:"wendu"`
}
type Yesterday struct {
	Date string `json:"date"`
	High string `json:"high"`
	Fx   string `json:"fx"`
	Low  string `json:"low"`
	Fl   string `json:"fl"`
	Type string `json:"type"`
}
type Forecast struct {
	Date      string `json:"date"`
	High      string `json:"high"`
	Fengli    string `json:"fengli"`
	Low       string `json:"low"`
	Fengxiang string `json:"fengxiang"`
	Type      string `json:"type"`
}
