package tushare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const baseUrl = "http://api.tushare.pro"

func newTransport() *http.Client {
	return &http.Client{}
}

type RequestData struct {
	ApiName string      `json:"api_name"`
	Token   string      `json:"token"`
	Params  interface{} `json:"params"`
	Fields  []string    `json:"fields"`
}

type Tushare struct {
	baseUrl string
	httpCli *http.Client
	token   string
}

func NewTushare(token string) *Tushare {
	return &Tushare{
		baseUrl: baseUrl,
		httpCli: newTransport(),
		token:   token,
	}
}

type Request[T any] struct {
	Params  *T
	ApiName string
	Fields  []string
}
type Response[T any] struct {
	Resp *T
}

type TushareResponse struct {
	Code int `json:"code,omitempty"`
	Data struct {
		Fields []string `json:"fields,omitempty"`
		Items  [][]any  `json:"items,omitempty"`
	} `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func RequestTushare[Q any, R any](ts *Tushare, req *Request[Q]) (*Response[R], error) {
	ret := &Response[R]{}
	_, err := ts.requestTushare(req.ApiName, req.Params, req.Fields, &ret.Resp)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *Tushare) requestTushare(apiName string, params interface{}, fields []string, v interface{}) (*TushareResponse, error) {
	requestBytes, err := json.Marshal(&RequestData{
		ApiName: apiName,
		Token:   c.token,
		Params:  params,
		Fields:  fields,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.baseUrl, bytes.NewBuffer(requestBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, err := c.httpCli.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code: %d", res.StatusCode)
	}

	resp := &TushareResponse{}
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, NewTushareError(resp.Code, resp.Message)
	}

	if resp.Data.Fields == nil || resp.Data.Items == nil {
		return nil, nil
	}

	mapList := []*map[string]any{}
	response := map[string][]*map[string]any{}
	for _, item := range resp.Data.Items {
		tmpMap := map[string]any{}
		for n, v := range resp.Data.Fields {
			switch item[n].(type) {
			case string:
				tmpMap[v] = item[n].(string)
			case float64:
				tmpMap[v] = item[n].(float64)
			case int64:
				tmpMap[v] = item[n].(int64)
			}
		}
		mapList = append(mapList, &tmpMap)
	}

	response["list"] = mapList
	responseBytes, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	if v != nil {
		err = json.Unmarshal(responseBytes, &v)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
}
