package spread

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// SpreadJS 授权
var (
	lic             = "E879948536774266#B12c3MC36MEhXd9hWelJjT5QlRyFFaihDZXxGRQd7ba3mRTRmMh56VDtmMxlkdUlHOotkMjN4Zld6ZiFDerhTb4JTRy2UQBpUeLVTcMJES6ZTN6VnZyo5dKN6KaJVY9B7SLVXOLxmU8F4UxNjdzYndQZ5T4RVQ5pGRld6d8RmbahjdrNXNGpGMXdTUqJUQaZzUapEWMd7L5pEeUl5MkRHbBN7Y74mcYh4UHtUeEV6YUdHN75kehxkcZhlTntkc8lEc4FTY4lXM8ITSnNDVkRTVrFTeRF6dv3CRjhEbQpFajdTW9IzahFEUhNFOyglNSBzUX3kRIR5SvhXR5J4LxJmZ4kDRyUHZxRlQxdEb4YlRTh7S494Kr8kI0IyUiwiIxgTO5MUQ4IjI0ICSiwSNyEjN4AjMzMTM0IicfJye#4Xfd5nIIlkSCJiOiMkIsICOx8idgMlSgQWYlJHcTJiOi8kI1tlOiQmcQJCLiUTMzIzMwAyNwUDM5IDMyIiOiQncDJCLiYDM6ATNyAjMiojIwhXRiwiIx8CMuAjL7ITMiojIz5GRiwiI8+Y9sWY9QmZ0Jyp93uL9hKI0Aqo9Re09cu19R619HWa96mp930b9J0a9iojIh94QiwSZ5JHd0ICb6VkIsIiN6IDN7cjNzUDO4kTO7gjI0ICZJJCL35lI4JXYoNUY4FGRiwiI4VWZoNFdy3GclJlIbpjInxmZiwSZzxWYmpjIyNHZisnOiwmbBJye0ICRiwiI34TUqRlQsJjdzlkVa3UeohFWXZnT7IFdwk4VZ36UE3CRRFFNlR5cyZTRDF6K6JHWyI6Z8s4Y8BHMy3yYrxUaeQZb"
	fullLic         = "10.1.150.152,E872583357245163#B1Wa4dEMQRWeuNjcBFFTrZVTwZDelJ5crQEeDB7Mtd6dNhmV5kVQ43iehFTY8lVV5kEd7V4a5InZycnWKV7dzFFbHllRKtEOvg7cUJlYKZTSD3CbrRlMrkVQ9lHaEp4aBV5ZvJjTTJnUzoENxRVQ4BnYYhTN5dHOhZVSz3yL68USoRnYShDV8BFbWFmehVEVY5WQuFTM9MDdvFWU9NHZ83iV5FnNGRXMzoUR7IDNJ5GdBxUMap6aONHMxomaJ3ScU5kMrZXbD5mQJhnUvhXaLplbXh6MHdHcGRnU7cHRTl7TMNmT6l6cotke6dDa8A7cyhzLC94VS96T5AjevoWYyhWW5QFdSF4VzVnSlZTdhNkZ7gkUZdlc0NHejpVTIl5Zhh6KrNXVIJiOiMlIsISM7ADO9AjRiojIIJCL7gzM7QTO7ETN0IicfJye#4Xfd5nIIlkSCJiOiMkIsICOx8idgMlSgQWYlJHcTJiOi8kI1tlOiQmcQJCLiYTMxMDNwASNxUDM5IDMyIiOiQncDJCLiUjM5ATNyAjMiojIwhXRiwiIyUTMuATNx8SMuATMiojIz5GRiwiI8+Y9sWY9QmZ0Jyp93uL9hKI0Aqo9Re09cu19R619HWa96mp930b9J0a9iojIh94QiwSZ5JHd0ICb6VkIsIyM6ETN4IzN5MzM8UjM7gjI0ICZJJCL35lI4VWZoNFd49WYHJCLiQnchh6QhRXYEJCLiQXZlh6U4J7bwVmUiwiIlxmYhRFdvZXaQJyW0IyZsZmIsU6csFmZ0IiczRmI1pjIs9WQisnOiQkIsISP3cWSxZGcZlVaoBla8MmYCFXSmh4anpkcQNzUkFGRFB5R09kMGRVUZVFO9hTathuRxp"
	designerLic     = "Designer-E336191996128716#B1RRf9mSEVFVOhmd7J6LvQme4dTRzR7aUFlb6VmSrIFbCNlZrgVWFJ7dYJVejtmRsNWbkJGSwcTZBNHb6ZXZFFGUmtCM6l4aaxEZXNTVLxmaXVjdpZ6UwAHahlXR0FVSP36LB5mapFkcxYDeZZzLrBnbwRnb5IHRZNDVEZlc4UDZxAHMRNzahV7b59GUypGU0tmbGFGNplXV9J7N7N7LvEHWvpEbzAHcmFTb5hXSi5WSPJzKxInQXhzQ5VVdo3CNSRjMnNzdL9UVwIHUalUNBtWRrIFc4RDWCJXOXJzTyE4T8IUNvtyd5dTavYHTHZjahtGUx4kd6EnNzNzLVZzMnVTZnVVNxpVR5okMI9UTiojITJCLiY4QFRzN5YENiojIIJCL6cDO6kjN6UzN0IicfJye35XX3JCSJpkQiojIDJCLigTMuYHITpEIkFWZyB7UiojIOJyebpjIkJHUiwiI6UTMxIDMggDM5ATNyAjMiojI4J7QiwiI6AjNwUjMwIjI0ICc8VkIsICMuAjLw8CMvIXZudWazVGZiojIz5GRiwiI8+Y9sWY9QmZ0Jyp93uL9hKI0Aqo9Re09cu19R619HWa96mp930b9J0a9iojIh94QiwSZ5JHd0ICb6VkIsIiNxcDOyEjN9kTM9EjNzMjI0ICZJJCL355W0IyZsZmIsUWdyRnOiI7ckJye0ICbuFkI1pjIEJCLi4TPnJnVrtEaOpWV6QDVyFkVDVGZChnR5FXdMNnYOBjd0lHSxN7dZhGWFljdxN6L6IFZxoVWNlWbklTU6xkMP3iNW9ENr3SaH32TMQ"
	designerFullLic = "10.1.150.152,E815289921648826#B1ucYQxM6QZVzUsdHO9R6an3Ea4NzLIZFeQ5GWlhTRllXc9hGdTl7Z7cza8EUcwJkTORnTDNFT7hjRMZ6QMF4QVh5SvI6LY34LSxmYvoHOz3WTSx6avllSZlkZaV4RYF4SWNmVuR4T8JGeqF7SK3Cd4Z4dmhje7QjQmd4VIVzdO54a6gmej3WaEd4KQFEa9FHREd5L7pFW4knYt5GUqFGZBxGS7tmayFDU09mM9Q5URJTOPp7VpdjaIRGaZFkcYJkUDVzd4gjWHJjNI36ZCFkQYhHcOdmTkJDM6UTYjpEMkF6Qyd7UE5GTmlTM8RUVlJDMRpFV5pXezZkcEN7NvJ4aotCNZJlI0IyUiwiIxIzNCV4Q5EjI0ICSiwyN7MTMygDN4AjM0IicfJye=#Qf35VfikTWzMjI0IyQiwiI8EjL6BibvRGZB5icl96ZpNXZE5yUKRWYlJHcTJiOi8kI1tlOiQmcQJCLiYTMxMDNwASNxUDM5IDMyIiOiQncDJCLiUjM5ATNyAjMiojIwhXRiwiIyUTMuATNx8SMuATMiojIz5GRiwiI8+Y9sWY9QmZ0Jyp93uL9hKI0Aqo9Re09cu19R619HWa96mp930b9J0a9iojIh94QiwSZ5JHd0ICb6VkIsIiNygDO4YTMykTO8ITNxgjI0ICZJJye0ICRiwiI34TQPRURBl6NWd6Yn3GeMdEWBJlNStmZGtUWxQnTDdHMqhXc9cTelRzMYxWR7Y5Y0djQ6IFZQdDZFt6NYJFMttUbSRFcRBTV7M7ZxJ7VMBTVxlde"
)

func TestPlugins(t *testing.T) {
	plugins := []SJSPlugin{
		PluginDesigner,
		PluginPivotTable,
		PluginReportSheet,
		PluginGanttSheet,
		PluginTableSheet,
		PluginDataChart,
		PluginCollaboration,
		PluginAI,
	}
	for _, prod := range products {
		for _, p := range plugins {
			t.Log(prod.Name, "Support Plugin", p.Code(), p.Name(), p.BitMask(), p.IsSupport(*prod))
		}
	}

}

func TestSpreadJSLicense_Decrypt(t *testing.T) {
	//	{
	//	   "_r": 1332046125,
	//	   "H": "24AC5981",
	//	   "S": "N++NtKxSFV4lGqBTqdu2D94fbq/BuExoKTHFOWS0R6X28SaPAak29Y7chZPlHcD/owaQy1kU4dT3gI281yta1tpIxrKgNXYrLazMw4wTceDyKGSHXrm7csAltd3YTxJu/wLXJZS6ZABjQ7W0jF5skv8ZndxwgeDjuATtOVPvv3v3qSAxRlK9uKKpyaRZ+cJwZ2fuv56vHBLq5KyJAAO2E2tm8kx1bggegCc2Kh8yTvIq2kCWma2dSFoZowPDlWd8bhQrFT5N2eyhyuxD3oB3W4lD3iLkc/r0pxcK8gb2Xrv+aCE6rsTe4QQD/DSoYWI0tR7NvWXXhyOZVIsv2lBTjQ==",
	//	   "D": {
	//	       "Anl": {
	//	           "dsr": false,
	//	           "flg": [
	//	               "ReportSheet",
	//	               "DataChart"
	//	           ]
	//	       },
	//	       "Id": "879948536774266",
	//	       "Evl": true,
	//	       "CNa": "安徽晶奇网络科技股份有限公司",
	//	       "Dms": "127.0.0.1",
	//	       "Exp": "20250606",
	//	       "Crt": "20250507 032315",
	//	       "Prd": [
	//	           {
	//	               "N": "Spread JS v.18",
	//	               "C": "BJIH"
	//	           }
	//	       ]
	//	   }
	//	}
	sjs := NewSpreadJSLicense()
	sjs.Read(lic)

	assert.Equal(t, 1332046125, sjs.R())
	assert.Equal(t, "24AC5981", sjs.HexHash())
	assert.Equal(t, "N++NtKxSFV4lGqBTqdu2D94fbq/BuExoKTHFOWS0R6X28SaPAak29Y7chZPlHcD/owaQy1kU4dT3gI281yta1tpIxrKgNXYrLazMw4wTceDyKGSHXrm7csAltd3YTxJu/wLXJZS6ZABjQ7W0jF5skv8ZndxwgeDjuATtOVPvv3v3qSAxRlK9uKKpyaRZ+cJwZ2fuv56vHBLq5KyJAAO2E2tm8kx1bggegCc2Kh8yTvIq2kCWma2dSFoZowPDlWd8bhQrFT5N2eyhyuxD3oB3W4lD3iLkc/r0pxcK8gb2Xrv+aCE6rsTe4QQD/DSoYWI0tR7NvWXXhyOZVIsv2lBTjQ==", sjs.Sign())
	data := sjs.GetData()
	assert.Equal(t, false, data.Annual.Distribution)
	assert.Equal(t, []string{"ReportSheet", "DataChart"}, data.Annual.PluginFlags)
	assert.Equal(t, "879948536774266", data.Id)
	assert.Equal(t, true, data.Evaluation)
	assert.Equal(t, "安徽晶奇网络科技股份有限公司", data.CNa)
	assert.Equal(t, "127.0.0.1", data.Domains)
	assert.Equal(t, "20250606", data.Expiration)
	assert.Equal(t, "20250507 032315", data.CreateTime)
	assert.Equal(t, 1, len(data.Products))
	assert.Equal(t, "Spread JS v.18", data.Products[0].Name)
	assert.Equal(t, "BJIH", data.Products[0].Code)

	sjs.Read(fullLic)
	data = sjs.GetData()

	data.Domains = "teamwork-test.jqk8s.jqsoft.net,teamwork.jqk8s.jqsoft.net,10.1.40.93,10.1.150.152"
	data.Evaluation = false
	data.Annual.Distribution = true
	data.Expiration = "20250517"
	data.Products = products[:2]

	_ = sjs.Output(os.Stdout)
	println()
}

func TestDesignerLic(t *testing.T) {
	sjs := NewSpreadJSLicense()
	sjs.Read(designerLic)

	data := sjs.GetData()
	assert.Equal(t, "designer/0.0.0.0", data.Domains)

	sjs.Read(designerFullLic)
	data.Products = products[:2]
	assert.Equal(t, 2, len(data.Products))
	assert.Equal(t, "BJIH", data.Products[0].Code)
	assert.Equal(t, "33Y9", data.Products[1].Code)
	_ = sjs.Output(os.Stdout)
}

func TestInternalV18(t *testing.T) {
	l1 := ReadLicense(GCSpreadSheetsLicenseKeyV18)
	l2 := ReadLicense(GCSpreadSheetsDesignerLicenseKeyV18)

	_ = l1.Output(os.Stdout)
	println()

	_ = l2.Output(os.Stdout)
	println()

}

func TestInternalV19(t *testing.T) {
	l1 := ReadLicense(GCSpreadSheetsLicenseKeyV19)
	l2 := ReadLicense(GCSpreadSheetsDesignerLicenseKeyV19)

	_ = l1.Output(os.Stdout)
	println()

	_ = l2.Output(os.Stdout)
	println()

}
