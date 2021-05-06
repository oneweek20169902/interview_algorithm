package https

import (
	"github.com/go-resty/resty/v2"
	"stock/models"
	"sync"
	"time"
)

type StockBasicResponse struct {
	RequestId string `json:"request_id"`
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Data      struct {
		Fields  []string   `json:"fields"`
		Items   [][]string `json:"items"`
		HasMore bool       `json:"has_more"`
	} `json:"data"`
	HasMore bool `json:"has_more"`
}

func updateStockBasic(stock *models.StockBasic, wg *sync.WaitGroup) error {
	defer wg.Done()
	id := stock.GetStockBasicByCode()
	if id > 0 {
		return stock.Update()
	} else {
		return stock.Insert()
	}
	return nil
}

func getStockBasicList() ([]models.StockBasic, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader(hdrContentTypeKey, "application/json; charset=utf-8").
		SetBody([]byte(`{"api_name":"stock_basic","token":"8502e804d14490734ae693b79ad64ecf5e4178ff8d6be01426bee185","fields":["ts_code","symbol","name","area","industry","fullname","list_date","enname","cnspell","curr_type","list_status","delist_date","is_hs","cnspell","market","exchange"]}`)).
		SetResult(&StockBasicResponse{}).
		Post(URL)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	var res = resp.Result().(*StockBasicResponse)
	stocks := make([]models.StockBasic, len(res.Data.Items), len(res.Data.Items))
	wg.Add(len(res.Data.Items))
	for i, v := range res.Data.Items {
		wg.Add(1)
		listDate, err := time.ParseInLocation("20060102", v[12], time.Local)
		if err != nil {
			listDate = time.Date(1991, 1, 1, 1, 1, 1, 1, time.Local)
		}
		deListDate, err := time.ParseInLocation("20060102", v[13], time.Local)
		if err != nil {

			deListDate = time.Date(1991, 1, 1, 1, 1, 1, 1, time.Local)
		}
		model := models.StockBasic{
			TsCode:     v[0],
			Code:       v[1],
			Name:       v[2],
			Area:       v[3],
			Industry:   v[4],
			FullName:   v[5],
			EnName:     v[6],
			CnSpell:    v[7],
			Market:     v[8],
			ExChange:   v[9],
			CurrType:   v[10],
			ListStatus: v[11],
			ListDate:   &listDate,
			DeListDate: &deListDate,
			IsHs:       v[14],
		}
		go updateStockBasic(&model, &wg)
		stocks[i] = model
	}
	wg.Wait()
	return stocks, nil
}
