/*
 * @PackageName: date_util
 * @Description: 时间日期工具
 * @Author: Casso-Wong
 * @Date: 2021-10-04 19:31:04
 * @Last Modified by: Casso-Wong
 * @Last Modified time: 2021-10-04 21:08:40
 */
package date_util

import (
	"errors"
	"time"
)

var (
	tem1 = "2006-01-02 15:04:05" // 默认格式
	tem2 = "2006/01/02 15:04:05" // 斜杠格式
	tem3 = "2006/01/02"          // 仅有年月日格式
	tem4 = "15:04:05"            // 仅有时分秒格式
)

// FulltimeNow 返回本地标准时间字符串--横杠
func FulltimeNow() string {
	t := time.Now().UTC().Local()
	return t.Format(tem1)
}

// FulltimeSlantNow 返回本地标准时间字符串--斜杠
func FulltimeSlantNow() string {
	t := time.Now().UTC().Local()
	return t.Format(tem2)
}

// YmdNow 返回本地标准时间字符串--年月日
func YmdNow() string {
	t := time.Now().UTC().Local()
	return t.Format(tem3)
}

// HmsNow 返回本地标准时间字符串--时分秒
func HmsNow() string {
	t := time.Now().UTC().Local()
	return t.Format(tem4)
}

// ParseTime 输入时间字符串，返回时间
func ParseTime(str string) (time.Time, error) {
	if str == "" {
		return time.Now(), errors.New("empty string")
	}
	return time.ParseInLocation(tem1, str, time.Local)
}

// ParseTimeStamp 输入时间字符串，返回时间戳
func ParseTimeStamp(str string) (int64, error) {
	if str == "" {
		return 0, errors.New("empty string")
	}
	// 时间字符串解析
	t, err := time.ParseInLocation(tem1, str, time.Local)
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
func GetDurationByString(start, end string) (int64, error) {
	if start == "" || end == "" {
		return 0, errors.New("empty string")
	}
	// 时间字符串解析
	startTime, errStart := time.ParseInLocation(tem1, start, time.Local)
	endTime, errEnd := time.ParseInLocation(tem1, end, time.Local)
	if errStart != nil {
		return 0, errStart
	}
	if errEnd != nil {
		return 0, errEnd
	}

	// 开始时间必须在结束时间之前
	if startTime.Before(endTime) {
		return endTime.Unix() - startTime.Unix(), nil
	}
	return 0, errors.New("开始时间晚于结束时间")
}

// GetDuration 传入时间，获取相差时间秒数，stime:开始时间;etime：结束时间
func GetDurationByTime(start, end time.Time) (int64, error) {
	// 开始时间必须在结束时间之前
	if start.Before(end) {
		return end.Unix() - start.Unix(), nil
	}
	return 0, errors.New("开始时间晚于结束时间")
}
