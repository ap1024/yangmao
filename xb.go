/*******************************************
QQ交流群：418353744
QQ线报群：263723430
********************************************/
package youhui

import (
	"encoding/json"
	"fmt"
	
	"github.com/beego/beego/v2/adapter/httplib"
	"github.com/cdle/sillyGirl/core"
)
var yangmao = core.NewBucket("yangmao")
type XB struct {
	Data struct {
		Type         string   `json:"type"`
		Title        string   `json:"title"`
		Time         string   `json:"Time"`
		Rule         string   `json:"rule"`
		Manner       string   `json:"manner"`
		Explanation  string   `json:"explanation"`
		Introduction string   `json:"Introduction"`
		Picture      []string `json:"Picture"`
	} `json:"data"`
	User struct {
		UserID int    `json:"userID"`
		Nick   string `json:"nick"`
	} `json:"user"`
}

func init() {
	core.AddCommand("", []core.Function{
		{ 	Rules: []string{"raw ^羊毛$"},
			Cron:  "30 8-18 * * *",
			Handle: func(s core.Sender) interface{} {
				msg:=getXb()
				if push, ok := core.GroupPushs["qq"]; ok {
					push(core.Int64(yangmao.Get("group")), nil, msg, "")
				}
				return msg
			},
		},
	})
}

func getXb() string {
	var rlt=""
	req := httplib.Get("http://xiaobai.klizi.cn/API/other/xb.php")
	data, _ := req.Bytes()
	fmt.Println(string(data))
	res := &XB{}
	json.Unmarshal([]byte(data), &res)
	rlt ="名称："+res.Data.Title+"\n"+
		"方法："+res.Data.Manner
	for _,pic:=range res.Data.Picture{
		rlt+="\n"+`[CQ:image,file=`+pic+`]`
	}		
	return rlt
}
