package spread

import (
	"encoding/base64"
	"fmt"
	"strings"
	"unicode"
)

func swap(runes []rune, start, end int, transform func(rune) rune) {
	if len(runes) > 1 {
		tmp := transform(runes[start])
		runes[start] = transform(runes[end])
		runes[end] = tmp
	}
}

// characterConversion 字符转换，处理单个字符：
// - 交换字母大小写（大写→小写，小写→大写）
// - 对数字进行循环移位（如 '5'+1 → '6'，'9'+1 → '0'）
// - 其他字符保持不变
func characterConversion(r rune, offset int) rune {
	offset %= 10
	if unicode.IsUpper(r) {
		return unicode.ToLower(r)
	} else if unicode.IsLower(r) {
		return unicode.ToUpper(r)
	} else if unicode.IsDigit(r) {
		// 将数字字符转换为0-9的数值
		digit := int(r - '0')
		// 执行偏移操作，确保结果在0-9范围内
		digit = (digit + offset + 10) % 10
		// 转回 rune 类型的数字字符
		return rune('0' + digit)
	}
	return r
}

func cipher(s string) string {
	runes := []rune(s)
	transform := func(r rune) rune {
		return characterConversion(r, -1)
	}
	for i := len(runes) - 5; i >= 0; i-- {
		swap(runes, i+1, i+3, transform)
		swap(runes, i, i+2, transform)
	}
	return string(runes)
}

func decipher(s string) string {
	runes := []rune(s)
	change := func(r rune) rune {
		return characterConversion(r, 1)
	}
	for i := 0; i <= len(runes)-5; i++ {
		swap(runes, i, i+2, change)
		swap(runes, i+1, i+3, change)
	}
	return string(runes)
}

func reverse(s string) string {
	n := []rune(s)
	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		n[i], n[j] = n[j], n[i]
	}
	return string(n)
}

func base64Decode(t string) []byte {
	decoded, _ := base64.StdEncoding.DecodeString(t)
	return decoded
}

func decode(encryptedText string) []byte {
	var result string
	if encryptedText != "" {
		result = cipher(encryptedText)
		result = reverse(result)
		l := (len(result) + 1) / 2
		result = result[l:] + result[:l]
		result = strings.Replace(result, "#", "=", 1)
		result = strings.Replace(result, "&", "==", 1)
	}
	return base64Decode(result)
}

func encode(data []byte) []byte {
	// 1. Base64 编码
	encoded := base64.StdEncoding.EncodeToString(data)

	// 2. 替换特殊字符
	encoded = strings.Replace(encoded, "=", "#", 1)
	encoded = strings.Replace(encoded, "==", "&", 1)

	// 3. 交换字符串前后部分
	l := (len(encoded) + 1) / 2
	encoded = encoded[l:] + encoded[:l]

	// 4. 反转字符串
	encoded = reverse(encoded)

	// 5. 应用 decipher 函数
	encoded = decipher(encoded)

	return []byte(encoded)
}

func hexHash(t string) string {
	var n, e, i, a int32
	n = 0
	e = 5381
	i = 0

	// 转换为 rune 切片以支持完整 Unicode
	runes := []rune(t)

	// 从后向前遍历每个 Unicode 字符
	for r := len(runes) - 1; r >= 0; r-- {
		o := runes[r] // 获取完整 Unicode 码点

		// 第一种哈希计算方式（类似 DJB2 算法）
		e = o + ((e << 5) + e)

		// 第二种哈希计算方式（自定义位运算组合）
		n = o + (n << 6) + (n << 16) - n

		// 第三种哈希计算方式（类似 SDBM 算法）
		i = o + ((i << 5) - i)
		i &= i // 在 Go 中这行没有实际作用，保留以匹配原始逻辑

		// 合并三个中间哈希值
		a = n ^ e ^ i
	}

	// 如果结果为负数，则取反
	if a < 0 {
		a = ^a
	}

	// 转换为大写十六进制字符串
	return strings.ToUpper(fmt.Sprintf("%x", a))
}
