package tushare

import (
	"os"
	"strings"
	"testing"

	v1 "github.com/isfk/tushare/gen/api/v1"
	"github.com/joho/godotenv"
)

func TestIndexBasic(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.IndexBasicRequest, v1.IndexBasicResponse](ts, &Request[v1.IndexBasicRequest]{
		Params:  &v1.IndexBasicRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiIndexBasic,
		Fields:  v1.FieldsIndexBasic,
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

func TestIndexDaily(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.IndexDailyRequest, v1.IndexDailyResponse](ts, &Request[v1.IndexDailyRequest]{
		Params:  &v1.IndexDailyRequest{Limit: "2", Offset: "0", TsCode: "399300.SZ"},
		ApiName: v1.ApiIndexDaily,
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

func TestIndexWeekly(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.IndexWeeklyRequest, v1.IndexWeeklyResponse](ts, &Request[v1.IndexWeeklyRequest]{
		Params:  &v1.IndexWeeklyRequest{Limit: "2", Offset: "0", TsCode: "399300.SZ"},
		ApiName: v1.ApiIndexWeekly,
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

func TestIndexMonthly(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.IndexMonthlyRequest, v1.IndexMonthlyResponse](ts, &Request[v1.IndexMonthlyRequest]{
		Params:  &v1.IndexMonthlyRequest{Limit: "2", Offset: "0", TsCode: "399300.SZ"},
		ApiName: v1.ApiIndexMonthly,
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

func TestIndexWeight(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.IndexWeightRequest, v1.IndexWeightResponse](ts, &Request[v1.IndexWeightRequest]{
		Params:  &v1.IndexWeightRequest{Limit: "2", Offset: "0", IndexCode: "399300.SZ"},
		ApiName: v1.ApiIndexWeight,
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

func TestIndexDailyBasic(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.IndexDailyBasicRequest, v1.IndexDailyBasicResponse](ts, &Request[v1.IndexDailyBasicRequest]{
		Params:  &v1.IndexDailyBasicRequest{Limit: "2", Offset: "0", TsCode: "399300.SZ"},
		ApiName: v1.ApiIndexDailyBasic,
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

func TestIndexClassify(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.IndexClassifyRequest, v1.IndexClassifyResponse](ts, &Request[v1.IndexClassifyRequest]{
		Params:  &v1.IndexClassifyRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiIndexClassify,
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

func TestIndexMember(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.IndexMemberRequest, v1.IndexMemberResponse](ts, &Request[v1.IndexMemberRequest]{
		Params:  &v1.IndexMemberRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiIndexMember,
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

func TestDailyInfo(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.DailyInfoRequest, v1.DailyInfoResponse](ts, &Request[v1.DailyInfoRequest]{
		Params:  &v1.DailyInfoRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiDailyInfo,
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

func TestSzDailyInfo(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.SzDailyInfoRequest, v1.SzDailyInfoResponse](ts, &Request[v1.SzDailyInfoRequest]{
		Params:  &v1.SzDailyInfoRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiSzDailyInfo,
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

func TestThsIndex(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.ThsIndexRequest, v1.ThsIndexResponse](ts, &Request[v1.ThsIndexRequest]{
		Params:  &v1.ThsIndexRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiThsIndex,
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

func TestThsDaily(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.ThsDailyRequest, v1.ThsDailyResponse](ts, &Request[v1.ThsDailyRequest]{
		Params:  &v1.ThsDailyRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiThsDaily,
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

func TestThsMember(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.ThsMemberRequest, v1.ThsMemberResponse](ts, &Request[v1.ThsMemberRequest]{
		Params:  &v1.ThsMemberRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiThsMember,
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

func TestCiDaily(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.CiDailyRequest, v1.CiDailyResponse](ts, &Request[v1.CiDailyRequest]{
		Params:  &v1.CiDailyRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiCiDaily,
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

func TestIndexGlobal(t *testing.T) {
	_ = godotenv.Load()
	ts := NewTushare(os.Getenv("token"))
	resp, err := RequestTushare[v1.IndexGlobalRequest, v1.IndexGlobalResponse](ts, &Request[v1.IndexGlobalRequest]{
		Params:  &v1.IndexGlobalRequest{Limit: "2", Offset: "0"},
		ApiName: v1.ApiIndexGlobal,
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
