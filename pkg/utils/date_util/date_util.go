/*
 * @PackageName: date_util
 * @Description: 时间日期工具
 * @Author: Casso-Wong
 * @Date: 2021-10-04 19:31:04
 * @Last Modified by: Casso-Wong
 * @Last Modified time: 2021-10-04 20:39:45
 */
package date_util

import (
	"errors"
	"time"
)

// LocalNow 返回本地标准时间字符串
func LocalNow() string {
	t := time.Now().UTC().Local()
	return t.Format("2006-01-02 15:04:05")
}

// ParseTime 输入时间字符串，返回时间
func ParseTime(str string) (time.Time, error) {
	if str == "" {
		return time.Now(), errors.New("empty string")
	}

	return time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
}

// ParseTimeStamp 输入时间字符串，返回时间戳
func ParseTimeStamp(str string) (int64, error) {
	if str == "" {
		return 0, errors.New("empty string")
	}

	t, err := time.ParseInLocation("2006-01-02 15:04:05", str, time.Local)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// StartToday 返回本地当天零点
func StartToday() time.Time {
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day()-1, 0, 0, 0, 0, t.Location())
	tm2 := tm1.AddDate(0, 0, 1)
	return tm2
}

// GetDuration 传入时间字符串，获取相差时间秒数，stime:开始时间;etime：结束时间
func GetDurationByString(stime, etime string) (res int64, errs error) {
	if stime == "" || etime == "" {
		return 0, errors.New("empty string")
	}
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", stime, time.Local)
	t2, err2 := time.ParseInLocation("2006-01-02 15:04:05", etime, time.Local)

	// 开始时间必须在结束时间之前
	if err == nil && err2 == nil && t1.Before(t2) {
		res = t2.Unix() - t1.Unix()
		errs = nil
	} else {
		res = 0
		errs = err
	}
	return
}

// GetDuration 传入时间，获取相差时间秒数，stime:开始时间;etime：结束时间
func GetDurationByTime(stime, etime time.Time) (res int64, errs error) {
	// 开始时间必须在结束时间之前
	if stime.Before(etime) {
		res = etime.Unix() - stime.Unix()
		errs = nil
	} else {
		res = 0
		errs = errors.New("开始时间晚于结束时间")
	}
	return
}
