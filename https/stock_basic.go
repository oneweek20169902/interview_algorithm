package https

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type StockBasicResponse struct {
	RequestId string `json:"request_id"`
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Data      struct {
		Fields  []string        `json:"fields"`
		Items   [][]interface{} `json:"items"`
		HasMore bool            `json:"has_more"`
	} `json:"data"`
	HasMore bool `json:"has_more"`
}

func Get() error {
	client := resty.New()
	resp, err := client.R().
		SetHeader(hdrContentTypeKey, "application/json; charset=utf-8").
		SetBody([]byte(`{"api_name":"stock_basic","token":"8502e804d14490734ae693b79ad64ecf5e4178ff8d6be01426bee185"}`)).
		SetResult(&StockBasicResponse{}).
		Post(URL)

	if err != nil {
		return err
	}
	fmt.Print(resp.String())
	var res = resp.Result().(*StockBasicResponse)
	fmt.Println(res)
	for k, v := range res.Data.Items {
		fmt.Println(k, v)
	}
	return nil
}
