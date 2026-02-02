package gcexcel

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/wenzhenxi/gorsa"
)

func TestGenKeys(t *testing.T) {
	expected := "Ldo4Gj8o2osII/U6AKHyLgbG/na2s5jjKGLcjkR+pp9m87xGFxUiedM0qMJCK12zRifG3+ZnycwuDMqPvRi26hbo29rtp0h9X9/mPazDAg76S44s+ewFBfYWc39XLP2liRv5FPd88m5n8BkHnBT6uOKTBGCBumkmkKAUNY2HZiNqq1Z9oJTnjSx+RK3oIUFUbFDSa93ZI+izWHRYgmsIqQd59Nd14y32f6rpLqgR+isBE5YlAcb57bb5+OdGjA10EQQ4KDnnOIQy6RgDr1xg00BAmDedpF/jr3nhylVH98RuzN1vU09NDUDHZNePBKofrLNFcU7W8b8vq1JJkP3sgw=="
	file, err := os.ReadFile("private.pem")
	assert.Nil(t, err)
	s := string(file)
	sign, err := gorsa.SignSha256WithRsa("123", s)
	assert.Nil(t, err)
	assert.Equal(t, expected, sign)

	pemBytes, _ := os.ReadFile("public.pem")
	err = gorsa.VerifySignSha256WithRsa("123", expected, string(pemBytes))
	assert.Nil(t, err)
}

func TestTime(t *testing.T) {
	l1, l2 := RelativeDays(time.Hour * 24 * 7)
	assert.Equal(t, 7, l2-l1)
	fmt.Println(l1, l2)
}

func TestName(t *testing.T) {
	decodeString, err := base64.StdEncoding.DecodeString("NjA2NDExMDdYWFhYWFhYWDA4Mg==")
	assert.Nil(t, err)
	assert.Equal(t, "60641107XXXXXXXX082", string(decodeString))

	decodeString, err = base64.StdEncoding.DecodeString("NjA2NDExMDdYWFhYWFhYWDA4Mg")
	assert.NotNil(t, err)
	assert.NotEqual(t, "60641107XXXXXXXX082", string(decodeString))
}

func TestLicence(t *testing.T) {
	// 初始化生成器，起始基数为 1（可改为任意大于 0 的数）
	ser, err := NewSerialNumber()
	assert.Nil(t, err)
	t.Log(ser.Value())
	t.Log(ser.MaskValue())

	uid, err := uuid.Parse("6bf630ea-22d3-47b5-bb9e-2102f3c52186")

	l := license{
		a:            false, // 是否过期
		uuid:         uid,
		serialNumber: ser.MaskValue(),
		hostname:     "",
		activeTime:   0,
		f:            true, // 是否使用
		expiryTime:   0,
		version:      "Standard",

		buf: bytes.NewBuffer(nil),
	}

	l1, l2 := RelativeDays(time.Hour * 24 * 7)
	l.activeTime = l1
	l.expiryTime = l2

	l.build()

	log.Print(l.buf.String())

}
