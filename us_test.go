package tushare

import (
	"os"
	"strings"
	"testing"

	v1 "github.com/isfk/tushare/gen/api/v1"
	"github.com/joho/godotenv"
)

func TestUsBasic(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.UsBasicRequest, v1.UsBasicResponse](ts, &Request[v1.UsBasicRequest]{
		Params:  &v1.UsBasicRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiUsBasic,
		Fields:  v1.FieldsUsBasic,
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

func TestUsTradecal(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.UsTradecalRequest, v1.UsTradecalResponse](ts, &Request[v1.UsTradecalRequest]{
		Params:  &v1.UsTradecalRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiUsTradecal,
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

func TestUsDaily(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.UsDailyRequest, v1.UsDailyResponse](ts, &Request[v1.UsDailyRequest]{
		Params:  &v1.UsDailyRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiUsDaily,
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
