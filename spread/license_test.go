package spread

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	//Please provide valid license key.
	GcSpreadSheetsLicenseKey         = "GrapeCity-Internal-Use-Only,E835481965572883#B17eYICbuFkI1pjIEJCLi4TPn96M7dVY7dWVkVTUMpGMoJzbWVEd8pFMT3UT7ZUSws6KElWamVXdDhDdplnZ8J5aTRmb626bvF5QNN6TH3UbkR5ao9EU8JnYvMVUShWVpdUR8dDO0JTZiNFWYplSDRlY9Nlc9lVYjNWWxRFUv4Ua6pGelNnNiR5NKVVZ8ITTRB5bZ9kVR9kdIhUaaJ7dJdnRshXcxoVbhxkeaJ4ZHFjY7dUaxUFeVZ6duVWWyRnYlVUbXp4Vy9EO9AlRKlXYGtGT9YnMtNXNKNTcud7LWJzNyIEZzwUckR6NzwUSxNmSycVWyRjULxETzMDTrh5U7k5Kn3md5FzYRdkV954V42UW9dDT5UndoZDSWRFTYJnbXVXSvs4QLhWMRp6TxgUYkV4R79EcIF4bzhDVLBzLGNjN6IHW8llTslkMLREay86SDJVexdWRHlWe896ZVhzZ5cmZjhEZVN6YiojITJCLiIzMGZTM8EkI0ICSiwyMwIDM6cDMzITM0IicfJye#4Xfd5nIIlkSCJiOiMkIsICOx8idgMlSgQWYlJHcTJiOi8kI1tlOiQmcQJCLiYzMxQTNwAiMwITM4IDMyIiOiQncDJCLiI7au26YukHdpNWZwFmcn9iKs46bj9yc5l6YzVWbuoCLytmLvNmLzVXajNXZt9iKsAnau26YukHdpNWZwFmcn9iKs46bj9Se4l6YlBXYydmLqwSbvNmL6VGZ9RXajVGchJ7ZuoCLw3GduMXdpN6cl5mLqwCcq9yc5l6YzVWbuoCLwpmLvNmLzVXajNXZt9iKs2WauMXdpN6cl5mLqwibj9SbvNmL9RXajVGchJ7ZuoCLzVnLzVXajNXZt9iKiojIz5GRiwiIzx6bvRlclB7bsVmdlRkI0ISYONkIsUWdyRnOiwmdFJCLiMDO8IzN5UjN9EDO4UzM8IiOiQWSiwSfdJCdyFGaDFGdhRkIsICdlVGaTRHduF6RiwiI4VWZoNFdy3GclJlIsISZsJWYUR7b6lGUislOicGbmJCLlNHbhZmOiIKc6J"
	GcSpreadSheetsDesignerLicenseKey = "GrapeCity-Internal-Use-Only,E395773961976736#B1JpMIyNHZisnOiwmbBJye0ICRiwiI34zdrJHd0t6SRN6NNBVM9UmZaB7dOV6TvJjUvcXdapXQD3WW0FGUoRUNwwkVWxkaD36RXlFbVxWb8kVekZjM0tWY0d5U4ZUZZh5c48GRMlWbNh7b8QHZrp7T7MmS9plMLVlSqBTePBzcLJzdxQXQl3SN9FmYBhHUO9UbT3SM9QVTQxGOlpmbU3GTYdGThFUaEpHNPFVSkJXdDVjaMVTQ7JkS4k4NaZ6LwljS7QUe5wUeH5WWpF7Yv5UN0NlQ7cDZJFXe4FkMTBnNClXQ89UZwlka8VVNntWOy5mTppnSEZGe4VWb4UVYLVHaQF4Q5g5dkRkYvl4U5hUZJB5S8cTNmlVWlBlSOJFRhdmTUdHVEtUUzpFNThUQEVlQHFGZ5MGOtlDb8sUd0R5axY5T8Y5YL3CeotGZ8t4b9lXVFZnU9dTbycjQTNVdxBDWSxWWVJmczI7RHd6bklkI0IyUiwiI5QEOwUDN6EjI0ICSiwSM6EDOxITOwkTM0IicfJye#4Xfd5nI9k5MzIiOiMkIsICOx8idg86bkRWQtIXZudWazVGRtMlSkFWZyB7UiojIOJyebpjIkJHUiwiIzAzN4UDMgIDMyEDNyAjMiojI4J7QiwiIw3GduMXdpN6cl5mLqwicr9ybj9Se4l6YlBXYydmLqwCcq9ybj9Se4l6YlBXYydmLqwCcq9ybj9yc5l6YzVWbuoCLuNmLt36YukHdpNWZwFmcn9iKs46bj9idlRWe4l6YlBXYydmLqwyc59yc5l6YzVWbuoCLt36YuMXdpN6cl5mLqwybp9yc5l6YzVWbuoCLwpmLzVXajNXZt9iKs46bj9Se4l6YlBXYydmLqwicr9ybj9yc5l6YzVWbuoiI0IyctRkIsIycs36bUJXZw3GblZXZEJiOiEmTDJCLlVnc4pjIsZXRiwiI6MzN6cTOxYTOzczN5kzMiojIklkIs4XXiQnchh6QhRXYEJCLiQXZlh6U4RnbhdkIsICdlVGaTRncvBXZSJCLiUGbiFGV43mdpBlIbpjInxmZiwSZwxZY"
)

func TestDesignerLicense(t *testing.T) {
	designer := ReadLicense(designerLic)
	assert.NotNil(t, designer)
	designer = ReadLicense(designerFullLic)
	assert.NotNil(t, designer)
	opts := make([]Options, 0)

	//opts = append(opts, WithCreateTime("2025-05-15 04:31:16"))
	opts = append(opts, WithDeadline("24h"))
	opts = append(opts, WithPlugin(int(PluginDesigner)))
	// 开发授权
	opts = append(opts, WithDistributionLicense())
	sjs := NewSpreadJSLicense(opts...)
	assert.NotNil(t, sjs)
	data := sjs.GetData()
	assert.NotNil(t, data)
	assert.NotNil(t, data.Annual)
	//assert.Equal(t, false, data.Annual.Distribution)
	assert.Equal(t, 0, len(data.Annual.PluginFlags))
	assert.Equal(t, "designer/0.0.0.0", data.Domains)
	assert.Equal(t, false, data.Evaluation)
	_ = sjs.Output(os.Stdout)
	println()
}

func TestDistributionLicense(t *testing.T) {
	opts := make([]Options, 0)
	// {
	//"Anl":{"dsr":false,"flg":["PivotTable","ReportSheet","DataChart","GanttSheet"]},
	//"Id":"872583357245163","Evl":true,"CNa":"安徽晶奇网络科技股份有限公司","Dms":"10.1.150.152",
	//"Exp":"20250525","Crt":"20250515 043116",
	//"Prd":[{"N":"Spread JS v.18","C":"BJIH"}]
	//}
	//opts = append(opts, WithCreateTime("2025-05-15 04:31:16"))
	opts = append(opts, WithDeadline("8760h"))
	opts = append(opts, WithLicenseType(Official))
	opts = append(opts, WithPlugin(int(PluginDesigner)))
	opts = append(opts, WithPlugin(int(PluginPivotTable)))
	opts = append(opts, WithPlugin(int(PluginReportSheet)))
	opts = append(opts, WithPlugin(int(PluginGanttSheet)))
	opts = append(opts, WithPlugin(int(PluginTableSheet)))
	opts = append(opts, WithPlugin(int(PluginDataChart)))
	opts = append(opts, WithDomain("*.jqk8s.jqsoft.net"))
	opts = append(opts, WithIP("10.1.*", "172.21.*"))
	sjs := NewSpreadJSLicense(opts...)
	assert.NotNil(t, sjs)
	data := sjs.GetData()
	assert.NotNil(t, data)
	//assert.Equal(t, "20250515 043116", data.CreateTime)
	//assert.Equal(t, "20250525", data.Expiration)
	file, err := os.OpenFile("license.js", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0644)
	assert.Nil(t, err)
	_ = sjs.Output(os.Stdout)
	println()

	_, _ = file.WriteString(fmt.Sprintf("// 表格授权，授权截止时间：%s\n", sjs.GetData().Expiration))
	_, _ = file.WriteString(fmt.Sprintf("GC.Spread.Sheets.LicenseKey = '%v';\n", sjs))
	_, _ = file.WriteString(fmt.Sprintf("// 设计器授权，授权截止时间：%s\n", sjs.GetData().Expiration))
	_, _ = file.WriteString(fmt.Sprintf("GC.Spread.Sheets.Designer.LicenseKey = '%v';\n", sjs))

}

func TestWebDesignerLicense(t *testing.T) {
	webDesigner := ReadLicense(designerFullLic)
	println(webDesigner)
	opts := make([]Options, 0)

	opts = append(opts, WebDesignerLicense())
	opts = append(opts, WithCreateTime("2025-05-15 04:31:16"))
	opts = append(opts, WithDeadline("240h"))
	opts = append(opts, WithPlugin(int(PluginDesigner)))
	//opts = append(opts, WithPlugin(int(PluginPivotTable)))
	//opts = append(opts, WithPlugin(int(PluginReportSheet)))
	//opts = append(opts, WithPlugin(int(PluginGanttSheet)))
	//opts = append(opts, WithPlugin(int(PluginTableSheet)))
	//opts = append(opts, WithPlugin(int(PluginDataChart)))
	opts = append(opts, WithDomain("127.0.0.1"))
	opts = append(opts, WithIP("10.1.*", "172.21.*"))
	// 正式授权需要加此选项
	//opts = append(opts, WithoutTrial())
	sjs := NewSpreadJSLicense(opts...)
	assert.NotNil(t, sjs)
	data := sjs.GetData()
	assert.NotNil(t, data)
	//assert.Equal(t, false, data.Annual.Distribution)
	//assert.Equal(t, 2, len(data.Annual.PluginFlags))
	assert.Equal(t, "20250515 043116", data.CreateTime)
	assert.Equal(t, "20250525", data.Expiration)
	_ = sjs.Output(os.Stdout)
	println()
}

func Test_license_GetData(t *testing.T) {

}

func Test_license_HexHash(t *testing.T) {

}

func Test_license_MarshalJSON(t *testing.T) {

}

func Test_license_Output(t *testing.T) {

}

func Test_license_PrefixGenerate(t *testing.T) {

}

func Test_license_R(t *testing.T) {

}

func Test_license_Read(t *testing.T) {

}

func Test_license_Sign(t *testing.T) {

}

func Test_license_UnmarshalJSON(t *testing.T) {

}
