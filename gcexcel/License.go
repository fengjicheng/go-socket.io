package gcexcel

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/wenzhenxi/gorsa"
)

type license struct {
	a            bool
	uuid         uuid.UUID
	serialNumber string
	hostname     string
	activeTime   uint // 与 2000-01-01 00:00:00 相差的天数
	f            bool
	expiryTime   uint // 与 2000-01-01 00:00:00 相差的天数
	version      string
	i            string
	j            string
	//private boolean a;
	//private UUID uuid;
	//private String serialNumber;
	//private String hostname;
	//private int activeTime;
	//private boolean f;
	//private int expiryTime;
	//private String version; // Unlimited
	//private String i;
	//private String j;

	buf *bytes.Buffer
}

func (l *license) writeValue(val string) {
	if val != "" {
		enc := base64.StdEncoding.EncodeToString([]byte(val))
		all := strings.ReplaceAll(enc, "=", "")
		l.buf.WriteString(all)
	}
}

func (l *license) writeSep() {
	l.buf.WriteByte(',')
}

func (l *license) build() {
	l.buf.Reset()
	// 0. uuid
	l.writeValue(l.uuid.String())
	l.writeSep()

	// 1. serialNumber
	l.writeValue(l.serialNumber)
	l.writeSep()

	// 2. hostname
	l.writeValue(l.hostname)
	l.writeSep()

	// 3. a
	l.writeValue(fmt.Sprintf("%v", l.a))
	l.writeSep()

	// 4. activeTime
	l.writeValue(fmt.Sprintf("%d", l.activeTime))
	l.writeSep()

	// 5. f
	l.writeValue(fmt.Sprintf("%v", l.f))
	l.writeSep()

	// 6. expiryTime
	l.writeValue(fmt.Sprintf("%d", l.expiryTime))
	l.writeSep()

	// 7.
	l.writeValue(fmt.Sprintf("%d", l.expiryTime))
	l.writeSep()

	// 8. version
	l.writeValue(l.version)
	l.writeSep()

	// 9. i
	l.writeValue(l.i)
	l.writeSep()

	// 10. j
	l.writeValue(l.j)

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
