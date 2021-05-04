package https

import (
	json2 "encoding/json"
	"testing"
)

func TestGetStockBasic(t *testing.T) {
	Get()
}

func TestJsonDeser(t *testing.T) {
	json := "{\"request_id\":\"e5e578d2acea11eb80eb1d82661c7a871620141044427830\",\"code\":11,\"msg\":\"test\",\"has_more\":false,\"data\":{\"fields\":[\"ts_code\",\"symbol\",\"name\",\"area\",\"industry\",\"enname\",\"cnspell\",\"market\",\"curr_type\",\"list_status\",\"list_date\",\"delist_date\",\"is_hs\"],\"items\":[[\"000001.SZ\",\"000001\",\"平安银行\",\"深圳\",\"银行\",\"Ping An Bank Co., Ltd.\",\"payh\",\"主板\",\"CNY\",\"L\",\"19910403\",null,\"S\"]]}}"
	res := StockBasicResponse{}
	json2.Unmarshal([]byte(json), &res)
}
