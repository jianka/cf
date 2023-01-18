package cf

// 时间函数

import (
	"time"

	"github.com/jianka/cf/cast"
)

// 时间戳转日期
// t 时间戳 秒
// types 格式 默认 1
// 1 2006-01-02 15:04:05
// 2 2006/01/02 15:04:05
// 3 2006年01月02日 15点04分05秒
// 4 2006-01-02
// 5 2006/01/02
// 6 2006年01月02日
func TransTime(t interface{}, types int) string {
	ret := ""
	timeobj := time.Unix(cast.ToInt64(t), 0)
	switch types {
	case 2:
		ret = timeobj.Format("2006/01/02 15:04:05")
	case 3:
		ret = timeobj.Format("2006年01月02日 15点04分05秒")
	case 4:
		ret = timeobj.Format("2006-01-02")
	case 5:
		ret = timeobj.Format("2006/01/02")
	case 6:
		ret = timeobj.Format("2006年01月02日")
	case 7:
		ret = timeobj.Format("2006-01-02-15-04-05")
	default:
		ret = timeobj.Format("2006-01-02 15:04:05")
	}
	return ret
}

// 格式化时间
// t 时间类型
// types 格式 默认 1
// 1 2006-01-02 15:04:05
// 2 2006/01/02 15:04:05
// 3 2006年01月02日 15点04分05秒
// 4 2006-01-02
// 5 2006/01/02
// 6 2006年01月02日
func TimeFormat(t time.Time, types int) string {
	layout := "2006-01-02 15:04:05"
	switch types {
	case 2:
		layout = "2006/01/02 15:04:05"
	case 3:
		layout = "2006年01月02日 15点04分05秒"
	case 4:
		layout = "2006-01-02"
	case 5:
		layout = "2006/01/02"
	case 6:
		layout = "2006年01月02日"
	case 7:
		layout = "2006-01-02-15-04-05"
	}
	return t.Format(layout)
}

func TransTimetamp(t string) uint32 {
	times, _ := time.Parse("2006-01-02 15:04:05", t)
	return uint32(times.Unix())
}

// 取今天日期的字符串格式
func GetDayString(t ...string) string {
	template := "20060102"
	if len(t) == 1 {
		template = t[0]
	}
	return time.Now().Format(template)
}

// 返回加秒的时间
func AddSecondDate(s int) time.Time {
	timeobj := time.Now().Add(time.Second * time.Duration(s))
	return timeobj
}

// 返回减秒的时间
func SubSecondDate(s int) time.Time {
	timeobj := time.Now().Add(-(time.Second * time.Duration(s)))
	return timeobj
}
