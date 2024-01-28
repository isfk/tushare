package tushare

import (
	"os"
	"strings"
	"testing"
	"time"

	v1 "github.com/isfk/tushare/gen/api/v1"
	"github.com/joho/godotenv"
)

func TestFutBasic(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.FutBasicRequest, v1.FutBasicResponse](ts, &Request[v1.FutBasicRequest]{
		Params:  &v1.FutBasicRequest{Limit: "2", Offset: "0", Exchange: "CFFEX"},
		ApiName: v1.ApiFutBasic,
		Fields:  v1.FieldsFutBasic,
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

func TestFutDaily(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.FutDailyRequest, v1.FutDailyResponse](ts, &Request[v1.FutDailyRequest]{
		Params:  &v1.FutDailyRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiFutDaily,
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

func TestFtMins(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.FtMinsRequest, v1.FtMinsResponse](ts, &Request[v1.FtMinsRequest]{
		Params:  &v1.FtMinsRequest{Limit: "2", Offset: "0", TsCode: "CU2310.SHF", Freq: "1min"},
		ApiName: v1.ApiFtMins,
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

func TestFutWsr(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.FutWsrRequest, v1.FutWsrResponse](ts, &Request[v1.FutWsrRequest]{
		Params:  &v1.FutWsrRequest{Limit: "2", Offset: "0", Symbol: "ZN"},
		ApiName: v1.ApiFutWsr,
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

func TestFutSettle(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.FutSettleRequest, v1.FutSettleResponse](ts, &Request[v1.FutSettleRequest]{
		Params:  &v1.FutSettleRequest{Limit: "2", Offset: "0", TradeDate: TushareDate(time.Now().AddDate(0, 0, -1)), Exchange: "SHFE"},
		ApiName: v1.ApiFutSettle,
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

func TestFutHolding(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.FutHoldingRequest, v1.FutHoldingResponse](ts, &Request[v1.FutHoldingRequest]{
		Params:  &v1.FutHoldingRequest{Limit: "2", Offset: "0", Symbol: "ZN"},
		ApiName: v1.ApiFutHolding,
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

func TestFutMapping(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.FutMappingRequest, v1.FutMappingResponse](ts, &Request[v1.FutMappingRequest]{
		Params:  &v1.FutMappingRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiFutMapping,
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

func TestFutWeeklyDetail(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.FutWeeklyDetailRequest, v1.FutWeeklyDetailResponse](ts, &Request[v1.FutWeeklyDetailRequest]{
		Params:  &v1.FutWeeklyDetailRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiFutWeeklyDetail,
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
