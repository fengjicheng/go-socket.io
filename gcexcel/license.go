package gcexcel

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/wenzhenxi/gorsa"
)

// com.grapecity.documents.excel.internals.aX.f V9
// com.grapecity.documents.excel.internals.aU.a V8
// GCEXCEL_JAVA_DEPLOY_LICENSE_V9
// GCEXCEL_JAVA_DEPLOY_LICENSE_V8
const (
	GrapeCity = "GrapeCity"

	DeployV2 = "6bf630ea-22d3-47b5-bb9e-2102f3c52186"
	DevV2    = "383d4eff-9ef1-4198-ad4d-eb11035a7bc6"

	DeployV2GCLMTest = "e9b2a94e-afa0-40ab-be43-b86d3719c5f7"
	DevV2GCLMTest    = "64c28c83-fab2-4c5d-bd94-7e2dee4da186"

	LicenseFile = ".license"
	KeyFile     = ".key"
)

type licOpts struct {
	dev      bool // 是否为开授权
	gcLMTest bool
	expire   time.Duration // 过期时间
}

type license struct {
	a            bool
	uuid         uuid.UUID
	serialNumber string
	hostname     string
	activeTime   uint // 与 2000-01-01 00:00:00 相差的天数
	f            bool
	expiryTime   uint   // 与 2000-01-01 00:00:00 相差的天数
	version      string // Standard Unlimited
	i            string
	j            string

	buf *bytes.Buffer
}

func (l *license) writeValue(val string, b64e bool) {
	if val != "" {
		if b64e {
			enc := base64.StdEncoding.EncodeToString([]byte(val))
			all := strings.ReplaceAll(enc, "=", "")
			l.buf.WriteString(all)
		} else {
			l.buf.WriteString(val)
		}
	}
}

func (l *license) writeSep() {
	l.buf.WriteByte(',')
}

func (l *license) build() {
	l.buf.Reset()
	// 0. uuid
	l.writeValue(l.uuid.String(), true)
	l.writeSep()

	// 1. serialNumber
	l.writeValue(l.serialNumber, true)
	l.writeSep()

	// 2. hostname
	l.writeValue(l.hostname, true)
	l.writeSep()

	// 3. a
	l.writeValue(fmt.Sprintf("%v", l.a), true)
	l.writeSep()

	// 4. activeTime
	l.writeValue(fmt.Sprintf("%d", l.activeTime), true)
	l.writeSep()

	// 5. f
	l.writeValue(fmt.Sprintf("%v", l.f), true)
	l.writeSep()

	// 6. expiryTime
	l.writeValue(fmt.Sprintf("%d", l.expiryTime), true)
	l.writeSep()

	// 7.
	l.writeValue(fmt.Sprintf("%d", l.expiryTime), true)
	l.writeSep()

	// 8. version
	l.writeValue(l.version, true)
	l.writeSep()

	// 9. i
	l.writeValue(l.i, true)
	l.writeSep()

	// 10. j
	l.writeValue(l.j, true)

	source := l.buf.String()
	pri, _ := os.ReadFile("private.pem")
	sign, _ := gorsa.SignSha256WithRsa(source, string(pri))
	l.buf.WriteByte(';')
	l.buf.WriteString(strings.ReplaceAll(sign, "=", ""))
}

//NmJmNjMwZWEtMjJkMy00N2I1LWJiOWUtMjEwMmYzYzUyMTg2,
//NjA2NDExMDdYWFhYWFhYWDA4Mg,
//bWFjLW1pbmk,
//RmFsc2U,
//OTUyNQ,
//VHJ1ZQ,
//OTUzMg,
//OTUzMg,
//U3RhbmRhcmQ,,
