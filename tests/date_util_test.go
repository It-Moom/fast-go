/*
 * @PackageName: tests
 * @Description: 时间日期工具测试
 * @Author: Casso-Wong
 * @Date: 2021-10-04 19:39:17
 * @Last Modified by: Casso-Wong
 * @Last Modified time: 2021-10-04 21:08:56
 */
package tests

import (
	"fast-go/pkg/utils/date_util"
	"testing"
	"time"
)

func TestLocalNow(t *testing.T) {
	res := date_util.FulltimeNow()
	if res == "" {
		t.Error(res)
	}
	t.Log(res)
}

func TestFulltimeSlantNow(t *testing.T) {
	res := date_util.FulltimeSlantNow()
	if res == "" {
		t.Error(res)
	}
	t.Log(res)
}

func TestYmdNow(t *testing.T) {
	res := date_util.YmdNow()
	if res == "" {
		t.Error(res)
	}
	t.Log(res)
}

func TestHmsNow(t *testing.T) {
	res := date_util.HmsNow()
	if res == "" {
		t.Error(res)
	}
	t.Log(res)
}

func TestStartToday(t *testing.T) {
	res := date_util.StartToday()
	t.Log(res)
}

func TestParseTime(t *testing.T) {
	source := "2021-10-04 13:50:30"
	res, err := date_util.ParseTime(source)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestParseTimeStamp(t *testing.T) {
	source := "2021-10-04 13:50:30"
	res, err := date_util.ParseTimeStamp(source)
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}

func TestGetDurationByString(t *testing.T) {
	start := "2021-10-04 13:50:30"
	end := "2021-10-04 13:50:35"
	target := 5
	res, err := date_util.GetDurationByString(start, end)
	if err != nil || res != int64(target) {
		t.Error(err)
	}
	t.Log(res)
}

func TestGetDurationByTime(t *testing.T) {
	t1 := time.Now()
	t2 := t1.Add(time.Second * 5)
	target := 5
	res, err := date_util.GetDurationByTime(t1, t2)
	if err != nil || res != int64(target) {
		t.Error(err)
	}
	t.Log(res)
}
