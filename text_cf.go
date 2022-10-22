package cf

// 文本、字符串类函数

import (
	"fmt"
	"strings"

	"github.com/jianka/cf/cast"
)

// 去除字符串最后一位字符
func SubStrLast(s string) string {
	bt := []rune(s)
	end := len(bt) - 1
	return string(bt[0:end])
}

// 去除字符串第一个字符
func SubStrFirst(s string) string {
	bt := []rune(s)
	end := len(bt)
	return string(bt[1:end])
}

// 如果最后一个字符串是"/"的话，去除字符串最后一位字符，如果不是返回原字符串
func SubStrLast2(s string) string {
	bt := []rune(s)
	if string(bt[len(bt)-1]) == "/" {
		end := len(bt) - 1
		return string(bt[0:end])
	}
	return s
}

// 超出范围的字符串显示...
func SubStrShow(s string, i int) string {
	arr := []rune(s)
	all := 0
	for _, v := range arr {
		if v > 127 {
			all += 2
		} else {
			all++
		}
		if all > i {
			break
		}
	}
	if all > i {
		n := i - 3
		all2 := 0
		var ar []rune
		for _, v := range arr {
			// 截取字符串
			if v > 127 {
				all2 += 2
			} else {
				all2++
			}
			if all2 <= n {
				ar = append(ar, v)
			} else {
				break
			}
		}
		return string(ar) + "..."
	} else {
		return s
	}
}

// 整型64切片转字符串
func SliceUint64ToString(s []uint64) string {
	str := strings.Replace(strings.Trim(fmt.Sprint(s), "[]"), " ", ",", -1)
	return str
}

// 状态转换为显示文本
func StatusForSpanText(i interface{}, s ...[]string) string {
	slice := []string{
		`<span class="layui-badge layui-bg-cyan">Off</span>`,
		`<span class="layui-badge layui-bg-blue">On</span>`,
	}
	ret := ""
	n := cast.ToInt(i)
	if len(s) == 3 && len(s[0]) == len(s[1]) && len(s[1]) == len(s[2]) {
		for k, v := range s[0] {
			if cast.ToInt(v) == n {
				ret = `<span class="layui-badge" style="background-color:` + s[2][k] + `;">` + s[1][k] + `</span>`
				break
			}
		}
	} else {
		if n >= 0 && n < len(slice) {
			ret = slice[n]
		}
	}
	return ret
}

// 状态转换为显示文本 第二版本
func StatusForSpanText2(i interface{}, s ...[]string) string {
	slice := []string{
		`<span class="tag-item-new tag-item-danger">Off</span>`,
		`<span class="tag-item-new">On</span>`,
	}
	ret := ""
	n := cast.ToInt(i)
	if len(s) == 3 && len(s[0]) == len(s[1]) && len(s[1]) == len(s[2]) {
		for k, v := range s[0] {
			if cast.ToInt(v) == n {
				ret = `<span class="tag-item-new" style="background-color:` + s[2][k] + `;">` + s[1][k] + `</span>`
				break
			}
		}
	} else if len(s) == 1 {
		if n >= 0 && n < len(slice) && len(s[0]) == 2 {
			slice2 := []string{
				`<span class="tag-item-new tag-item-danger">` + s[0][0] + `</span>`,
				`<span class="tag-item-new">` + s[0][1] + `</span>`,
			}
			ret = slice2[n]
		}
	} else {
		if n >= 0 && n < len(slice) {
			ret = slice[n]
		}
	}
	return ret
}

// 检测字符串是否为数字
func IsNumeric(s string) bool {
	n := CountWords(s, 1, "numberCount")
	return len(s) == int(n)
}
