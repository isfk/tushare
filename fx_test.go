package tushare

import (
	"os"
	"strings"
	"testing"

	v1 "github.com/isfk/tushare/gen/api/v1"
	"github.com/joho/godotenv"
)

func TestFxObasic(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.FxObasicRequest, v1.FxObasicResponse](ts, &Request[v1.FxObasicRequest]{
		Params:  &v1.FxObasicRequest{},
		ApiName: v1.ApiFxObasic,
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

func TestFxDaily(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.FxDailyRequest, v1.FxDailyResponse](ts, &Request[v1.FxDailyRequest]{
		Params:  &v1.FxDailyRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiFxDaily,
		Fields:  v1.FieldsFxDaily,
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
