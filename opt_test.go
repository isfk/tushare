package tushare

import (
	"os"
	"strings"
	"testing"

	v1 "github.com/isfk/tushare/gen/api/v1"
	"github.com/joho/godotenv"
)

func TestOptBasic(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.OptBasicRequest, v1.OptBasicResponse](ts, &Request[v1.OptBasicRequest]{
		Params:  &v1.OptBasicRequest{Limit: "2", Offset: "0", Exchange: "CFFEX"},
		ApiName: v1.ApiOptBasic,
		Fields:  v1.FieldsOptBasic,
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

func TestOptDaily(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.OptDailyRequest, v1.OptDailyResponse](ts, &Request[v1.OptDailyRequest]{
		Params:  &v1.OptDailyRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiOptDaily,
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
