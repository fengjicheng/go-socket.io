package spread

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	//Please provide valid license key.
	GCSpreadSheetsLicenseKeyV18         = "GrapeCity-Internal-Use-Only,E835481965572883#B17eYICbuFkI1pjIEJCLi4TPn96M7dVY7dWVkVTUMpGMoJzbWVEd8pFMT3UT7ZUSws6KElWamVXdDhDdplnZ8J5aTRmb626bvF5QNN6TH3UbkR5ao9EU8JnYvMVUShWVpdUR8dDO0JTZiNFWYplSDRlY9Nlc9lVYjNWWxRFUv4Ua6pGelNnNiR5NKVVZ8ITTRB5bZ9kVR9kdIhUaaJ7dJdnRshXcxoVbhxkeaJ4ZHFjY7dUaxUFeVZ6duVWWyRnYlVUbXp4Vy9EO9AlRKlXYGtGT9YnMtNXNKNTcud7LWJzNyIEZzwUckR6NzwUSxNmSycVWyRjULxETzMDTrh5U7k5Kn3md5FzYRdkV954V42UW9dDT5UndoZDSWRFTYJnbXVXSvs4QLhWMRp6TxgUYkV4R79EcIF4bzhDVLBzLGNjN6IHW8llTslkMLREay86SDJVexdWRHlWe896ZVhzZ5cmZjhEZVN6YiojITJCLiIzMGZTM8EkI0ICSiwyMwIDM6cDMzITM0IicfJye#4Xfd5nIIlkSCJiOiMkIsICOx8idgMlSgQWYlJHcTJiOi8kI1tlOiQmcQJCLiYzMxQTNwAiMwITM4IDMyIiOiQncDJCLiI7au26YukHdpNWZwFmcn9iKs46bj9yc5l6YzVWbuoCLytmLvNmLzVXajNXZt9iKsAnau26YukHdpNWZwFmcn9iKs46bj9Se4l6YlBXYydmLqwSbvNmL6VGZ9RXajVGchJ7ZuoCLw3GduMXdpN6cl5mLqwCcq9yc5l6YzVWbuoCLwpmLvNmLzVXajNXZt9iKs2WauMXdpN6cl5mLqwibj9SbvNmL9RXajVGchJ7ZuoCLzVnLzVXajNXZt9iKiojIz5GRiwiIzx6bvRlclB7bsVmdlRkI0ISYONkIsUWdyRnOiwmdFJCLiMDO8IzN5UjN9EDO4UzM8IiOiQWSiwSfdJCdyFGaDFGdhRkIsICdlVGaTRHduF6RiwiI4VWZoNFdy3GclJlIsISZsJWYUR7b6lGUislOicGbmJCLlNHbhZmOiIKc6J"
	GCSpreadSheetsDesignerLicenseKeyV18 = "GrapeCity-Internal-Use-Only,E395773961976736#B1JpMIyNHZisnOiwmbBJye0ICRiwiI34zdrJHd0t6SRN6NNBVM9UmZaB7dOV6TvJjUvcXdapXQD3WW0FGUoRUNwwkVWxkaD36RXlFbVxWb8kVekZjM0tWY0d5U4ZUZZh5c48GRMlWbNh7b8QHZrp7T7MmS9plMLVlSqBTePBzcLJzdxQXQl3SN9FmYBhHUO9UbT3SM9QVTQxGOlpmbU3GTYdGThFUaEpHNPFVSkJXdDVjaMVTQ7JkS4k4NaZ6LwljS7QUe5wUeH5WWpF7Yv5UN0NlQ7cDZJFXe4FkMTBnNClXQ89UZwlka8VVNntWOy5mTppnSEZGe4VWb4UVYLVHaQF4Q5g5dkRkYvl4U5hUZJB5S8cTNmlVWlBlSOJFRhdmTUdHVEtUUzpFNThUQEVlQHFGZ5MGOtlDb8sUd0R5axY5T8Y5YL3CeotGZ8t4b9lXVFZnU9dTbycjQTNVdxBDWSxWWVJmczI7RHd6bklkI0IyUiwiI5QEOwUDN6EjI0ICSiwSM6EDOxITOwkTM0IicfJye#4Xfd5nI9k5MzIiOiMkIsICOx8idg86bkRWQtIXZudWazVGRtMlSkFWZyB7UiojIOJyebpjIkJHUiwiIzAzN4UDMgIDMyEDNyAjMiojI4J7QiwiIw3GduMXdpN6cl5mLqwicr9ybj9Se4l6YlBXYydmLqwCcq9ybj9Se4l6YlBXYydmLqwCcq9ybj9yc5l6YzVWbuoCLuNmLt36YukHdpNWZwFmcn9iKs46bj9idlRWe4l6YlBXYydmLqwyc59yc5l6YzVWbuoCLt36YuMXdpN6cl5mLqwybp9yc5l6YzVWbuoCLwpmLzVXajNXZt9iKs46bj9Se4l6YlBXYydmLqwicr9ybj9yc5l6YzVWbuoiI0IyctRkIsIycs36bUJXZw3GblZXZEJiOiEmTDJCLlVnc4pjIsZXRiwiI6MzN6cTOxYTOzczN5kzMiojIklkIs4XXiQnchh6QhRXYEJCLiQXZlh6U4RnbhdkIsICdlVGaTRncvBXZSJCLiUGbiFGV43mdpBlIbpjInxmZiwSZwxZY"

	//Please provide valid license key.
	GCSpreadSheetsLicenseKeyV19         = "GrapeCity-Internal-Use-Only,E562651991126827#B1swIZzxWYmpjIyNHZisnOiwmbBJye0ICRiwiI34TUNZTd596K6tiWGJWeuNVQ4oEOlFWN0l5R9UHTQFERoVUM7kGanZWZzlzd0pkUtpUW6EjYxp5QsxUWh54TPl6UCl4NrQ5Sud5Z9dnQE9GWkRDeTJ4bwMkZCd5dkBlNpZFdv2kb7EVc9NHcwxEZIJVaS94UoNGc0pEZQJ7MMd5UWdWYOZGRnxkZWlTUm56bBBzTZRjd0JlSOdVR9dnWqlGN8tUT8QUUFhEU4tUZQlXSxlmYxlmczo6K03yVxNnMPZFMyUmeCdVezBFUvgGOOBjbFd6RyFHVHVncFJTQ4I5R6VXVmNEZmRzT8k4NZFGasdHUrImeY9GV6VWVSxmY7ZHeGh5ak56dkJ7KhZXYUFXVUJjZkdWe8IlcpVmS0VzLqZnSLVnbTlkZQZjM9MTYZdTSuFUaqBHTyJjQyolNNdTQpFjZC9kdGJGaBVncwcGcNx6NldkTkxkI0IyUiwiI7UTM7UzQ9YjI0ICSiwSN6kjNzgTMxgTM0IicfJye=#Qf35VfiEzQRdlI0IyQiwiI9EjL6BSKONEKTpEIkFWZyB7UiojIOJyebpjIkJHUiwiI6EDOzIDMggTMxETNyAjMiojI4J7QiwiIwpmLzVXajNXZt9iKs2WauMXdpN6cl5mLqwCcq9ybj9Se4l6YlBXYydmLqwicr9ybj9yc5l6YzVWbuoCLt36YuMXdpN6cl5mLqwibj9SbvNmL9RXajVGchJ7ZuoCLt36YuYXZklHdpNWZwFmcn9iKsA7b49yc5l6YzVWbuoCLt36YukHdpNWZwFmcn9iKsMXduMXdpN6cl5mLqwCcq9ybj9yc5l6YzVWbuoCLytmLvNmL9RXajVGchJ7ZuoiI0IyctRkIsIycs36bUJXZw3GblZXZEJiOiEmTDJCLlVnc4pjIsZXRiwiI7IDO6ITMxkTOxUjNyYTNiojIklkIs4XXiQnchh6QhRXYEJCLi86bpRXYy3mYhxGbvNkIsISSBJCLiQXZlh6U4RnbhdkIsICdlVGaTRncvBXZSJCLiUGbiFGV43mdpBlIbpjIMxNZ"
	GCSpreadSheetsDesignerLicenseKeyV19 = "GrapeCity-Internal-Use-Only,E355896284189812#B1JpBInxmZiwSZzxWYmpjIyNHZisnOiwmbBJye0ICRiwiI34zZpZzZrYGS0VzMVp5TEd5dJx4LMhXbzEFV8gWNRJ7dSJTMxRHRK3SbtdWUtlXUhx6cSd7K6hDUlVkSEBjSoh4KwEje8E7NiFneR3WSxcmYT3iSIZDcHljZMdWRr2EO0tWOxkjTUlHdjN4UKd4T5N5dLNkZlVzR6AnYEFzLu5WM6R5cBJzbVNTbytkW5QTRQdXan3ibp3UUFBneadHcIx6c5lGZtZ7M6hGaCN7YLlVO4wWUDZnZZdWbNpENrFXO4RzKpRkV8EVWB5WSxRnc4tEZvNDS7hDbJpHd4old7pFS4hzRll7b9dzMTtGaKhUOkxmZN5ESXZWehNkR7clNzomTnBja0FlZzRmYw3kb0JUcVpGaoNXQYF7M8FnaNp6Mj3CVOJWe8dndvVTZnxmTyVGZSVjVi9USrcXd5llMGZkZCVFRzNzSRp7VMh4ZjhUNzNjWJxmdYNkI0IyUiwiI5YkRyczMxQjI0ICSiwCM9MDM4kzN8ETM0IicfJye35XX3JySUljRiojIDJCLikTMuYHIp84Qo86bkRWQtIXZudWazVGRtMlSkFWZyB7UiojIOJyebpjIkJHUiwiIzADN4IDMggTMxETNyAjMiojI4J7QiwiIwpmLvNmL9RXajVGchJ7ZuoCLwpmLvNmLzVXajNXZt9iKs46bj9Se4l6YlBXYydmLqwybp9yc5l6YzVWbuoCLytmLvNmL9RXajVGchJ7ZuoCLuNmLt36YukHdpNWZwFmcn9iKs46bj9idlRWe4l6YlBXYydmLqwyc59yc5l6YzVWbuoCLw3GduMXdpN6cl5mLqwSbvNmLzVXajNXZt9iKsAnauMXdpN6cl5mLqwicr9ybj9yc5l6YzVWbuoiI0IyctRkIsIycs36bUJXZw3GblZXZEJiOiEmTDJCLlVnc4pjIsZXRiwiIyEDO9gTM4gjM6kDO5UzMiojIklkIs4XXiQnchh6QhRXYEJCLiQXZlh6U4RnbhdkIsIibvlGdhJ7biFGbs36QiwiIJFkIsICdlVGaTRncvBXZSJCLiUGbiFGV43mdLBPI"
)

func TestDesignerLicense(t *testing.T) {
	designer := ReadLicense(designerLic)
	assert.NotNil(t, designer)
	designer = ReadLicense(designerFullLic)
	assert.NotNil(t, designer)
	opts := make([]Options, 0)

	//opts = append(opts, WithCreateTime("2025-05-15 04:31:16"))
	opts = append(opts, WithDeadline("24h"))
	opts = append(opts, WithWebDesigner())
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

func TestDesignerLicenseV19(t *testing.T) {
	designer := ReadLicense(GCSpreadSheetsDesignerLicenseKeyV19)
	assert.NotNil(t, designer)

	opts := make([]Options, 0)
	//opts = append(opts, WithCreateTime("2025-05-15 04:31:16"))
	opts = append(opts, WithMajor(19))
	opts = append(opts, WithDeadline("24h"))
	opts = append(opts, WithWebDesigner())
	opts = append(opts, WithPlugin(255))
	// 开发授权
	//opts = append(opts, WithLicenseType(Official))
	opts = append(opts, WithDistributionLicense())
	sjs := NewSpreadJSLicense(opts...)
	assert.NotNil(t, sjs)
	data := sjs.GetData()
	assert.NotNil(t, data)
	assert.NotNil(t, data.Annual)
	//assert.Equal(t, false, data.Annual.Distribution)
	//assert.Equal(t, 0, len(data.Annual.PluginFlags))
	assert.Equal(t, "designer/0.0.0.0", data.Domains)
	//assert.Equal(t, true, data.Evaluation)
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
	opts = append(opts, WithPlugin(uint(V18Plugins)))
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
	opts = append(opts, WithPlugin(PluginDesigner.BitMask()))
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
