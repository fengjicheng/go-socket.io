package gcexcel

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/google/uuid"
)

type SerialNumber interface {
	Value() string
	MaskValue() string
}

type serialNumber struct {
	value     string
	maskValue string
}

func (sn *serialNumber) Value() string {
	return sn.value
}

func (sn *serialNumber) MaskValue() string {
	return sn.maskValue
}

func NewSerialNumber() (SerialNumber, error) {
	val, err := serialNumberGenerate()
	if err != nil {
		return nil, err
	}
	segments := strings.Split(val, "-")
	segments[2] = "XXXX"
	segments[3] = "XXXX"
	maskValue := strings.Join(segments, "")
	return &serialNumber{
		value:     val,
		maskValue: maskValue,
	}, nil
}

// 生成格式：xxxx-xxxx-xxxx-xxxx-xxx 的自增序列化编号
func serialNumberGenerate() (string, error) {
	// 步骤1：生成标准 UUID v7（核心：毫秒时间戳+随机数，有序且唯一）
	uuid7, err := uuid.NewV7()
	if err != nil {
		return "", fmt.Errorf("生成 UUID v7 失败：%w", err)
	}

	// 步骤2：按 - 拆分 UUID 为各段
	segments := strings.Split(uuid7.String(), "-")
	if len(segments) != 5 {
		return "", fmt.Errorf("无效的 UUID 格式，必须包含 4 个 - 分隔符")
	}

	// 步骤3：定义结果切片，存储每段转换后的 4 位数字
	var resultSegments []string

	// 步骤4：遍历每一段，进行十六进制→4 位十进制转换
	for _, seg := range segments {
		// 4.1：将十六进制字符串转为 big.Int 十进制大数（支持超长段，如 12 位十六进制）
		var segBigInt big.Int
		_, ok := segBigInt.SetString(seg, 16)
		if !ok {
			return "", fmt.Errorf("分段 %s 不是有效的十六进制字符串", seg)
		}

		// 4.2：映射到 0~9999 范围（% 10000），保证转为 4 位数字
		mod := big.NewInt(10000)
		seg4DigitBigInt := new(big.Int).Mod(&segBigInt, mod)
		seg4Digit := seg4DigitBigInt.Int64()

		// 4.3：补零到 4 位，存入结果切片
		resultSegments = append(resultSegments, fmt.Sprintf("%04d", seg4Digit))
	}

	// 步骤4：用 - 拼接各段，返回最终结果
	return strings.Join(resultSegments, "-"), nil
}
