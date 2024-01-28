package tushare

import (
	"os"
	"strings"
	"testing"

	v1 "github.com/isfk/tushare/gen/api/v1"
	"github.com/joho/godotenv"
)

func TestSgeBasic(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.SgeBasicRequest, v1.SgeBasicResponse](ts, &Request[v1.SgeBasicRequest]{
		Params:  &v1.SgeBasicRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiSgeBasic,
		Fields:  v1.FieldsSgeBasic,
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

func TestSgeDaily(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.SgeDailyRequest, v1.SgeDailyResponse](ts, &Request[v1.SgeDailyRequest]{
		Params:  &v1.SgeDailyRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiSgeDaily,
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
