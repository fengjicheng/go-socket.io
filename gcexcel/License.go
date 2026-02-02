package gcexcel

import (
	"bytes"
	"fmt"

	"github.com/google/uuid"
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

func (l *license) build() {
	l.buf.Reset()
	// 0. uuid
	l.buf.WriteString(l.uuid.String())
	l.buf.WriteByte(',')

	// 1. serialNumber
	l.buf.WriteString(l.serialNumber)
	l.buf.WriteByte(',')

	// 2. hostname
	l.buf.WriteString(l.hostname)
	l.buf.WriteByte(',')

	// 3. a
	if l.a {
		l.buf.WriteString("True")
	} else {
		l.buf.WriteString("False")
	}
	l.buf.WriteByte(',')

	// 4. activeTime
	if l.activeTime > 0 {
		l.buf.WriteString(fmt.Sprintf("%d", l.activeTime))
	}
	l.buf.WriteByte(',')

	// 5. f
	if l.f {
		l.buf.WriteString("True")
	} else {
		l.buf.WriteString("False")
	}
	l.buf.WriteByte(',')

	// 6. expiryTime
	if l.expiryTime > 0 {
		l.buf.WriteString(fmt.Sprintf("%d", l.expiryTime))
	}
	l.buf.WriteByte(',')

	// 7.
	if l.expiryTime > 0 {
		l.buf.WriteString(fmt.Sprintf("%d", l.expiryTime))
	}
	l.buf.WriteByte(',')

	// 8. version
	l.buf.WriteString(l.version)
	l.buf.WriteByte(',')

	// 9. i
	l.buf.WriteString(l.i)
	l.buf.WriteByte(',')

	// 10. j
	l.buf.WriteString(l.j)
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
