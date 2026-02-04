package spread

import (
	"fmt"
	"strings"
)

var (
	products = []*Product{
		{18, "Spread JS v.18", "BJIH"},
		{18, "SpreadJS-Designer-Addon v.18", "33Y9"},
		{19, "Spread JS(CN) v.19", "WQC1"},
		{19, "SpreadJS-Designer-Addon(CN) v.19", "F9TK"},
	}

	PluginDesigner    SJSPlugin = &plugin{"表格编辑器", "Designer", 1}
	PluginPivotTable  SJSPlugin = &plugin{"数据透视表", "PivotTable", 2}
	PluginReportSheet SJSPlugin = &plugin{"报表", "ReportSheet", 3}
	PluginGanttSheet  SJSPlugin = &plugin{"甘特图", "GanttSheet", 4}
	PluginTableSheet  SJSPlugin = &plugin{"集算表", "TableSheet", 5}
	PluginDataChart   SJSPlugin = &plugin{"数据图表", "DataChart", 6}

	PluginCollaboration SJSPlugin = &plugin{"协同插件", "Collaboration", 7}
	PluginAI            SJSPlugin = &plugin{"AI 智能助手", "AI", 8}

	// V18Plugins V18 支持插件
	V18Plugins = 1<<6 - 1
	// V19Plugins V19 支持插件
	V19Plugins = 1<<8 - 1
)

type Product struct {
	Major uint   `json:"-"`
	Name  string `json:"N"` // 产品名称
	Code  string `json:"C"` // 产品代码
}

func (p *Product) IsSpreadJSDesigner() bool {
	return strings.Contains(p.Name, "SpreadJS-Designer")
}

type Data struct {
	Annual     *Annual    `json:"Anl,omitempty"` // 授权信息
	Id         string     `json:"Id,omitempty"`  // 授权标识
	Evaluation bool       `json:"Evl"`           // 是否为评估使用版
	CNa        string     `json:"CNa,omitempty"` // 客户名称
	Domains    string     `json:"Dms,omitempty"` // 域名
	Ips        string     `json:"Ips,omitempty"` // IP 地址
	Expiration string     `json:"Exp,omitempty"` // 过期时间
	CreateTime string     `json:"Crt,omitempty"` // 创建时间
	Products   []*Product `json:"Prd,omitempty"` // 产品信息
}

type Annual struct {
	Distribution bool     `json:"dsr"` // 是否为分发版
	PluginFlags  []string `json:"flg"` // 插件标志
}

// SJSPlugin
// -------------------------------------------------
// V8
// 表格编辑器 	Designer
// 数据透视表		PivotTable
// 报表			ReportSheet
// 甘特图		GanttSheet
// 集算表		TableSheet
// 数据图表		DataChart
// -------------------------------------------------
// V9
// 协同插件		Collaboration
// AI 智能助手	AI
// -------------------------------------------------
type SJSPlugin interface {
	// Name 插件名称
	Name() string
	// Code 插件标识
	Code() string
	// BitMask 插件位掩码
	BitMask() uint
	// IsSupport 判断该产品是否支持
	IsSupport(product Product) bool

	fmt.Stringer
}

type plugin struct {
	name string
	code string
	mask uint
}

func (p *plugin) Name() string {
	return p.name
}

func (p *plugin) Code() string {
	return p.code
}

func (p *plugin) BitMask() uint {
	return 1 << (p.mask - 1)
}

func (p *plugin) IsSupport(product Product) bool {
	m := 1 << (p.BitMask() - 1)
	if product.Major == 18 {
		return m&V18Plugins == m
	} else if product.Major == 19 {
		return m&V19Plugins == m
	}
	return false
}

func (p *plugin) String() string {
	return p.code
}

func PluginsFrom(mask uint) []string {
	plugins := make([]SJSPlugin, 0)
	add(&plugins, mask, PluginPivotTable)
	add(&plugins, mask, PluginReportSheet)
	add(&plugins, mask, PluginGanttSheet)
	add(&plugins, mask, PluginTableSheet)
	add(&plugins, mask, PluginDataChart)

	add(&plugins, mask, PluginCollaboration)
	add(&plugins, mask, PluginAI)
	result := make([]string, len(plugins))
	for i, sjsPlugin := range plugins {
		result[i] = sjsPlugin.String()
	}
	return result
}

func add(plugins *[]SJSPlugin, mask uint, plugin SJSPlugin) {
	if mask&plugin.BitMask() == plugin.BitMask() {
		*plugins = append(*plugins, plugin)
	}
}
