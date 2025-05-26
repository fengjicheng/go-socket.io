package spread

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

type LicenseReader interface {
	Read(lic string) License           // 读取解析授权文件
	ReadFromFile(file os.File) License // 从文件读取解析授权文件
}

type License interface {
	GetData() *Data // 授权数据

	R() int          // todo 目前不清楚如何生成
	Sign() string    // 签名 todo 目前不清楚如何生成
	HexHash() string // 十六进制哈希

	LicenseReader

	PrefixGenerate(data Data) string // 生成前缀
	Bytes() []byte                   // 授权文件内容
	Output(writer io.Writer) error   // 输出授权文件

	json.Marshaler
	json.Unmarshaler
}

type PrefixGen = func(data Data) string

type options struct {
	major    int       // 主版本号
	sep      string    // 分隔符
	prefixFn PrefixGen // 前缀生成函数

	createdAt time.Time     // 创建时间
	duration  time.Duration // 有效时长

	plugins     int    // 插件编码数字和
	description string // 描述

	domains        []string // 域名
	noLimitDomains bool     // 是否无限制域名
	ips            []string // IP
	noLimitIps     bool     // 是否无限制 IP

	localDesigner bool // 是否本地设计器
	webDesigner   bool // 是否在线设计器

	//designer int // 1 本地设计器 2 在线设计器 3 本地设计器+在线设计器
	licType int // 授权类型 0 试用版 1 正式版 2 分发版
}

type license struct {
	opts options

	prefix  string // 前缀
	licData string // 加密数据

	r         int    // _r
	hash      string // H 授权数据 16 进制
	signature string // S 签名
	data      *Data  // D 授权数据
}

type jsonLic struct {
	R int    `json:"_r,omitempty"` //
	H string `json:"H,omitempty"`  // H 授权数据 Hash 16 进制
	S string `json:"S,omitempty"`  // S 签名
	D Data   `json:"D,omitempty"`  // D 授权数据
}

func (l *license) GetData() *Data {
	return l.data
}

func (l *license) R() int {
	return l.r
}

func (l *license) Sign() string {
	return l.signature
}

func (l *license) HexHash() string {
	enc, _ := json.Marshal(l.data)
	// <前缀>#<分隔符><授权JSON数据>
	s := fmt.Sprintf("%s#%s%s", l.PrefixGenerate(*l.data), l.opts.sep, string(enc))
	var n, e, i, a int32
	n = 0
	e = 5381
	i = 0

	// 转换为 rune 切片以支持完整 Unicode
	runes := []rune(s)

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

func (l *license) Read(lic string) License {
	reg := fmt.Sprintf("(%s)#(%s)(%s)", ".*", l.opts.sep, ".*")
	all := regexp.MustCompile(reg).FindStringSubmatch(lic)
	l.prefix = all[1]
	l.licData = all[3]

	bData := decode(l.licData)
	_ = json.Unmarshal(bData, l)
	return l
}

func (l *license) ReadFromFile(file os.File) License {
	// todo 实现从文件读取
	return l
}

func (l *license) MarshalJSON() ([]byte, error) {
	jl := &jsonLic{}
	jl.R = l.r
	jl.H = l.hash
	jl.S = l.signature
	jl.D = *l.data
	return json.Marshal(jl)
}

func (l *license) UnmarshalJSON(data []byte) error {
	jl := &jsonLic{}
	err := json.Unmarshal(data, jl)
	l.r = jl.R
	l.hash = jl.H
	l.signature = jl.S
	l.data = &jl.D
	return err
}

func (l *license) PrefixGenerate(data Data) string {
	if l.opts.prefixFn != nil {
		return l.opts.prefixFn(data)
	}
	return fmt.Sprintf("%s%s", "E", data.Id)
}

func (l *license) Bytes() []byte {
	l.build()
	buf := bytes.Buffer{}
	buf.WriteString(l.prefix)
	buf.WriteByte('#')
	buf.WriteString(l.opts.sep)
	buf.WriteString(l.licData)
	return buf.Bytes()
}

func (l *license) String() string {
	return string(l.Bytes())
}

func (l *license) Output(writer io.Writer) error {
	_, err := writer.Write(l.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func (l *license) build() {
	// 只有新建的时候才需要 build
	if l.data != nil {
		return
	}

	opts := l.opts

	evaluation := true
	distribution := false
	switch opts.licType {
	case 0:
	case 1:
		evaluation = false
	case 2:
		distribution = true
		evaluation = false
	default:
	}

	l.data = &Data{
		Annual:     &Annual{distribution, PluginsFrom(opts.plugins)},
		Id:         fmt.Sprintf("%d", opts.createdAt.Unix()),
		Evaluation: evaluation, // 试用即评估
		CNa:        opts.description,
		Domains:    "",
		Ips:        strings.Join(opts.ips, ","),
		Expiration: opts.createdAt.Add(opts.duration).Format("20060102"),
		CreateTime: opts.createdAt.Format("20060102 150405"),
	}

	// 1. 允许 ip
	if opts.noLimitIps {
		l.data.Ips = ""
	} else {
		ips := deduplicate(opts.ips)
		opts.ips = ips
		l.data.Ips = strings.Join(ips, ",")
	}

	// 2. 允许域名
	if opts.noLimitDomains {
		l.data.Domains = ""
	} else {
		dms := deduplicate(opts.domains)
		designer := int(PluginDesigner)
		if (opts.plugins & designer) == designer {
			dms = append(dms, "designer/0.0.0.0")
		}
		// todo 前端代码并未校验 Ips，暂时将所有 ip 放入 Domains
		dms = deduplicate(append(dms, opts.ips...))
		l.data.Domains = strings.Join(dms, ",")
	}

	// 3. 构建 data
	var products []*Product
	if ps, ok := prods[opts.major]; ok {
		for _, p := range ps {
			if p.designer == false {
				// 基础功能
				products = append(products, p)
				break
			}
		}
	}

	designer := int(PluginDesigner)
	if (opts.plugins & designer) == designer {
		if ps, ok := prods[opts.major]; ok {
			for _, p := range ps {
				if p.designer == true {
					// web designer 在线表格编辑器
					products = append(products, p)
					break
				}
			}
		}
	}
	l.data.Products = products

	// 2 todo 不知道怎么生成，暂不处理
	//l.r = 0

	// 3. 生成 hash
	h := l.HexHash()
	if l.hash != h {
		l.hash = h
	}

	// 4. 生成 signature
	s := l.Sign()
	if l.signature != s {
		l.signature = s
	}

	// 5. 生成 prefix
	l.prefix = fmt.Sprintf("%s", l.PrefixGenerate(*l.data))

	// 6. 生成 license 加密文本
	licJson, _ := json.Marshal(l)
	l.licData = string(encode(licJson))
}

type Options func(opts *options)

func NewSpreadJSLicense(opts ...Options) License {
	// licType = 2 页面 js 需要 GC.Spread.Sheets.Workbook["lm"] = 1;
	o := &options{sep: "B1", major: 18, createdAt: time.Now(), duration: time.Hour * 24 * 30}
	for _, opt := range opts {
		opt(o)
	}
	lic := &license{opts: *o}
	lic.build()
	return lic
}

func ReadLicense(txt string, sep ...string) License {
	o := &options{sep: "B1", major: 18, createdAt: time.Now(), duration: time.Hour * 24 * 30}
	if len(sep) > 0 && sep[0] != "" && sep[0] != "B1" {
		o.sep = sep[0]
	}
	lic := &license{opts: *o}
	lic.Read(txt)
	return lic
}

func WithSeparator(sep string) Options {
	return func(opts *options) {
		if sep != "" {
			opts.sep = sep
		}
	}
}

// WithCreateTime 时间格式 YYYY-MM-DD hh:mm:ss
func WithCreateTime(timeString string) Options {
	return func(opts *options) {
		if timeString != "" {
			if t, err := time.Parse("2006-01-02 15:04:05", timeString); err == nil {
				opts.createdAt = t
				return
			}
		}
		opts.createdAt = time.Now()
	}
}

// WithDeadline 24h30m60s
func WithDeadline(duration string) Options {
	return func(opts *options) {
		if duration != "" {
			if d, err := time.ParseDuration(strings.ToLower(duration)); err == nil {
				opts.duration = d
				return
			}
		}
		opts.duration = time.Hour * 24 * 30
	}
}

func WithDomain(domain ...string) Options {
	return func(opts *options) {
		opts.domains = append(opts.domains, domain...)
	}
}

func WithNoLimitDomains() Options {
	return func(opts *options) {
		opts.noLimitDomains = true
	}
}

func WithIP(ip ...string) Options {
	return func(opts *options) {
		opts.ips = append(opts.ips, ip...)
	}
}

func WithNoLimitIps() Options {
	return func(opts *options) {
		opts.noLimitIps = true
	}
}

func WithPrefix(gen PrefixGen) Options {
	return func(opts *options) {
		if gen != nil {
			opts.prefixFn = gen
		}
	}
}

func WithLicenseType(licType int) Options {
	return func(opts *options) {
		opts.licType = licType
	}
}

// WithFormalLicense 正式授权
func WithFormalLicense() Options {
	return WithLicenseType(1)
}

// WithDistributionLicense 分发授权
func WithDistributionLicense() Options {
	return WithLicenseType(2)
}

func WithPlugin(mask int) Options {
	return func(opts *options) {
		opts.plugins = opts.plugins | mask
	}
}

func WebDesignerLicense() Options {
	return func(opts *options) {
		opts.webDesigner = true
	}
}

func WithWebDesigner() Options {
	return WithPlugin(int(PluginDesigner))
}

func deduplicate(slice []string) []string {
	seen := make(map[string]struct{}) // 使用struct{}节省内存
	result := make([]string, 0, len(slice))

	for _, s := range slice {
		if _, exists := seen[s]; !exists {
			seen[s] = struct{}{}
			result = append(result, s)
		}
	}
	return result
}
