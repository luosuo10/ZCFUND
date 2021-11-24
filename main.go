package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"zc_box/app/entities"
)
//https://open.feishu.cn/open-apis/bot/v2/hook/9504d0ff-50b3-424c-b182-3e6fb03e449a

var(
	API = "https://fundgz.1234567.com.cn/js/%d.js"
)


func HttpGet2(url string) (data []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode == 200 {
		return ioutil.ReadAll(response.Body)
	}
	return nil, errors.New("http code != 200")
}

func getFundList() string {
	resp, err := http.Get("http://fund.eastmoney.com/js/fundcode_search.js")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("status = %v\n", resp.Status)
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	body := string(data)
	//fmt.Printf("body = %v\n", body)
	return body
}

func main() {
	//code := "012414"
	//url.Parse()
	//url := fmt.Sprint("https://fundgz.1234567.com.cn/js/012414.js")
	//fmt.Printf("url = %s\n", url)
	//
	//body, err := HttpGet2(url)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(string(body))
	//var FundResult FundItem
	//err = json.Unmarshal(body[8:len(body)-2], &FundResult)
	//fmt.Printf("err = %v\n", err)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%+v", FundResult)
	fundList := getFundList()
	fundList = fundList[9:len(fundList)-2]
	//fmt.Printf("%v\n", fundList)
	FundItems := make([]entities.FundBaseInfo, len(fundList))
	arr := strings.Split(fundList, "],[")
	for idx, item := range arr {
		if idx == 0 {
			item = strings.TrimPrefix(item, "= [[")
		}
		if idx == len(arr) - 1 {
			item = strings.TrimSuffix(item, "]]")
		}
		items := strings.Split(item, ",")
		//fmt.Printf("%v\n", items)
		//["000001" "HXCZHH" "华夏成长混合" "混合型-偏股" "HUAXIACHENGZHANGHUNHE"]
		FundItems[idx] = entities.FundBaseInfo{
			FundCode: items[0],
			Name: items[2],
			Type: items[3],
		}
		fmt.Printf("%+v\n", FundItems[idx])
	}
}