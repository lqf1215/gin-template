package pkg

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

const StatTimeSuffix = "00:30:00" // set next schedule running timestamp (set 00:30:00 run schedule)

type ValuedBlockRange struct {
	Code  int //Code = 0-收益未开始，1-开始生效 2-有收益，3-有收益且今日到期 4-收益结算今日结束 9-收益结算完毕 -1 err
	Start string
	End   string
}

func TimeNowUnixStr() string {
	return strconv.FormatInt(TimeNow().Unix(), 10)
}
func TimeNowFormatString() string {
	return ConvertTimestampToTimeStr(TimeNowUnixStr(), "2006-01-02 15:04:05")
}

// 获取中国时区当前时间
func TimeNow() time.Time {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	return time.Now().In(cstSh)
}

/*
获取 日期两端 时间戳
start 2021-12-30 00:00:01
end   2021-12-30 23:59:59
*/
func TimeDayUnix(time2 time.Time) (start time.Time, end time.Time) {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	start, _ = time.ParseInLocation("2006-01-02 15:04:05", time2.Format("2006-01-02 ")+"00:00:00", cstSh)
	end, _ = time.ParseInLocation("2006-01-02 15:04:05", time2.Format("2006-01-02 ")+"23:59:59", cstSh)
	return
}

/*
"2021-09-03 17:11:00" -> 转 时间搓
*/
func TimeStrToUnix(timeStr string) int64 {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	t2, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, cstSh)
	if err != nil {
		return -1
	}
	return t2.Unix()
}

func TimeStrToUnixStr(timeStr string) string {
	timeUnix := TimeStrToUnix(timeStr)
	if timeUnix == -1 {
		return ""
	}
	return strconv.FormatInt(timeUnix, 10)
}

/*
	 具有特定格式（yyyy-MM-dd）的字符串，转换为Time对象
		timeStr yyyy-MM-dd
		suffix hh:mm:ss
		location 如Asia/Shanghai
*/
func TimeStringToTime(timeStr string, suffix string, location string) time.Time {
	if location == "" {
		location = "Asia/Shanghai"
	}
	var cstSh, _ = time.LoadLocation(location) //上海
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" "+suffix, cstSh)
	return t
}

/*
*

	根据所传时间字符串（格式：yyyy-MM-dd）计算与当前统计时间之间的间隔天数
*/
func GetDuratinDaysFromCurrentStatTime(timeStr string) (int, error) {
	statTimestamp := TimeStringToTime(timeStr, "00:00:00", "").Unix()
	curTimeStr := TimeNow().AddDate(0, 0, -1).Format("2006-01-02")
	curTimestamp := TimeStringToTime(curTimeStr, "00:00:00", "").Unix()
	if statTimestamp > curTimestamp {
		return 0, errors.New("the statTimeStr set beyond the current statistic time:" + timeStr + ">" + curTimeStr)
	}
	statTime, _ := time.Parse("2006-01-02", timeStr)
	curTime, _ := time.Parse("2006-01-02", curTimeStr)
	durationOfDays := curTime.Sub(statTime).Hours() / 24
	return int(durationOfDays), nil
}

func TimeStampBeautify(timestamp string) string {
	if timestamp == "" {
		return ""
	}
	return ConvertTimestampToTimeStr(timestamp, "2006-01-02 15:04:05")
}
func ConvertTimestampToTimeStr(timestamp string, format string) string {
	timeString, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return ""
	}
	return time.Unix(timeString, 0).Format(format)
}

/*
粗略计算两天间隔，不考虑具体小时，如2021-09-01 23:59:59 和2021-09-02 00:00:00 相差一天
*/
func GetDurationDaysForTimestamp(foreTsStr string, lateTsStr string) string {
	foreTs, err := strconv.ParseInt(foreTsStr, 10, 64)
	if err != nil {
		return ""
	}
	lateTs, err := strconv.ParseInt(lateTsStr, 10, 64)
	if err != nil {
		return ""
	}
	foreTimeTs := TimeStringToTime(time.Unix(foreTs, 0).Format("2006-01-02"), "00:00:00", "")
	lateTimeTs := TimeStringToTime(time.Unix(lateTs, 0).Format("2006-01-02"), "00:00:00", "")
	durationDays := int64(lateTimeTs.Sub(foreTimeTs).Hours() / 24)
	return strconv.FormatInt(durationDays, 10)
}

func GetWeekMonday() time.Time {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	return weekStart
}

func TimeAgo(t time.Time) string {
	now := time.Now()
	diff := now.Sub(t)

	seconds := int(diff.Seconds())
	minutes := int(diff.Minutes())
	hours := int(diff.Hours())
	days := int(diff.Hours() / 24)
	months := int(diff.Hours() / 720)

	switch {
	case months > 0:
		return fmt.Sprintf("%d 个月前", months)
	case days > 0:
		return fmt.Sprintf("%d 天前", days)
	case hours > 0:
		return fmt.Sprintf("%d 小时前", hours)
	case minutes > 0:
		return fmt.Sprintf("%d 分钟前", minutes)
	default:
		return fmt.Sprintf("%d 秒前", seconds)
	}
}
