package tushare

import (
	"os"
	"strings"
	"testing"

	v1 "github.com/isfk/tushare/gen/api/v1"
	"github.com/joho/godotenv"
)

func TestHkBasic(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.HkBasicRequest, v1.HkBasicResponse](ts, &Request[v1.HkBasicRequest]{
		Params:  &v1.HkBasicRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiHkBasic,
		Fields:  v1.FieldsHkBasic,
	})
	if err != nil {
		if strings.Contains(err.Error(), "code=40203") {
			t.Logf("权限问题: %v", err.Error())
			return
		}
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestHkTradecal(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.HkTradecalRequest, v1.HkTradecalResponse](ts, &Request[v1.HkTradecalRequest]{
		Params:  &v1.HkTradecalRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiHkTradecal,
		Fields:  []string{},
	})
	if err != nil {
		if strings.Contains(err.Error(), "code=40203") {
			t.Logf("权限问题: %v", err.Error())
			return
		}
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestHkDaily(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.HkDailyRequest, v1.HkDailyResponse](ts, &Request[v1.HkDailyRequest]{
		Params:  &v1.HkDailyRequest{Limit: "2", Offset: "0", TsCode: "00001.HK"},
		ApiName: v1.ApiHkDaily,
		Fields:  []string{},
	})
	if err != nil {
		if strings.Contains(err.Error(), "code=40203") {
			t.Logf("权限问题: %v", err.Error())
			return
		}
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}

func TestHkMins(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.HkMinsRequest, v1.HkMinsResponse](ts, &Request[v1.HkMinsRequest]{
		Params:  &v1.HkMinsRequest{Limit: "2", Offset: "0", TsCode: "00001.HK"},
		ApiName: v1.ApiHkMins,
		Fields:  []string{},
	})
	if err != nil {
		if strings.Contains(err.Error(), "code=40203") {
			t.Logf("权限问题: %v", err.Error())
			return
		}
		t.Errorf("%v", err)
		return
	}
	if resp.Resp != nil {
		t.Logf("%v", resp.Resp.String())
	}
}
