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
