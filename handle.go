package strutils

import "strings"

// 字符串操作相关函数

// 截取指定长度的字符串 支持中文截取
// length 要截取的字符串长度, 此处使用uint16 0--65535长度足够了! 这里避免使用 int 或者 int64 可省去判断负数和少用内存 这样代码更精简更高效
func Substr(str string, length int) string {
	rstr := []rune(str)
	if length > len(rstr) {
		length = len(rstr)
	}
	return string(rstr[:length])
}

// 去除字符串中的空白字符包含 回车 换行 制表符等, 注意是字符串中的所有的空白符全部去除
func TrimWhiteSpace(s string) string {
	// 使用strings包中的Replacer对空白字符串进行批量 这里的规则都是成对的, 前面是查找字符串 后面是要替换的字符串
	replacer := strings.NewReplacer(" ", "", "\t", "", "\n", "", "\r", "", "\f", "")
	return replacer.Replace(s)
}

// 转换字符串为go语言中安全的命名样式, 即英文字母或者数字与 _ 组合,不能以数字开头
// SafeString会将所有的非字母或者数字全部转换为_  如果是数字开头则转换为_数字, 如123abc 转换为 _123abc
func SafeString(in string) string {
	if len(in) == 0 {
		return in
	}
	// 对输入的字符串的每个字符进行map映射操作处理
	data := strings.Map(func(r rune) rune {
		if isSafeRune(r) {
			return r
		}
		return '_' //将非安全命名字符全部替换为下划线 _
	}, in)
	// 判断第一个字符是否是数字
	firstStr := rune(data[0])
	if isNumber(firstStr) {
		return "_" + data
	}
	return data
}

// 判断是否是go语言中的安全命名字符  即字母,数字 或者下划线 _
func isSafeRune(r rune) bool {
	return isLetter(r) || isNumber(r) || r == '_'
}

func isLetter(r rune) bool {
	return 'A' <= r && r <= 'z'
}

func isNumber(r rune) bool {
	return '0' <= r && r <= '9'
}

// 字符串索引位置查找,找到返回对应的索引位置,  未找到返回 -1
func Index(slice []string, item string) int {
	for i := range slice {
		if slice[i] == item {
			return i
		}
	}
	return -1
}
