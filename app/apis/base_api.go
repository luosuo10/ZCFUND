package apis

import (
	"net/http"
	"io/ioutil"
	"zc_box/app/entities"
	"fmt"
)

var (
	// 基金列表API
	FUND_LIST_API = "http://fund.eastmoney.com/js/fundcode_search.js"
)

func getFundListAPIData() (data string, err error) {
	var resp *http.Response
	if resp, err = http.Get(FUND_LIST_API); err != nil {
		return "", err
	}
	var respBody []byte
	if respBody, err = ioutil.ReadAll(resp.Body); err != nil {
		return "", err
	}
	data = string(respBody)
	data = data[9 : len(data) - 2]
	return data, nil
}

func GetFundList(url string) (FundBaseInfos []entities.FundBaseInfo, err error) {
	var data string
	if data, err = getFundListAPIData(); err != nil {
		return nil, err
	}
	fmt.Printf("data = %v\n", data)
	return nil, nil
}